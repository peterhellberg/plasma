package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/peterhellberg/plasma"
	"github.com/peterhellberg/plasma/palette"
)

const (
	gamepadID = 0
)

var (
	width  = flag.Int("w", 256, "Width of the screen")
	height = flag.Int("h", 240, "Height of the screen")
	scale  = flag.Int("s", 2, "Scaling factor")
	size   = flag.Float64("size", 17.0, "Size of the plasma")

	count int

	x float64
	y float64

	p *plasma.Plasma
	m *image.RGBA

	pa = palette.DefaultGradient
)

func pressed(buttonID ebiten.GamepadButton) bool {
	return ebiten.IsGamepadButtonPressed(gamepadID, buttonID)
}

func update(screen *ebiten.Image) error {
	count++

	if count%2 == 0 {
		m = p.Image(*width, *height, count, pa)
	}

	plasmaImage, err := ebiten.NewImageFromImage(m, ebiten.FilterNearest)
	if err == nil {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(*width)/2, -float64(*height)/2)
		op.GeoM.Rotate(float64(count%360) * 2 * math.Pi / 360)

		op.GeoM.Scale(
			1.0+ebiten.GamepadAxis(gamepadID, 2),
			1.0+ebiten.GamepadAxis(gamepadID, 3),
		)

		x += ebiten.GamepadAxis(gamepadID, 0)
		y += ebiten.GamepadAxis(gamepadID, 1)

		// Up
		if pressed(14) {
			y--
		}

		// Right
		if pressed(15) {
			x++
		}

		// Down
		if pressed(16) {
			y++
		}

		// Left
		if pressed(17) {
			x--
		}

		// X for exit
		if pressed(1) {
			os.Exit(0)
		}

		op.GeoM.Translate(x+float64(*width)/2, y+float64(*height)/2)

		if err := screen.DrawImage(plasmaImage, op); err != nil {
			return err
		}
	}

	pressedButtons := []string{}
	maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(gamepadID))
	for b := ebiten.GamepadButton(gamepadID); b < maxButton; b++ {
		if ebiten.IsGamepadButtonPressed(gamepadID, b) {
			pressedButtons = append(pressedButtons, strconv.Itoa(int(b)))
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"\n %s\n FPS: %.2f\n POS: %.2fx%.2f\n BTN: %s",
		time.Now(),
		ebiten.CurrentFPS(),
		x, y,
		strings.Join(pressedButtons, ", "),
	))

	return nil
}

func main() {
	flag.Parse()

	p = plasma.New(*width, *height, *size)
	m = p.Image(*width, *height, count, pa)

	if err := ebiten.Run(update, *width, *height, *scale, "Plasma Joystick GUI"); err != nil {
		log.Fatal(err)
	}
}
