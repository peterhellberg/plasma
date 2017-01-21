package main

import (
	"image"
	"log"
	"time"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
)

const (
	width  = 1024
	height = 768
)

func main() {
	visible := true

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{width, height})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		b, err := s.NewBuffer(image.Point{width, height})
		if err != nil {
			log.Fatal(err)
		}
		defer b.Release()

		go func(p *plasma.Plasma) {
			i := 0

			for {
				time.Sleep(16 * time.Millisecond)

				if visible {
					p.Draw(b.RGBA(), i, palette.DefaultGradient)

					w.Upload(image.Point{0, 0}, b, b.Bounds())
					w.Publish()

					i += 1
				}
			}
		}(plasma.New(b.Size().X, b.Size().Y, 64.0))

		for {
			e := w.NextEvent()

			switch e := e.(type) {
			case lifecycle.Event:
				visible = !(e.From == lifecycle.StageFocused && e.To == lifecycle.StageVisible)
			case key.Event:
				if e.Code == key.CodeEscape || e.Code == key.CodeQ {
					return
				}
			}
		}
	})
}
