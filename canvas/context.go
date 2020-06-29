package canvas

import (
	"goi/colors"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Context interface for context struct
type Context interface {
	ClearScreen()
	BackgroundColor(colors.Color)
}

// context interface
type context struct {
	canvas *canvas
}

func (c *context) ClearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (c *context) BackgroundColor(col colors.Color) {
	colorArray := col.Normalize()
	gl.ClearColor(colorArray[0], colorArray[1], colorArray[2], colorArray[3])
}

// newContext --> Creates a new context struct
func newContext(canv *canvas) Context {
	return &context{
		canvas: canv,
	}
}
