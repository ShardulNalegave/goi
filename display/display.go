package display

import (
	"goi/canvas"
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Display interface for display struct
type Display interface {
	init()
	SetCanvas(*canvas.Canvas)
	Show()
}

// display struct
type display struct {
	canvas     *canvas.Canvas
	title      string
	glfwWindow *glfw.Window
}

func (d *display) init() {
	if err := glfw.Init(); err != nil {
		log.Panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	d.glfwWindow, err = glfw.CreateWindow(0, 0, d.title, nil, nil)
	if err != nil {
		log.Panic(err)
	}
}

func (d *display) SetCanvas(canv *canvas.Canvas) {
	d.canvas = canv
}

func (d *display) Show() {
	defer glfw.Terminate()
}

// NewDisplay --> Creates a new display struct
func NewDisplay(title string) Display {
	dis := display{
		title: title,
	}
	dis.init()
	return &dis
}
