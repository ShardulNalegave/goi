package canvas

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Canvas interface for canvas struct
type Canvas interface {
	Init()
	GetGlProgram() uint32
	ClearScreen()
}

// canvas struct
type canvas struct {
	glProgram uint32
}

func (c *canvas) Init() {
	if err := gl.Init(); err != nil {
		log.Panic(err)
	}
	c.glProgram = gl.CreateProgram()
	gl.LinkProgram(c.glProgram)
}

func (c *canvas) GetGlProgram() uint32 {
	return c.glProgram
}

func (c *canvas) ClearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// NewCanvas --> Creates a new canvas struct
func NewCanvas() Canvas {
	canv := canvas{}
	return &canv
}
