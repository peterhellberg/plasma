package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, 1024, 512),
		VSync:       true,
		Resizable:   false,
		Undecorated: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	var s *pixel.Sprite

	p := plasmaPicture(512, 256, 10.0, 0)
	s = pixel.NewSprite(p, p.Bounds())

	centerMatrix := pixel.IM.Moved(win.Bounds().Center()).Scaled(
		win.Bounds().Center(), 2,
	)

	size := 12.0

	go func() {
		c := time.Tick(32 * time.Millisecond)

		var i int

		for range c {
			i++

			p := plasmaPicture(256, 128, size, i)

			s = pixel.NewSprite(p, p.Bounds())
		}
	}()

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		s.Draw(win, centerMatrix)

		if win.Pressed(pixelgl.KeyUp) {
			size += 0.2
		}

		if win.Pressed(pixelgl.KeyDown) {
			size -= 0.2
		}

		win.Update()
	}
}

func plasmaPicture(w, h int, s float64, i int) *pixel.PictureData {
	return pixel.PictureDataFromImage(plasma.New(w, h, s).
		Image(w, h, i, palette.MaterialDesign700))
}

func main() {
	pixelgl.Run(run)
}
