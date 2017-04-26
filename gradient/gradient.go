package gradient

import (
	"github.com/lucasb-eyer/go-colorful"
)

// Default is the default gradient
var Default = Table{
	{Hex("#005994"), 0.00},

	{Hex("#6F8525"), 0.12},
	{Hex("#B2C85B"), 0.14},
	{Hex("#EAECB8"), 0.16},
	{Hex("#002440"), 0.18},
	{Hex("#005994"), 0.20},

	{Hex("#6F8525"), 0.32},
	{Hex("#B2C85B"), 0.34},
	{Hex("#EAECB8"), 0.36},
	{Hex("#002440"), 0.38},
	{Hex("#005994"), 0.40},

	{Hex("#F1F334"), 0.52},
	{Hex("#6B0103"), 0.54},
	{Hex("#F03C02"), 0.56},
	{Hex("#1C0113"), 0.58},
	{Hex("#D30B07"), 0.60},

	{Hex("#002440"), 0.72},
	{Hex("#005994"), 0.74},
	{Hex("#004675"), 0.76},
	{Hex("#002440"), 0.78},
	{Hex("#005994"), 0.80},

	{Hex("#6F8525"), 0.92},
	{Hex("#B2C85B"), 0.94},
	{Hex("#EAECB8"), 0.96},
	{Hex("#002440"), 0.98},

	{Hex("#005994"), 1.00},
}

// RainbowDash is based on http://www.color-hex.com/color-palette/807
var RainbowDash = Table{
	{Hex("#ee4035"), 0.00},
	{Hex("#f37736"), 0.25},
	{Hex("#fdf498"), 0.50},
	{Hex("#7bc043"), 0.75},
	{Hex("#0392cf"), 1.00},
}

// MaterialDesign500 is based on https://www.materialpalette.com/colors (A500)
var MaterialDesign500 = Table{
	{Hex("#f44336"), 0.0000},
	{Hex("#e91e63"), 0.0666},
	{Hex("#9c27b0"), 0.1333},
	{Hex("#673ab7"), 0.1999},
	{Hex("#3f51b5"), 0.2666},
	{Hex("#2196f3"), 0.3333},
	{Hex("#03a9f4"), 0.3999},
	{Hex("#00bcd4"), 0.4666},
	{Hex("#009688"), 0.5332},
	{Hex("#4caf50"), 0.5999},
	{Hex("#8bc34a"), 0.6666},
	{Hex("#cddc39"), 0.7326},
	{Hex("#ffeb3b"), 0.7992},
	{Hex("#ffc107"), 0.8658},
	{Hex("#ff9800"), 0.9324},
	{Hex("#ff5722"), 1.0000},
}

// MaterialDesign700 is based on https://www.materialpalette.com/colors (A700)
var MaterialDesign700 = Table{
	{Hex("#d50000"), 0.0000},
	{Hex("#c51162"), 0.0666},
	{Hex("#aa00ff"), 0.1333},
	{Hex("#6200ea"), 0.1999},
	{Hex("#304ffe"), 0.2666},
	{Hex("#2962ff"), 0.3333},
	{Hex("#0091ea"), 0.3999},
	{Hex("#00b8d4"), 0.4666},
	{Hex("#00bfa5"), 0.5332},
	{Hex("#00c853"), 0.5999},
	{Hex("#64dd17"), 0.6666},
	{Hex("#aeea00"), 0.7326},
	{Hex("#ffd600"), 0.7992},
	{Hex("#ffab00"), 0.8658},
	{Hex("#ff6d00"), 0.9324},
	{Hex("#dd2c00"), 1.0000},
}

// Table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type Table []struct {
	Col colorful.Color
	Pos float64
}

// GetInterpolatedColorFor returns a HCL-blend between the two colors around `n`.
// This is the meat of the gradient computation.
// Note: gradient keypoints must be sorted
func (t Table) GetInterpolatedColorFor(n float64) colorful.Color {
	for i := 0; i < len(t)-1; i++ {
		c1 := t[i]
		c2 := t[i+1]
		if c1.Pos <= n && n <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			n := (n - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, n).Clamped()
		}
	}
	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return t[len(t)-1].Col
}

// Hex parses hex color string
func Hex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}
