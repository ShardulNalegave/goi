package goi

import (
	"goi/canvas"
	"goi/display"
	"time"
)

// GOI interface for goi struct
type GOI interface {
	init(string)
	Run()
}

// goi struct
type goi struct {
	Display display.Display
	Canvas  canvas.Canvas
}

func (g *goi) init(winTitle string) {
	g.Canvas = canvas.NewCanvas()
	g.Display = display.NewDisplay(winTitle, &g.Canvas)
}

func (g *goi) Run() {
	g.Display.Show()
	defer g.Display.Destroy()
	time.Sleep(time.Second * 2)
}

// NewGoiApp --> Creates a new goi struct
func NewGoiApp(title string) GOI {
	app := goi{
		Display: nil,
		Canvas:  nil,
	}
	app.init(title)
	return &app
}
