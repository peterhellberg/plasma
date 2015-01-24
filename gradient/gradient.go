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
