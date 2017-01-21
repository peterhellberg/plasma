package plasma

import (
	"image"
	"math"

	"github.com/peterhellberg/plasma/palette"
)

// Plasma contains the plasma field data
type Plasma struct {
	Field [][]uint8
}

// Image generates a image of the plasma field
func (p *Plasma) Image(w, h, s int, pa *palette.Palette) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	p.Draw(m, s, pa)

	return m
}

// Draw draws the plasma field on the provided image
func (p *Plasma) Draw(m *image.RGBA, s int, pa *palette.Palette) {
	w, h := m.Bounds().Max.X, m.Bounds().Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			m.Set(x, y, pa[uint8(int64(p.Field[x][y])+int64(s)%255)])
		}
	}
}

// New generates a new plasma field
func New(w, h int, s float64) *Plasma {
	var f = make([][]uint8, w)

	for x := 0; x < w; x++ {
		f[x] = make([]uint8, h)

		i := float64(w) / 2.0
		j := float64(h) / 3.0

		for y := 0; y < h; y++ {
			c := uint8(
				(i + (j * math.Sin(float64(x)/s*1.5)) +
					i + (j * math.Sin(float64(y)/(s*2.5))) +
					i + (j * math.Sin(math.Sqrt((float64(x-w)/2)*(float64(x-w)/2)+(float64(y-h)/2)*(float64(y-h)/2))/s)) +
					i + (j * math.Sin(math.Sqrt(float64(x*x+y*y))/s))) / 4.0)

			f[x][y] = c
		}
	}

	return &Plasma{Field: f}
}
