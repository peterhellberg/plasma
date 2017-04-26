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

	size := 12.0

	p := plasmaPicture(256, 128, size, 0)
	s = pixel.NewSprite(p, p.Bounds())

	centerMatrix := pixel.IM.Moved(win.Bounds().Center()).Scaled(
		win.Bounds().Center(), 4,
	)

	go func() {
		c := time.Tick(32 * time.Millisecond)

		var i int

		for range c {
			i++

			p := plasmaPicture(256, 128, size, i)

			s = pixel.NewSprite(p, p.Bounds())
			s.SetMatrix(centerMatrix)
		}
	}()

	s.SetMatrix(centerMatrix)

	for !win.Closed() {
		win.Update()

		s.Draw(win)

		if win.Pressed(pixelgl.KeyUp) {
			size += 0.2
		}

		if win.Pressed(pixelgl.KeyDown) {
			size -= 0.2
		}

		if win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ) {
			return
		}
	}
}

func plasmaPicture(w, h int, s float64, i int) *pixel.PictureData {
	return pixel.PictureDataFromImage(plasma.New(w, h, s).
		Image(w, h, i, palette.DefaultGradient))
}

func main() {
	pixelgl.Run(run)
}
