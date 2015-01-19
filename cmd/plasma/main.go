package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"time"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"
)

var (
	width  = flag.Int("w", 512, "Width of the image")
	height = flag.Int("h", 512, "Height of the image")
	frames = flag.Int("n", 1, "Number of frames to generate")
	scale  = flag.Float64("s", 16.0, "Scale of the plasma")
	ofn    = flag.String("o", "plasma.png", "Output file name")
	pfn    = flag.String("p", "palette.png", "Palette file name")

	show = flag.Bool("show", false, "Show the generated image")
)

func main() {
	flag.Parse()

	p := plasma.New(*width, *height, *scale)
	pa := palette.DefaultGradient

	renderPalette(pa, *pfn)
	renderPlasma(*width, *height, p, pa, *ofn)
}

func renderPalette(pa *palette.Palette, fn string) {
	if file, err := os.Create(fn); err == nil {
		defer file.Close()

		if err := png.Encode(file, pa.Image()); err == nil {
			open(fn)
		}
	}
}

func renderPlasma(w, h int, p *plasma.Plasma, pa *palette.Palette, fn string) {
	if *frames > 1 {
		for s := 0; s < *frames; s++ {
			m := p.Image(w, h, s, pa)

			if file, err := os.Create(fmt.Sprintf("%03d-%s", s, fn)); err == nil {
				defer file.Close()

				png.Encode(file, m)
			}
		}
	} else {
		m := p.Image(w, h, int(time.Now().Unix()), pa)

		if file, err := os.Create(fn); err == nil {
			defer file.Close()

			if err := png.Encode(file, m); err == nil {
				open(fn)
			}
		}
	}
}

func open(fn string) {
	if *show {
		exec.Command("open", fn).Run()
	}
}
