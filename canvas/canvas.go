package canvas

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Canvas interface for canvas struct
type Canvas interface {
	Setup()
	Draw()
	ClearScreen()
	GetGlProgram() uint32
	OnSetup(func(Context))
	OnDraw(func(Context))
}

// canvas struct
type canvas struct {
	glProgram uint32
	_onSetup  func(Context)
	_onDraw   func(Context)
}

func (c *canvas) Setup() {
	if err := gl.Init(); err != nil {
		log.Panic(err)
	}
	c.glProgram = gl.CreateProgram()
	gl.LinkProgram(c.glProgram)
	c._onSetup(newContext(c))
}

func (c *canvas) Draw() {
	c._onDraw(newContext(c))
}

func (c *canvas) ClearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (c *canvas) GetGlProgram() uint32 {
	return c.glProgram
}

func (c *canvas) OnSetup(handler func(Context)) {
	c._onSetup = func(ctx Context) {
		handler(ctx)
	}
}

func (c *canvas) OnDraw(handler func(Context)) {
	c._onDraw = func(ctx Context) {
		handler(ctx)
	}
}

// NewCanvas --> Creates a new canvas struct
func NewCanvas() Canvas {
	canv := canvas{
		_onSetup: func(ctx Context) {},
		_onDraw:  func(ctx Context) {},
	}
	return &canv
}
