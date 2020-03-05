package egg

import (
	"github.com/nsf/termbox-go"
)

// View - a component that represents a section of the application's viewport
type View struct {
	id                  string
	bounds              Bounds
	foreground          Color
	background          Color
	attribute           Attribute
	focusable           bool
	transparent         bool
	visible             bool
	subViews            []*View
	superView           *View
	mouseEventHandler   func(*MouseEvent)
	keyEventHandler     func(*KeyEvent)
	boundsUpdateHandler func(Bounds, Bounds)
	drawHandler         func(Canvas)
	viewportAccessor    func(Bounds) *Bounds
}

// MakeView - make a new view.
func MakeView() *View {
	view := View{}
	view.focusable = true
	view.visible = true
	view.transparent = false
	view.subViews = make([]*View, 0)
	view.foreground = ColorDefault
	view.background = ColorDefault
	return &view
}

// ReceiveKeyEvent - bare implementation
func (v *View) ReceiveKeyEvent(ke *KeyEvent) {
	if v.keyEventHandler != nil {
		v.keyEventHandler(ke)
	}
}

// ReceiveMouseEvent - bare implementation
func (v *View) ReceiveMouseEvent(me *MouseEvent) {
	if v.mouseEventHandler != nil {
		v.mouseEventHandler(me)
	}
}

// ReceiveBoundsUpdate - bare implementation
func (v *View) ReceiveBoundsUpdate(old, new Bounds) {
	if v.boundsUpdateHandler != nil {
		v.boundsUpdateHandler(old, new)
	}
}

// Unmount - remove this view from it's super view
func (v *View) Unmount() {
	if v.superView != nil {
		superViewSubViews := v.superView.getSubviews()
		for i, subv := range superViewSubViews {
			if subv == v {
				newSubviews := append(superViewSubViews[:i], superViewSubViews[i+1:]...)
				v.superView.setSubviews(newSubviews)
				break
			}
		}
		v.superView = nil
	}
}

// AddSubView - add subview to this view
func (v *View) AddSubView(sv *View) {
	if sv == nil {
		return
	}
	for _, vw := range v.subViews {
		if vw == sv {
			return
		}
	}
	v.subViews = append(v.subViews, sv)
	sv.setSuperview(v)
}

// GainFocus - set this view to gain focus
func (v *View) GainFocus() {
	_APP.focusedView = v
}

// IsFocused - is this view focused?
func (v *View) IsFocused() bool {
	return _APP.focusedView == v
}

// ReDraw - re draw this view
func (v *View) ReDraw() {
	v.redraw()
	termbox.Flush()
}

func (v *View) getSubviews() []*View {
	return v.subViews
}

func (v *View) setSubviews(svs []*View) {
	v.subViews = svs
}

func (v *View) setSuperview(sv *View) {
	v.superView = sv
}

func (v *View) getSuperview() *View {
	return v.superView
}

func (v *View) redraw() {
	if !v.visible {
		return
	}
	bounds := absoluteBounds(v)
	var viewPort *Bounds = nil
	if v.superView != nil {
		sv := v.superView
		svport := sv.GetViewport()
		if svport != nil {
			svb := absoluteBounds2(*svport, sv)
			viewPort = &svb
		}
	}
	if v.IsTransparent() {
		// don't clear background
	} else if viewPort != nil {
		intersect := intersectionBounds(bounds, *viewPort)
		clearBounds(intersect, v.foreground, v.background)
	} else {
		clearBounds(bounds, v.foreground, v.background)
	}
	c := makeCanvasWithViewPort(bounds, viewPort, v.foreground, v.background, v.attribute)
	if v.drawHandler != nil {
		v.drawHandler(c)
	}
	for _, subv := range v.subViews {
		subv.redraw()
	}
}
