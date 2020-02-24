package main

import (
	"github.com/thomgray/egg"
)

var app *egg.Application
var labelStr = ""
var v *egg.View

func main() {
	app = egg.InitOrPanic()
	defer app.Start()

	app.SetBackground(egg.ColorMagenta)
	app.OnKeyEvent(func(e *egg.KeyEvent) {
		if e.Key == egg.KeyEsc {
			app.Stop()
		}
	})

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
	v.SetSize(e.Width-20, e.Height-20)
	app.ReDraw()
}

func handleKey(e *egg.KeyEvent) {
	if e.Key == egg.KeyEsc {
		app.Stop()
	} else {
		labelStr += "!"
		v.ReDraw()
	}
}
