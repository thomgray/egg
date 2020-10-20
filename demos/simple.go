package main

import (
	"fmt"

	"github.com/thomgray/egg"
)

func main() {
	app := egg.InitOrPanic()

	defer app.Start()

	app.SetState("Hello")
	app.OnDraw(func(c egg.Canvas, state egg.State) {
		switch state := state.(type) {
		case string:
			c.DrawString2(state, 0, 0)
		}

	})

	app.OnKeyEvent(func(ke *egg.KeyEvent) {
		app.SetState(fmt.Sprintf("Key pressed [%c]", ke.Char))
		app.ReDraw()
	})
}
