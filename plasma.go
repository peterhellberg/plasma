package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"os/exec"

	"github.com/lucasb-eyer/go-colorful"
)

var (
	width  = flag.Int("w", 512, "Width of the image")
	height = flag.Int("h", 512, "Height of the image")
	scale  = flag.Float64("s", 16.0, "Scale of the plasma")
	ofn    = flag.String("o", "plasma.png", "Output file name")
	pfn    = flag.String("p", "palette.png", "Palette file name")

	show = flag.Bool("show", false, "Show the generated image")
)

func main() {
	flag.Parse()

	w := *width
	h := *height

	palette := generatePalette()

	renderPalette(palette)

	plasma := generatePlasma(w, h, *scale)

	renderPlasma(w, h, plasma, palette)
}

func generatePalette() [255]colorful.Color {
	var palette [255]colorful.Color

	for x := 0; x < 255; x++ {
		palette[x] = colorful.Hsv(float64(x)*1.411764706, 1, 1)
	}

	return palette
}

func renderPalette(palette [255]colorful.Color) {
	w := 255
	h := 10

	p := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p.Set(x, y, palette[x])
		}
	}

	if file, err := os.Create(*pfn); err == nil {
		defer file.Close()
		if err := png.Encode(file, p); err == nil {
			open(*pfn)
		}
	}
}

func generatePlasma(w, h int, s float64) [][]uint8 {
	var plasma = make([][]uint8, w)

	for x := 0; x < w; x++ {
		plasma[x] = make([]uint8, h)

		for y := 0; y < h; y++ {
			c := uint8(
				(128.0 + (128.0 * math.Sin(float64(x)/s)) +
					128.0 + (128.0 * math.Sin(float64(y)/s))) / 2.0)

			plasma[x][y] = c
		}
	}
	return plasma
}

func renderPlasma(w, h int, plasma [][]uint8, palette [255]colorful.Color) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for s := 0; s < 64; s++ {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				m.Set(x, y, palette[(plasma[x][y]+uint8(s))%255])
			}
		}

		fn := fmt.Sprintf("%03d-%s", s, *ofn)

		if file, err := os.Create(fn); err == nil {
			defer file.Close()
			png.Encode(file, m)
		}
	}
}

func open(fn string) {
	if *show {
		exec.Command("open", fn).Run()
	}
}
