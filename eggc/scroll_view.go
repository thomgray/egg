package eggc

import (
	"math"
	"strings"

	"github.com/thomgray/egg"
)

type attributedRune struct {
	char rune
	fg   egg.Color
	bg   egg.Color
	att  egg.Attribute
}

// ScrollView ...
type ScrollView struct {
	*egg.View
	scrollVertical      bool
	scrollHorizontal    bool
	drawBarHorizontal   bool
	drawBarVertical     bool
	horizontalBlockChar attributedRune
	horizontalSpaceChar attributedRune
	verticalBlockChar   attributedRune
	verticalSpaceChar   attributedRune
	didScroll           func()
	batcher             *Batcher
}

// MakeScrollView - create a scroll view
func MakeScrollView() *ScrollView {
	sv := ScrollView{
		View:              egg.MakeView(),
		scrollVertical:    true,
		scrollHorizontal:  true,
		drawBarHorizontal: true,
		drawBarVertical:   true,
		verticalBlockChar: attributedRune{
			char: '█', bg: egg.ColorDefault, fg: egg.ColorDefault, att: egg.Attribute(0),
		},
		verticalSpaceChar: attributedRune{
			char: ' ', bg: egg.ColorDefault, fg: egg.ColorDefault, att: egg.Attribute(0),
		},
		horizontalBlockChar: attributedRune{
			char: '▄', bg: egg.ColorDefault, fg: egg.ColorDefault, att: egg.Attribute(0),
		},
		horizontalSpaceChar: attributedRune{
			char: ' ', bg: egg.ColorDefault, fg: egg.ColorDefault, att: egg.Attribute(0),
		},
	}
	sv.OnKeyEvent(sv.keyHandler)
	sv.OnDraw(sv.draw)
	sv.SetViewport(sv.viewport)
	sv.batcher = MakeBatcher()
	sv.batcher.Receive(sv.batchReceive)
	// sv.OnKeyEvent(sv.batcher.AsKeyEventHandler())
	sv.OnKeyEvent(sv.keyHandler)

	return &sv
}

// GetView ...
func (sv *ScrollView) GetView() *egg.View {
	return sv.View
}

// SetCanScrollVertically ...
func (sv *ScrollView) SetCanScrollVertically(v bool) {
	sv.scrollVertical = v
}

// SetCanScrollHorizontally ...
func (sv *ScrollView) SetCanScrollHorizontally(v bool) {
	sv.scrollHorizontal = v
}

// SetVerticalScrollBarFormat -
func (sv *ScrollView) SetVerticalScrollBarFormat(block, space rune) {
	sv.verticalBlockChar.char = block
	sv.verticalSpaceChar.char = space
}

// SetHorizontalScrollBarFormat -
func (sv *ScrollView) SetHorizontalScrollBarFormat(block, space rune) {
	sv.horizontalBlockChar.char = block
	sv.horizontalSpaceChar.char = space
}

func (sv *ScrollView) viewport(b egg.Bounds) *egg.Bounds {
	if sv.drawBarHorizontal && sv.scrollHorizontal {
		b.Height--
	}
	if sv.drawBarVertical && sv.scrollVertical {
		b.Width--
	}
	return &b
}

func (sv *ScrollView) draw(c egg.Canvas) {
	bigBounds := sv.subViewUnionBounds()
	if sv.scrollVertical && sv.drawBarVertical {
		sv.drawVerticalScrollBar(c, bigBounds)
	}
	if sv.scrollHorizontal && sv.drawBarHorizontal {
		sv.drawHorizontalScrollBar(c, bigBounds)
	}
}

func (sv *ScrollView) drawHorizontalScrollBar(c egg.Canvas, unionBounds egg.Bounds) {
	svBounds := sv.GetBounds()
	fullScrollLength := unionBounds.Width - svBounds.Width
	if fullScrollLength <= 0 {
		return
	}
	scrollXOff := unionBounds.X * -1
	percentOff := float64(scrollXOff) / float64(fullScrollLength)
	barWidth := svBounds.Width / fullScrollLength
	if barWidth < 1 {
		barWidth = 1
	}
	tokenOff := int((float64(svBounds.Width - barWidth - 1)) * percentOff)
	block := strings.Repeat(string(sv.horizontalBlockChar.char), barWidth)

	c.DrawString2(block, tokenOff, svBounds.Height-1)
}

func (sv *ScrollView) drawVerticalScrollBar(c egg.Canvas, unionBounds egg.Bounds) {
	svBounds := sv.GetBounds()
	fullScrollLength := float64(unionBounds.Height) - float64(svBounds.Height)
	if fullScrollLength <= 0 {
		return
	}
	scrollYOff := float64(unionBounds.Y * -1)
	percentOff := float64(scrollYOff) / float64(fullScrollLength)
	barHeight := int(float64(svBounds.Height) / fullScrollLength)
	if barHeight < 1 {
		barHeight = 1
	}
	tokenOff := int((float64(svBounds.Height - barHeight)) * percentOff)
	for i := tokenOff; i < tokenOff+barHeight; i++ {
		c.DrawRune2(sv.verticalBlockChar.char, svBounds.Width-1, i)
	}
}

