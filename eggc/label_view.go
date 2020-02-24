package eggc

import (
	"unicode/utf8"

	"github.com/thomgray/egg"
)

// LabelView - a view controller for a simple view containing text
type LabelView struct {
	*egg.View
	label     []byte
	alignment aligned
}

// MakeLabelView - label view constructor
func MakeLabelView() *LabelView {
	lv := LabelView{
		View:  egg.MakeView(),
		label: make([]byte, 0),
		alignment: aligned{
			h: AlignedLeft,
			v: AlignedTop,
		},
	}
	lv.OnDraw(lv.draw)
	return &lv
}

// SetLabelBytes - set the label content in bytes
func (lv *LabelView) SetLabelBytes(l []byte) {
	lv.label = l
}

// SetLabel - set the label content as a string
func (lv *LabelView) SetLabel(l string) {
	lv.label = []byte(l)
}

// SetAlignment - set the label horizontal and vertical alignment
func (lv *LabelView) SetAlignment(h AlignmentHorizontal, v AlignmentVertical) {
	lv.alignment.h = h
	lv.alignment.v = v
}

// GetView -
func (lv *LabelView) GetView() *egg.View {
	return lv.View
}

func (lv *LabelView) draw(c egg.Canvas) {
	bounds := lv.View.GetBounds()
	x := 0
	y := 0
	stringLen := utf8.RuneCountInString(string(lv.label))

	switch lv.alignment.h {
	case AlignedCenterHorizontal:
		x = (bounds.Width - stringLen) / 2
		if x < 0 {
			x = 0
		}
	case AlignedRight:
		x = bounds.Width - stringLen
	}
	switch lv.alignment.v {
	case AlignedCenterVertical:
		y = (bounds.Height / 2)
	case AlignedBottom:
		y = bounds.Height
	}
	c.DrawString2(string(lv.label), x, y)
}
