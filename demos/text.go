package main

import (
	"github.com/thomgray/egg"
	"github.com/thomgray/egg/eggc"
)

func main() {
	app := egg.InitOrPanic()
	defer app.Start()

	textView := eggc.MakeTextView()
	textView.SetBounds(egg.MakeBounds(0, 0, 80, 20))
	textView.SetTextContentString("Hello there")

	textView.GainFocus()

	app.AddViewController(textView)
	app.ReDraw()
}
