package eggc

import (
	"strings"
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

// GetLabel -
func (lv *LabelView) GetLabel() string {
	return string(lv.label)
}

func (lv *LabelView) draw(c egg.Canvas) {
	bounds := lv.View.GetBounds()
	labelStr := string(lv.label)
	labelLines := strings.Split(labelStr, "\n")

	stringLines := len(labelLines)
	y := 0
	switch lv.alignment.v {
	case AlignedCenterVertical:
		y = (bounds.Height + stringLines/2)
	case AlignedBottom:
		y = bounds.Height - stringLines
	}

	for _, s := range labelLines {
		x := 0
		lineLength := utf8.RuneCountInString(s)
		switch lv.alignment.h {
		case AlignedCenterHorizontal:
			x = (bounds.Width + lineLength) / 2
		case AlignedRight:
			x = bounds.Width - lineLength
		}
		c.DrawString2(s, x, y)
		y++
	}
}
