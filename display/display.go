package display

import "goi/canvas"

// Display interface for display struct
type Display interface {
	init()
	SetCanvas(*canvas.Canvas)
}

// display struct
type display struct {
	canvas *canvas.Canvas
}

func (d *display) init() {
	//
}

func (d *display) SetCanvas(canv *canvas.Canvas) {
	d.canvas = canv
}

// NewDisplay --> Creates a new display struct
func NewDisplay() Display {
	return &display{}
}