func (sv *ScrollView) keyHandler(ke *egg.KeyEvent) {
	switch ke.Key {
	case egg.KeyLeft:
		sv.ScrollLeft(1)
	case egg.KeyRight:
		sv.ScrollRight(1)
	case egg.KeyUp:
		sv.ScrollUp(1)
	case egg.KeyDown:
		sv.ScrollDown(1)
	}
}

func (sv *ScrollView) batchReceive(e []interface{}) {
	h := 0
	v := 0
	for _, i := range e {
		ke := i.(*egg.KeyEvent)
		switch ke.Key {
		case egg.KeyLeft:
			h--
		case egg.KeyRight:
			h++
		case egg.KeyUp:
			v++
		case egg.KeyDown:
			v--
		}
	}

	if h > 0 {
		sv.ScrollRight(h)
	} else if h < 0 {
		sv.ScrollLeft(h * -1)
	}

	if v > 0 {
		sv.ScrollUp(v)
	} else if v < 0 {
		sv.ScrollDown(v * -1)
	}
}

// ScrollLeft ...
func (sv *ScrollView) ScrollLeft(amount int) {
	if !sv.scrollHorizontal {
		return
	}
	subviews := sv.GetSubViews()
	if len(subviews) == 0 {
		return
	}
	union := sv.subViewUnionBounds()
	maxScroll := union.X * -1
	toScroll := amount
	if maxScroll < toScroll {
		toScroll = maxScroll
	}

	if maxScroll <= 0 {
		return
	}
	for _, subv := range sv.GetSubViews() {
		subv.UpdateBounds(func(b egg.Bounds) egg.Bounds {
			b.X += toScroll
			return b
		})
	}
	if sv.didScroll != nil {
		sv.didScroll()
	}
}

// ScrollRight ...
func (sv *ScrollView) ScrollRight(amount int) {
	if !sv.scrollHorizontal {
		return
	}
	thisBounds := sv.GetBounds()
	subviews := sv.GetSubViews()
	if len(subviews) == 0 {
		return
	}
	union := sv.subViewUnionBounds()
	maxScroll := union.X + union.Width - thisBounds.Width
	if maxScroll <= 0 {
		return
	}

	toScroll := amount
	if maxScroll < toScroll {
		toScroll = maxScroll
	}

	for _, subv := range sv.GetSubViews() {
		subv.UpdateBounds(func(b egg.Bounds) egg.Bounds {
			b.X -= toScroll
			return b
		})
	}
	if sv.didScroll != nil {
		sv.didScroll()
	}
}

// ScrollUp ...
func (sv *ScrollView) ScrollUp(amount int) {
	if !sv.scrollVertical {
		return
	}
	subviews := sv.GetSubViews()
	if len(subviews) == 0 {
		return
	}
	union := sv.subViewUnionBounds()
	maxScroll := union.Y * -1
	if maxScroll <= 0 {
		return
	}
	toScroll := amount
	if toScroll > maxScroll {
		toScroll = maxScroll
	}

	for _, subv := range sv.GetSubViews() {
		subv.UpdateBounds(func(b egg.Bounds) egg.Bounds {
			b.Y += toScroll
			return b
		})
	}

	if sv.didScroll != nil {
		sv.didScroll()
	}
}

// ScrollDown ...
func (sv *ScrollView) ScrollDown(amount int) {
	if !sv.scrollVertical {
		return
	}
	thisBounds := sv.GetBounds()
	subviews := sv.GetSubViews()
	if len(subviews) == 0 {
		return
	}
	union := sv.subViewUnionBounds()
	maxScroll := union.Y + union.Height - thisBounds.Height
	if maxScroll <= 0 {
		return
	}

	toScroll := amount
	if toScroll > maxScroll {
		toScroll = maxScroll
	}

	for _, subv := range sv.GetSubViews() {
		subv.UpdateBounds(func(b egg.Bounds) egg.Bounds {
			b.Y -= toScroll
			return b
		})
	}
	if sv.didScroll != nil {
		sv.didScroll()
	}
}

// OnDidScroll - calls this function after scrolling occurred
func (sv *ScrollView) OnDidScroll(f func()) {
	sv.didScroll = f
}

func (sv *ScrollView) subViewUnionBounds() egg.Bounds {
	subviews := sv.GetSubViews()
	minx := math.MaxInt8
	miny := math.MaxInt8

	for _, subview := range subviews {
		b := subview.GetBounds()
		if b.X < minx {
			minx = b.X
		}
		if b.Y < miny {
			miny = b.Y
		}
	}

	maxW := 0
	maxH := 0

	for _, subview := range subviews {
		b := subview.GetBounds()
		w := b.X + b.Width - minx
		h := b.Y + b.Height - miny
		if w > maxW {
			maxW = w
		}
		if h > maxH {
			maxH = h
		}
	}

	return egg.Bounds{
		Origin: egg.Origin{
			X: minx,
			Y: miny,
		},
		Size: egg.Size{
			Width:  maxW,
			Height: maxH,
		},
	}
}
