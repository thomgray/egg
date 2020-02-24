package egg

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

// Canvas ...
type Canvas struct {
	Bounds
	Foreground Color
	Background Color
	Attribute  Attribute
}

// DrawRune - draw a rune at the specified relative point
func (c Canvas) DrawRune(r rune, x, y int, fg, bg Color, attr Attribute) {
	fgAtts := termbox.Attribute(fg) | termbox.Attribute(attr)
	termbox.SetCell(c.X+x, c.Y+y, r, fgAtts, termbox.Attribute(bg))
}

// DrawRune2 - draw rune at the specified point with the canvas attributes
func (c Canvas) DrawRune2(r rune, x, y int) {
	c.DrawRune(r, x, y, c.Foreground, c.Background, c.Attribute)
}

// DrawString - draw a string at the specified relative point
func (c Canvas) DrawString(s string, x, y int, fg, bg Color, attr Attribute) {
	bytes := []byte(s)
	for len(bytes) > 0 {
		r, w := utf8.DecodeRune(bytes)
		bytes = bytes[w:]
		c.DrawRune(r, x, y, fg, bg, attr)
		x++
	}
}

// DrawString2 - draw a string at the specified relative point with the canvas attributes
func (c Canvas) DrawString2(s string, x, y int) {
	c.DrawString(s, x, y, c.Foreground, c.Background, c.Attribute)
}

// DrawCursor - draw the cursor at the specified x/y.
func (c Canvas) DrawCursor(x, y int) {
	termbox.SetCursor(c.X+x, c.Y+y)
}

func makeCanvas(bounds Bounds, fg, bg Color, atts Attribute) Canvas {
	return Canvas{
		Bounds:     bounds,
		Foreground: fg,
		Background: bg,
		Attribute:  atts,
	}
}
