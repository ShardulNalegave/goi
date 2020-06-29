package colors

// Color interface for color struct
type Color interface {
	GetR() float32
	GetG() float32
	GetB() float32
}

// color struct
type color struct {
	R, G, B float32
}

func (c *color) GetR() float32 {
	return c.R
}

func (c *color) GetG() float32 {
	return c.G
}

func (c *color) GetB() float32 {
	return c.B
}

// NewColor --> Creates a new color struct
func NewColor(r float32, g float32, b float32) Color {
	return &color{
		R: r,
		G: g,
		B: b,
	}
}
