package canvas

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Canvas interface for canvas struct
type Canvas interface {
	init()
	GetGlProgram() uint32
}

// canvas struct
type canvas struct {
	glProgram uint32
}

func (c *canvas) init() {
	if err := gl.Init(); err != nil {
		log.Panic(err)
	}
	c.glProgram = gl.CreateProgram()
}

func (c *canvas) GetGlProgram() uint32 {
	return c.glProgram
}

// NewCanvas --> Creates a new canvas struct
func NewCanvas() Canvas {
	canv := canvas{}
	canv.init()
	return &canv
}
