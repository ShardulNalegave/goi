package goi

import (
	"goi/canvas"
	"goi/display"
	"goi/vector"
)

// GOI interface for goi struct
type GOI interface {
	init(string, vector.Vector2D)
	Run()
}

// goi struct
type goi struct {
	Display display.Display
	Canvas  canvas.Canvas
}

func (g *goi) init(winTitle string, winSize vector.Vector2D) {
	g.Canvas = canvas.NewCanvas()
	g.Display = display.NewDisplay(winTitle, winSize, &g.Canvas)
}

func (g *goi) Run() {
	g.Display.SetInLoop(func() {
		//
	})
	g.Display.SetOnExit(func() {
		defer g.Display.Destroy()
	})
	g.Display.Show()
}

// NewGoiApp --> Creates a new goi struct
func NewGoiApp(title string, size vector.Vector2D) GOI {
	app := goi{
		Display: nil,
		Canvas:  nil,
	}
	app.init(title, size)
	return &app
}
