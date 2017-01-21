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
	"golang.org/x/mobile/event/paint"
)

const (
	width  = 512
	height = 512
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

		go func() {
			p := plasma.New(b.Size().X, b.Size().Y, 64.0)
			i := 0

			for {
				time.Sleep(16 * time.Millisecond)

				if visible {
					i += 1
					p.Draw(b.RGBA(), i, palette.DefaultGradient)
					w.Send(paint.Event{})
				}
			}
		}()

		for {
			e := w.NextEvent()

			switch e := e.(type) {
			case lifecycle.Event:
				visible = !(e.From == lifecycle.StageFocused && e.To == lifecycle.StageVisible)
			case key.Event:
				if e.Code == key.CodeEscape || e.Code == key.CodeQ {
					return
				}
			case paint.Event:
				w.Upload(image.Point{0, 0}, b, b.Bounds())
				w.Publish()
			}
		}
	})
}
