package egg

import (
	"github.com/nsf/termbox-go"
)

// WindowWidth ...
func WindowWidth() int {
	w, _ := termbox.Size()
	return w
}

// WindowHeight ...
func WindowHeight() int {
	_, h := termbox.Size()
	return h
}

// WindowSize ...
func WindowSize() (width int, height int) {
	return termbox.Size()
}

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
