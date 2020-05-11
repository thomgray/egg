package main

import (
	"log"
	"os"

	"github.com/thomgray/egg"
	"github.com/thomgray/egg/eggc"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	defer file.Close()
	app := egg.InitOrPanic()
	defer app.Start()

	view := eggc.MakeLabelView()
	view.SetLabel("Hello\nThere\nThis\nIs\nA\nLabel")
	view.SetBounds(egg.MakeBounds(0, 0, egg.WindowWidth()+10, egg.WindowHeight()+20))

	scroll := eggc.MakeScrollView()
	scroll.SetCanScrollHorizontally(true)
	scroll.SetBounds(egg.MakeBounds(0, 0, egg.WindowWidth(), egg.WindowHeight()))
	scroll.OnDidScroll(func() {
		log.Println("scrikked")
		app.ReDraw()
	})

	label := eggc.MakeLabelView()
	label.SetLabel("HELLLOOOOO")
	label.SetForeground(egg.ColorAnsi(100))
	label.SetBounds(egg.MakeBounds(10, 10, 20, 1))

	scroll.AddSubView(view.View)
	scroll.AddSubView(label.View)

	app.AddViewController(scroll)

	app.SetFocusedViewController(scroll)
}
