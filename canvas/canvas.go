package canvas

import (
	"goi/colors"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Canvas interface for canvas struct
type Canvas interface {
	Init()
	GetGlProgram() uint32
	ClearScreen()
	BackgroundColor(colors.Color)
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

func (c *canvas) BackgroundColor(col colors.Color) {
	colorArray := col.Normalize()
	gl.ClearColor(colorArray[0], colorArray[1], colorArray[2], colorArray[3])
}

// NewCanvas --> Creates a new canvas struct
func NewCanvas() Canvas {
	canv := canvas{}
	return &canv
}
