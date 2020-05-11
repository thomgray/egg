package main

import (
	"log"
	"os"

	"github.com/thomgray/egg"
	"github.com/thomgray/egg/eggc"
)

/*
This demo showcases the use of a simple label eggc contained in a border eggc.

The origin of the label is fixed at 1,1 (inside the border view).
The size should be 2 less in either dimention than the size of the border containing the view.

Commands are implemented to move and grow/shrink the border view.
Since the origin of a view is relative to its container, the label inside the border moves as the border view does.
However, we need to manually ensure that the label view grows/shrinks with the border view.
*/

var app *egg.Application
var labeleggc *eggc.LabelView

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	defer file.Close()
	app = egg.InitOrPanic()
	defer app.Start()

	labeleggc = eggc.MakeLabelView()
	labeleggc.SetBounds(egg.MakeBounds(1, 1, 9, 1))
	labeleggc.SetBackground(egg.ColorCyan)

	bv := eggc.MakeBorderView()
	bv.SetBounds(egg.MakeBounds(10, 10, 11, 3))
	bv.AddSubView(labeleggc.View)
	bv.SetForeground(egg.ColorMagenta)
	bv.SetBackground(egg.ColorBlack)

	labeleggc.SetLabel("Hello")
	labeleggc.SetAlignment(eggc.AlignedCenterHorizontal, eggc.AlignedCenterVertical)

	bv.OnBoundsSet(func(_, new egg.Bounds) {
		log.Println("bounbds updated")
		labeleggc.View.UpdateBounds(func(b egg.Bounds) egg.Bounds {
			b.Width = new.Width - 2
			b.Height = new.Height - 2
			return b
		})
	})

	app.AddView(bv.View)
	app.OnKeyEvent(func(e *egg.KeyEvent) {
		switch e.Key {
		case egg.KeyEsc:
			app.Stop()
		case egg.KeyArrowLeft:
			bounds := bv.View.GetBounds()
			bounds.X--
			bv.SetBounds(bounds)
		case egg.KeyArrowRight:
			bounds := bv.View.GetBounds()
			bounds.X++
			bv.SetBounds(bounds)
		case egg.KeyArrowUp:
			bounds := bv.View.GetBounds()
			bounds.Y--
			bv.SetBounds(bounds)
		case egg.KeyArrowDown:
			bounds := bv.View.GetBounds()
			bounds.Y++
			bv.SetBounds(bounds)
		}

		switch e.Char {
		case 'w':
			bounds := bv.View.GetBounds()
			bv.View.SetBounds(egg.MakeBounds(bounds.X, bounds.Y, bounds.Width+1, bounds.Height))
		case 'n':
			bounds := bv.View.GetBounds()
			bv.View.SetBounds(egg.MakeBounds(bounds.X, bounds.Y, bounds.Width-1, bounds.Height))
		case 't':
			bounds := bv.View.GetBounds()
			bv.View.SetBounds(egg.MakeBounds(bounds.X, bounds.Y, bounds.Width, bounds.Height+1))
		case 's':
			bounds := bv.View.GetBounds()
			bv.View.SetBounds(egg.MakeBounds(bounds.X, bounds.Y, bounds.Width, bounds.Height-1))
		}
		app.ReDraw()
	})
}
