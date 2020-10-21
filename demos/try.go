package main

import (
	"fmt"

	"github.com/thomgray/egg"
)

func main() {
	app, err := egg.Init()
	// var msg string = "Hello"
	var count int = 0
	if err != nil {
		panic(err)
	}

	app.OnDraw(func(c egg.Canvas) {
		c.DrawString(fmt.Sprintf("Hello %d", count), 10, 5, egg.ColorRed, egg.ColorWhite, egg.AttrBold)
	})

	app.OnKeyEvent(func(ke *egg.KeyEvent) {
		count++
		app.ReDraw()
	})
	defer app.Start()
}
