package eggc

import "github.com/thomgray/egg"

// BorderView ...
type BorderView struct {
	*egg.View
	borderChars borderChars
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
	return &bv
}

// SetChars - set the border characters in clockwise order starting from the top-left corner and ending the left side
func (bv *BorderView) SetChars(tl, t, tr, r, br, b, bl, l rune) {
	bv.borderChars = borderChars{
		tl, t, tr, r, br, b, bl, l,
	}
}

func (bv *BorderView) draw(c egg.Canvas) {
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
}

func (bv *BorderView) GetView() *egg.View {
	return bv.View
}
