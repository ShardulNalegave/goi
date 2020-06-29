package colors

// Color interface for color struct
type Color interface {
	GetR() float32
	GetG() float32
	GetB() float32
	GetA() float32
	Normalize() [4]float32
}

// color struct
type color struct {
	R, G, B, A float32
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

func (c *color) GetA() float32 {
	return c.A
}

func (c *color) Normalize() [4]float32 {
	return [4]float32{
		c.R / 255,
		c.G / 255,
		c.B / 255,
		c.A / 255,
	}
}

// NewColor --> Creates a new color struct
func NewColor(r float32, g float32, b float32, a float32) Color {
	return &color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
