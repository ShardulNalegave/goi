package goi

import (
	"goi/canvas"
	"goi/display"
)

// GOI interface for goi struct
type GOI interface {
	init()
	Run()
}

// goi struct
type goi struct {
	Display display.Display
	Canvas  canvas.Canvas
}

func (g *goi) init() {
	g.Canvas = canvas.NewCanvas()
	g.Display = display.NewDisplay()
	g.Display.SetCanvas(&g.Canvas)
}

func (g *goi) Run() {
	g.Display.Show()
}

// NewGoiApp --> Creates a new goi struct
func NewGoiApp() GOI {
	return &goi{
		Display: nil,
		Canvas:  nil,
	}
}
