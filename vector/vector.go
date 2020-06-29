package vector

import "math"

// Vector2D interface for vector2D struct
type Vector2D interface {
	GetX() float32
	GetY() float32
	Magnitude() float64
}

// vector2D struct
type vector2D struct {
	X, Y, Z float32
}

func (v *vector2D) Magnitude() float64 {
	return math.Sqrt(float64((v.X * v.X) + (v.Y * v.Y)))
}

func (v *vector2D) GetX() float32 {
	return v.X
}

func (v *vector2D) GetY() float32 {
	return v.Y
}

// NewVector2D --> Creates a new vector2D struct
func NewVector2D(x float32, y float32) Vector2D {
	return &vector2D{
		X: x,
		Y: y,
		// Takes a Z value because for it to be 2D it should have Z=0.0
		Z: 0.0,
	}
}
