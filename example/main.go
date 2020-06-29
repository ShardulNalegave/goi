package main

import (
	goi "goi/main"
	"goi/vector"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	app := goi.NewGoiApp("Test App", vector.NewVector2D(600, 600))
	app.Run()
}
