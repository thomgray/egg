package main

import (
	"log"
	"os"

	"github.com/thomgray/egg"
)

var app *egg.Application
var labelStr = ""
var v *egg.View

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	defer file.Close()
	log.Println("Starting...")
	app = egg.InitOrPanic()
	defer app.Start()

	app.SetBackground(egg.ColorMagenta)

	v = egg.MakeView()
	v.SetForeground(egg.ColorWhite)
	v.SetBackground(egg.ColorRed)
	v.SetAttributes(egg.AttrBold, egg.AttrUnderline)
	v.SetBounds(egg.MakeBounds(1, 1, egg.WindowWidth()-20, egg.WindowHeight()-20))
	v.OnDraw(drawView)
	app.AddView(v)

	app.OnKeyEvent(handleKey)
	app.OnResizeEvent(handleResize)
}

func drawView(c egg.Canvas) {
	c.DrawString(labelStr, 0, 0, c.Foreground, c.Background, c.Attribute)
}

func handleResize(e *egg.ResizeEvent) {
	v.UpdateBounds(func(b egg.Bounds) egg.Bounds {
		b.Width = e.Width - 20
		b.Height = e.Height - 20
		return b
	})
	app.ReDraw()
}

func handleKey(e *egg.KeyEvent) {
	if e.Key == egg.KeyEsc {
		app.Stop()
	} else {
		log.Println("??????????")
		labelStr += "!"
		v.ReDraw()
	}
}
