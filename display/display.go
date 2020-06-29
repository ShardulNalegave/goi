package display

import (
	"goi/canvas"
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Display interface for display struct
type Display interface {
	init()
	Show()
	SetVisible(bool)
	Destroy()
	SetInLoop(func())
	SetOnExit(func())
}

// display struct
type display struct {
	canvas     *canvas.Canvas
	title      string
	glfwWindow *glfw.Window
	_inLoop    func()
	_onExit    func()
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
}

func (d *display) Show() {
	var err error
	d.glfwWindow, err = glfw.CreateWindow(600.0, 600.0, d.title, nil, nil)
	if err != nil {
		log.Panic(err)
	}
	d.glfwWindow.MakeContextCurrent()

	for !d.glfwWindow.ShouldClose() {
		d._inLoop()
		glfw.PollEvents()
		d.glfwWindow.SwapBuffers()
	}
	d._onExit()
}

func (d *display) SetVisible(shouldShow bool) {
	if shouldShow {
		d.glfwWindow.Show()
	} else {
		d.glfwWindow.Hide()
	}
}

func (d *display) Destroy() {
	d.glfwWindow.Destroy()
	glfw.Terminate()
}

func (d *display) SetInLoop(handler func()) {
	d._inLoop = handler
}

func (d *display) SetOnExit(handler func()) {
	d._onExit = handler
}

// NewDisplay --> Creates a new display struct
func NewDisplay(title string, canv *canvas.Canvas) Display {
	dis := display{
		title:   title,
		canvas:  canv,
		_inLoop: func() {},
		_onExit: func() {},
	}
	dis.init()
	return &dis
}
