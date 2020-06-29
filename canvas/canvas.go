package canvas

// Canvas interface for canvas struct
type Canvas interface{}

// canvas struct
type canvas struct{}

// NewCanvas --> Creates a new canvas struct
func NewCanvas() Canvas {
	return &canvas{}
}
