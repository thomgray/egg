package egg

import (
	"unicode/utf8"

	"github.com/gdamore/tcell"
)

// Canvas ...
type Canvas struct {
	Bounds
	ViewPort   *Bounds
	Foreground Color
	Background Color
	Attribute  Attribute
}

// DrawRune - draw a rune at the specified relative point
func (c Canvas) DrawRune(r rune, x, y int, fg, bg Color, attr Attribute) {
	absx := c.X + x
	absy := c.Y + y

	if c.ViewPort != nil && !c.ViewPort.Contains(absx, absy) {
		return
	}

	s := tcell.StyleDefault
	s = s.Foreground(tcell.Color(fg))
	s = s.Background(tcell.Color(bg))
	_APP.screen.SetContent(absx, absy, r, nil, s)
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

// DrawAttributedString -
func (c Canvas) DrawAttributedString(s AttributedString, x, y int, fg, bg Color, attr Attribute) {
	bytes := []byte(s.str)
	for len(bytes) > 0 {
		r, w := utf8.DecodeRune(bytes)
		attsPlus := s.GetAttributesAt(x)
		bytes = bytes[w:]
		c.DrawRune(r, x, y, fg, bg, attr|attsPlus)
		x++
	}
}

// DrawCursor - draw the cursor at the specified x/y.
func (c Canvas) DrawCursor(x, y int) {
	_APP.screen.ShowCursor(c.X+x, c.Y+y)
}

func makeCanvas(bounds Bounds, fg, bg Color, atts Attribute) Canvas {
	return Canvas{
		Bounds:     bounds,
		Foreground: fg,
		Background: bg,
		Attribute:  atts,
	}
}

func makeCanvasWithViewPort(bounds Bounds, viewport *Bounds, fg, bg Color, atts Attribute) Canvas {
	return Canvas{
		Bounds:     bounds,
		ViewPort:   viewport,
		Foreground: fg,
		Background: bg,
		Attribute:  atts,
	}
}
