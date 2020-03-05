package main

import (
	"github.com/thomgray/egg"
	"github.com/thomgray/egg/eggc"
)

func main() {
	app := egg.InitOrPanic()
	defer app.Start()

	view := eggc.MakeLabelView()
	view.SetLabel("Hello\nThere\nThis\nIs\nA\nLabel")
	view.SetBounds(egg.MakeBounds(0, 0, egg.WindowWidth()+10, egg.WindowHeight()+20))

	scroll := eggc.MakeScrollView()
	scroll.SetCanScrollHorizontally(true)
	scroll.SetBounds(egg.MakeBounds(0, 0, egg.WindowWidth(), egg.WindowHeight()))
	scroll.OnDidScroll(func() {
		app.ReDraw()
	})
	// scroll.SetCanScrollHorizontally(false)

	label := eggc.MakeLabelView()
	label.SetLabel("HELLLOOOOO")
	label.SetForeground(egg.ColorAnsi(100))
	label.SetBounds(egg.MakeBounds(10, 10, 20, 1))

	scroll.AddSubView(view.View)
	scroll.AddSubView(label.View)

	app.AddViewController(scroll)

	app.SetFocusedViewController(scroll)
}
