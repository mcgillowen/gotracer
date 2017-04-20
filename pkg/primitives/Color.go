package primitives

import (
	"math"
)

// Color is the structure of a color
type Color struct {
	R, G, B, A float64
}

// Add two colors together
func (a Color) Add(b Color) Color {
	return Color{a.R + b.R, a.G + b.G, a.B + b.B, a.A + b.A}
}

// Subtract two colors
func (a Color) Subtract(b Color) Color {
	return Color{a.R - b.R, a.G - b.G, a.B - b.B, a.A - b.A}
}

// Multiply a color by a scalar
func (a Color) Multiply(s float64) Color {
	return Color{a.R * s, a.G * s, a.B * s, a.A * s}
}

// Clamp a color to a max value
func (a *Color) Clamp(m float64) {
	a.R = math.Max(0.0, math.Min(m, a.R))
	a.G = math.Max(0.0, math.Min(m, a.G))
	a.B = math.Max(0.0, math.Min(m, a.B))
	a.A = math.Max(0.0, math.Min(m, a.A))
}
