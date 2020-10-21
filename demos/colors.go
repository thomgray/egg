package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/thomgray/egg"
)

func main() {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	var views []*ColorView

	if err != nil {
		panic(err)
	}
	log.SetOutput(io.Writer(f))
	egg.UseTrueColor(false)
	app := egg.InitOrPanic()

	for i := 0; i < 500; i++ {
		cv := makeColorView(i)
		views = append(views, cv)
		app.AddView(cv.View)
	}
	defer app.Start()

	app.OnResizeEvent(func(e *egg.ResizeEvent) {
		for _, sv := range views {
			bnds := boundsForI(sv.ansiCode, e.Height)
			sv.SetBounds(bnds)
		}
		app.ReDraw()
	})

	// fmt.Println("\033[1;31mhello")
}

func boundsForI(ansi, height int) egg.Bounds {
	y := ansi * 2
	x := (y / height) * 5
	y = y % height
	return egg.MakeBounds(x, y, 5, 2)
}

func makeColorView(ansi int) *ColorView {
	h := egg.WindowHeight()
	v := ColorView{}.New(ansi)
	bnds := boundsForI(ansi, h)
	v.SetBounds(bnds)
	v.SetBackground(egg.ColorAnsi(ansi))
	v.OnDraw(func(c egg.Canvas) {
		c.DrawString2(fmt.Sprintf("%d", v.ansiCode), 0, 0)
	})
	return v
}

type ColorView struct {
	*egg.View
	ansiCode int
}

func (cv ColorView) New(ansiCode int) *ColorView {
	view := egg.MakeView()

	cv.View = view
	cv.ansiCode = ansiCode

	return &cv
}
