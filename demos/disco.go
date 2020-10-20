package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/thomgray/egg/eggc"

	"github.com/thomgray/egg"
)

var app *egg.Application

func main() {
	rand.Seed(time.Now().UnixNano())
	app = egg.InitOrPanic()
	defer app.Start()
	done := make(chan bool)
	label := eggc.MakeLabelView()
	label.SetTransparent(true)
	label.SetBounds(egg.MakeBounds(0, 0, 100, 100))
	app.AddViewController(label)

	app.OnDraw(draw)

	ticker := time.NewTicker(300 * time.Millisecond)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// fireAtTheDisco()
				app.ReDraw()
				// label.SetLabel(fmt.Sprintf("Tick at %s", t))
				// app.ReDraw()
			}
		}
	}()

	app.OnKeyEvent(func(ke *egg.KeyEvent) {
		switch ke.Char {
		case 'q':
			ticker.Stop()
			done <- true
			app.Stop()
		}
	})
}

func draw(c egg.Canvas, _ egg.State) {
	w, h := egg.WindowSize()
	diam := w
	if diam > h {
		diam = h
	}
	radius := diam / 2
	midx := w / 2
	midy := h / 2

	discoBall := egg.ColorAnsi(245)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// realX := x / 2
			dx := float64(x - midx)
			dy := float64(y - midy)
			dz := math.Sqrt((dx * dx) + (dy * dy))
			realX := (midx-x)*2 + midx

			col := egg.ColorDefault
			if int(dz) <= radius {
				// within the circle!
				col = discoBall
				randomIn := rand.Intn(100)
				if randomIn > 90 {
					randcol := rand.Intn(5)
					col = egg.ColorAnsi(251 + randcol)
				}
			}

			c.DrawRune(' ', realX, y, egg.ColorDefault, col, egg.AttrNormal)
			c.DrawRune(' ', realX+1, y, egg.ColorDefault, col, egg.AttrNormal)
		}
	}
}
