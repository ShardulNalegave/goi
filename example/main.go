package main

import (
	goi "goi/main"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	app := goi.NewGoiApp("Test App")
	app.Run()
}
