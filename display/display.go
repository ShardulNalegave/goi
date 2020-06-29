package display

import (
	"goi/canvas"
	"goi/vector"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
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
	ChangeTitle(string)
}

// display struct
type display struct {
	canvas     *canvas.Canvas
	title      string
	size       vector.Vector2D
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
	d.glfwWindow, err = glfw.CreateWindow(int(d.size.GetX()), int(d.size.GetY()), d.title, nil, nil)
	if err != nil {
		log.Panic(err)
	}
	d.glfwWindow.MakeContextCurrent()
	(*d.canvas).Setup()

	for !d.glfwWindow.ShouldClose() {
		(*d.canvas).ClearScreen()
		(*d.canvas).Draw()
		gl.UseProgram((*d.canvas).GetGlProgram())
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

func (d *display) ChangeTitle(title string) {
	d.title = title
	d.glfwWindow.SetTitle(title)
}

// NewDisplay --> Creates a new display struct
func NewDisplay(title string, size vector.Vector2D, canv *canvas.Canvas) Display {
	dis := display{
		title:   title,
		size:    size,
		canvas:  canv,
		_inLoop: func() {},
		_onExit: func() {},
	}
	dis.init()
	return &dis
}
