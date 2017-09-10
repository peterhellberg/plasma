package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"
)

var (
	width  = flag.Int("w", 256, "Width of the screen")
	height = flag.Int("h", 256, "Height of the screen")
	scale  = flag.Float64("s", 2, "Scaling factor")
	size   = flag.Float64("size", 12.0, "Size of the plasma")

	p *plasma.Plasma
	m *image.RGBA

	count int

	pa = palette.DefaultGradient
)

func update(screen *ebiten.Image) error {
	count++

	if count%2 == 0 {
		m = p.Image(*width, *height, count, pa)
	}

	screen.ReplacePixels(m.Pix)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("\n %s\n FPS: %.2f", time.Now(), ebiten.CurrentFPS()))

	return nil
}

func main() {
	flag.Parse()

	p = plasma.New(*width, *height, *size)
	m = p.Image(*width, *height, count, pa)

	if err := ebiten.Run(update, *width, *height, *scale, "Plasma Pixels GUI"); err != nil {
		log.Fatal(err)
	}
}
