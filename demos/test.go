package main

import (
	"github.com/thomgray/egg"
	"github.com/thomgray/egg/eggc"
)

func main() {
	app := egg.InitOrPanic()
	defer app.Start()

	labelView := eggc.MakeLabelView()
	labelView.SetID("label")
	labelView.SetLabel("Hello")
	labelView.SetBounds(egg.MakeBounds(1, 1, 8, 8))
	labelView.SetBackground(egg.ColorRed)

	bv := eggc.MakeBorderView()
	bv.SetBounds(egg.MakeBounds(1, 1, 10, 10))
	bv.AddSubView(labelView.View)

	app.AddViewController(bv)
	// app.SetBackground(egg.ColorGreen)
}
