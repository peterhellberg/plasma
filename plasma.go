package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/peterhellberg/plasma/gradient"
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

	w := *width
	h := *height

	palette := generateGradientPalette(gradient.Table{
		{gradient.Hex("#005994"), 0.00},

		{gradient.Hex("#6F8525"), 0.12},
		{gradient.Hex("#B2C85B"), 0.14},
		{gradient.Hex("#EAECB8"), 0.16},
		{gradient.Hex("#002440"), 0.18},
		{gradient.Hex("#005994"), 0.20},

		{gradient.Hex("#6F8525"), 0.32},
		{gradient.Hex("#B2C85B"), 0.34},
		{gradient.Hex("#EAECB8"), 0.36},
		{gradient.Hex("#002440"), 0.38},
		{gradient.Hex("#005994"), 0.40},

		{gradient.Hex("#F1F334"), 0.52},
		{gradient.Hex("#EFC50F"), 0.54},
		{gradient.Hex("#E5A50F"), 0.56},
		{gradient.Hex("#E1800B"), 0.58},
		{gradient.Hex("#D30B07"), 0.60},

		{gradient.Hex("#6F8525"), 0.72},
		{gradient.Hex("#B2C85B"), 0.74},
		{gradient.Hex("#EAECB8"), 0.76},
		{gradient.Hex("#002440"), 0.78},
		{gradient.Hex("#005994"), 0.80},

		{gradient.Hex("#6F8525"), 0.92},
		{gradient.Hex("#B2C85B"), 0.94},
		{gradient.Hex("#EAECB8"), 0.96},
		{gradient.Hex("#002440"), 0.98},

		{gradient.Hex("#005994"), 1.00},
	})

	renderPalette(palette, *pfn)

	plasma := generatePlasma(w, h, *scale)

	renderPlasma(w, h, plasma, palette, *ofn)
}

func generateGradientPalette(t gradient.Table) [256]colorful.Color {
	var palette [256]colorful.Color

	for x := 0; x < 256; x++ {
		palette[x] = t.GetInterpolatedColorFor(float64(x) / 256.0)
	}

	return palette
}

func generatePalette(s float64) [256]colorful.Color {
	var palette [256]colorful.Color

	for x := 0; x < 256; x++ {
		r := (128.0 + 128*math.Sin(math.Pi*float64(x)/16.0)) / 256.0
		g := (128.0 + 128*math.Sin(math.Pi*float64(x)/128.0)) / 256.0
		b := (8.0) / 256.0

		palette[x] = colorful.Color{r, g, b}
	}

	return palette
}

func renderPalette(palette [256]colorful.Color, fn string) {
	w := 256
	h := 10

	p := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p.Set(x, y, palette[x])
		}
	}

	if file, err := os.Create(fn); err == nil {
		defer file.Close()
		if err := png.Encode(file, p); err == nil {
			open(fn)
		}
	}
}

func generatePlasma(w, h int, s float64) [][]uint8 {
	var plasma = make([][]uint8, w)

	for x := 0; x < w; x++ {
		plasma[x] = make([]uint8, h)

		i := float64(w) / 2.0
		j := float64(h) / 3.0

		for y := 0; y < h; y++ {
			c := uint8(
				(i + (j * math.Sin(float64(x)/s*1.5)) +
					i + (j * math.Sin(float64(y)/(s*2.5))) +
					i + (j * math.Sin(math.Sqrt((float64(x-w)/2)*(float64(x-w)/2)+(float64(y-h)/2)*(float64(y-h)/2))/s)) +
					i + (j * math.Sin(math.Sqrt(float64(x*x+y*y))/s))) / 4.0)

			plasma[x][y] = c
		}
	}
	return plasma
}

func renderPlasma(w, h int, plasma [][]uint8, palette [256]colorful.Color, fn string) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	if *frames > 1 {
		for s := 0; s < *frames; s++ {
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					m.Set(x, y, palette[uint8(int64(plasma[x][y])+int64(s)%255)])
				}
			}

			if file, err := os.Create(fmt.Sprintf("%03d-%s", s, fn)); err == nil {
				defer file.Close()
				png.Encode(file, m)
			}
		}
	} else {
		s := time.Now().Unix()

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				m.Set(x, y, palette[uint8(int64(plasma[x][y])+(s)%255)])
			}
		}

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
