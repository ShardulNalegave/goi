package main

import (
	"goi/canvas"
	"goi/colors"
	goi "goi/main"
	"goi/vector"
	"log"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	app := goi.NewGoiApp("Test App", vector.NewVector2D(600, 600))

	(*app.GetCanvas()).OnSetup(func(ctx canvas.Context) {
		log.Println("Canvas Setup....")
		ctx.BackgroundColor(colors.NewColor(50, 50, 50, 1))
	})

	(*app.GetCanvas()).OnDraw(func(ctx canvas.Context) {
		//
	})

	app.Run()
}
