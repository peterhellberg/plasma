package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/exec"
)

var (
	width  = flag.Int("w", 512, "Width of the image")
	height = flag.Int("h", 512, "Height of the image")
	scale  = flag.Float64("s", 16.0, "Scale of the plasma")
	output = flag.String("o", "plasma.png", "Output image")
	show   = flag.Bool("show", true, "Show the generated image")
)

func main() {
	flag.Parse()

	w := *width
	h := *height

	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < *width; x++ {
		for y := 0; y < *height; y++ {
			c := uint8(128.0 + (128.0 * math.Sin(math.Sqrt(
				(float64(x-w)/2.0)*(float64(x-w)/2.0)+
					(float64(y-h)/2.0)*(float64(y-h)/2.0)) / *scale)))

			m.Set(x, y, color.RGBA{c, c, c, 0xff})
		}
	}

	if file, err := os.Create(*output); err == nil {
		defer file.Close()
		if err := png.Encode(file, m); err == nil && *show {
			exec.Command("open", *output).Run()
		}
	}
}
