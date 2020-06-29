package colors

// Color interface for color struct
type Color interface{}

// color struct
type color struct{}

// NewColor --> Creates a new color struct
func NewColor() Color {
	return &color{}
}
