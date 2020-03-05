package egg

import (
	"github.com/nsf/termbox-go"
)

// SetCursorVisible ...
func SetCursorVisible(visible bool) {
	termbox.HideCursor()
}

func clearBounds(bounds Bounds, fg, bg Color) {
	for x := bounds.X; x < bounds.X+bounds.Width; x++ {
		for y := bounds.Y; y < bounds.Y+bounds.Height; y++ {
			termbox.SetCell(x, y, '\000', termbox.Attribute(fg), termbox.Attribute(bg))
		}
	}
}

func absoluteBounds(view *View) Bounds {
	bounds := view.GetBounds()
	v := view.getSuperview()

	for v != nil {
		vBounds := v.GetBounds()
		bounds.X += vBounds.X
		bounds.Y += vBounds.Y

		v = v.getSuperview()
	}
	return bounds
}

func absoluteBounds2(bounds Bounds, view *View) Bounds {
	v := view.getSuperview()

	for v != nil {
		vBounds := v.GetBounds()
		bounds.X += vBounds.X
		bounds.Y += vBounds.Y

		v = v.getSuperview()
	}
	return bounds
}

func intersectionBounds(bounds1, bounds2 Bounds) Bounds {
	maxX := maxInts(bounds1.X, bounds2.Y)
	maxY := maxInts(bounds1.Y, bounds2.Y)

	minXX := minInts(bounds1.X+bounds1.Width, bounds2.X+bounds2.Width)
	minYY := minInts(bounds1.Y+bounds1.Height, bounds2.Y+bounds2.Height)

	w := 0
	h := 0

	if minXX >= maxX {
		w = minXX - maxX
	}
	if minYY >= maxY {
		h = minYY - maxY
	}
	return Bounds{
		Origin{X: maxX, Y: maxY},
		Size{Width: w, Height: h},
	}
}

func minInts(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInts(x, y int) int {
	if x > y {
		return x
	}
	return y
}
