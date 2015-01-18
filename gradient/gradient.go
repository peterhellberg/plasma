package gradient

import "github.com/lucasb-eyer/go-colorful"

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
