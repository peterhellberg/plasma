package palette

import (
	"image"
	"math"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/peterhellberg/plasma/gradient"
)

// DefaultGradient contains a palette based on the default gradient
var DefaultGradient = FromGradient(gradient.Default)

// A Palette contains 256 colorful colors
type Palette [256]colorful.Color

// Image returns an image representation of the palette
func (p *Palette) Image() *image.RGBA {
	w := 256
	h := 10
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			m.Set(x, y, p[x])
		}
	}

	return m
}

// FromGradient generates a palette based on the given gradient table
func FromGradient(t gradient.Table) *Palette {
	var p Palette

	for x := 0; x < 256; x++ {
		p[x] = t.GetInterpolatedColorFor(float64(x) / 256.0)
	}

	return &p
}

// Default returns the default palette
func Default(s float64) *Palette {
	var p Palette

	for x := 0; x < 256; x++ {
		r := (128.0 + 128*math.Sin(math.Pi*float64(x)/16.0)) / 256.0
		g := (128.0 + 128*math.Sin(math.Pi*float64(x)/128.0)) / 256.0
		b := (8.0) / 256.0

		p[x] = colorful.Color{r, g, b}
	}

	return &p
}
