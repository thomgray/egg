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
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.Writer(f))
	egg.UseTrueColor(false)
	app := egg.InitOrPanic()

	for i := 0; i < 500; i++ {
		app.AddView(makeColorView(i))
	}
	defer app.Start()

	app.OnResizeEvent(func(e *egg.ResizeEvent) {
		for _, sv := range app.GetSubViews() {
			switch ansi := sv.GetState().(type) {
			case int:
				bnds := boundsForI(ansi, e.Height)
				sv.SetBounds(bnds)
			}
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

func makeColorView(ansi int) *egg.View {
	h := egg.WindowHeight()
	v := egg.MakeView()
	bnds := boundsForI(ansi, h)
	log.Printf("x=%d y=%d", bnds.X, bnds.Y)
	v.SetBounds(bnds)
	v.SetBackground(egg.ColorAnsi(ansi))
	v.SetState(ansi)
	v.OnDraw(func(c egg.Canvas, s egg.State) {
		switch s := s.(type) {
		case int:
			c.DrawString2(fmt.Sprintf("%d", s), 0, 0)
		}
	})
	return v
}
