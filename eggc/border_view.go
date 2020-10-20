package eggc

import (
	"github.com/thomgray/egg"
)

// BorderView ...
type BorderView struct {
	*egg.View
	borderChars borderChars
	title       string
}

type borderChars struct {
	tl, t, tr, r, br, b, bl, l rune
}

// MakeBorderView ...
func MakeBorderView() *BorderView {
	bv := BorderView{
		View:        egg.MakeView(),
		borderChars: borderChars{'┏', '━', '┓', '┃', '┛', '━', '┗', '┃'},
	}
	bv.OnDraw(bv.draw)
	bv.SetViewport(func(b egg.Bounds) *egg.Bounds {
		b.Width -= 2
		b.Height -= 2
		b.X++
		b.Y++
		return &b
	})
	return &bv
}

// SetChars - set the border characters in clockwise order starting from the top-left corner and ending the left side
func (bv *BorderView) SetChars(tl, t, tr, r, br, b, bl, l rune) {
	bv.borderChars = borderChars{
		tl, t, tr, r, br, b, bl, l,
	}
}

// SetTitle - set the border tile, or set to a blank string to unset
func (bv *BorderView) SetTitle(title string) {
	bv.title = title
}

func (bv *BorderView) draw(c egg.Canvas, _ egg.State) {
	bWidth := c.Width - 1
	bHeight := c.Height - 1

	for i := 1; i < bHeight; i++ {
		c.DrawRune2(bv.borderChars.l, 0, i)
		c.DrawRune2(bv.borderChars.r, c.Width-1, i)
	}
	for i := 1; i < bWidth; i++ {
		c.DrawRune2(bv.borderChars.t, i, 0)
		c.DrawRune2(bv.borderChars.b, i, c.Height-1)
	}

	c.DrawRune2(bv.borderChars.tl, 0, 0)
	c.DrawRune2(bv.borderChars.tr, c.Width-1, 0)
	c.DrawRune2(bv.borderChars.br, c.Width-1, c.Height-1)
	c.DrawRune2(bv.borderChars.bl, 0, c.Height-1)
	if bv.title != "" {
		title := " " + bv.title + " "
		c.DrawString2(title, 3, 0)
	}
}

// GetView - get the view
func (bv *BorderView) GetView() *egg.View {
	return bv.View
}
