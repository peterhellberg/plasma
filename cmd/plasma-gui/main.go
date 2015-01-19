package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"
)

var (
	width  = flag.Int("w", 128, "Width of the screen")
	height = flag.Int("h", 120, "Height of the screen")
	scale  = flag.Float64("s", 8.0, "Scale of the plasma")

	count int
	p     *plasma.Plasma
	m     *image.RGBA

	pa = palette.DefaultGradient
)

func update(screen *ebiten.Image) error {
	count++

	if count%2 == 0 {
		m = p.Image(*width, *height, count, pa)
	}

	plasmaImage, err := ebiten.NewImageFromImage(m, ebiten.FilterLinear)
	if err == nil {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(*width)/2, -float64(*height)/2)
		op.GeoM.Rotate(float64(count%360) * 2 * math.Pi / 360)
		op.GeoM.Translate(float64(*width)/2, float64(*height)/2)

		if err := screen.DrawImage(plasmaImage, op); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	p = plasma.New(*width, *height, *scale)
	m = p.Image(*width, *height, count, pa)

	if err := ebiten.Run(update, *width, *height, 5, "Plasma GUI"); err != nil {
		log.Fatal(err)
	}
}
