package egg

// SetID - set an id for the view
// The view id is not used for anything internal, nor is there any requirement
// for it to be unique. It is simply there in case you want to use it.
func (v *View) SetID(id string) {
	v.id = id
}

// GetID - get the id for the view
func (v *View) GetID() string {
	return v.id
}

// SetFocusable - set whether this view can be focused
func (v *View) SetFocusable(f bool) {
	v.focusable = f
}

// IsFocusable -
func (v *View) IsFocusable() bool {
	return v.focusable
}

// SetVisible -
func (v *View) SetVisible(visible bool) {
	v.visible = visible
}

// IsVisible -
func (v *View) IsVisible() bool {
	return v.visible
}

// SetTransparent -
func (v *View) SetTransparent(t bool) {
	v.transparent = t
}

// IsTransparent -
func (v *View) IsTransparent() bool {
	return v.transparent
}

// GetBounds - get the current bounds for the view
func (v *View) GetBounds() Bounds {
	return v.bounds
}

// GetSubViews - get all subviews of this view
func (v *View) GetSubViews() []*View {
	return v.subViews
}

// SetViewport - sets the viewport of ths view.
// This would likely depend on the current state of the view, espeically the views current bounds.
// The function takes the views current bounds (relative to the parent).
// The viwport bounds should be returned (also relative to the views parent origin).
// Returning nil indicates that this view does not constrain its child views to a viewport
// (This is the default case)
func (v *View) SetViewport(f func(Bounds) *Bounds) {
	v.viewportAccessor = f
}

// GetViewport - get the viewport for this view
// This defines the area within the view that can be draws by its subviews
// Override this with SetViewport()
func (v *View) GetViewport() *Bounds {
	if v.viewportAccessor != nil {
		return v.viewportAccessor(v.bounds)
	}
	return &v.bounds
}

// SetBounds - set the bounds for the view. Note this will trigger am OnBoundsSet event
func (v *View) SetBounds(bounds Bounds) {
	old := v.bounds
	v.bounds = bounds
	v.ReceiveBoundsUpdate(old, v.bounds)
}

// UpdateBounds - has the effect of SetBounds, only you provide a function to mutate and return the existing bounds.
func (v *View) UpdateBounds(updater func(Bounds) Bounds) {
	old := v.bounds
	new := updater(old)
	v.SetBounds(new)
}

// SetForeground - set the foreground color for the view
func (v *View) SetForeground(fg Color) {
	v.foreground = fg
}

// SetBackground - set the background colour for the view
func (v *View) SetBackground(bg Color) {
	v.background = bg
}

// GetForeground - set the foreground color for the view
func (v *View) GetForeground() Color {
	return v.foreground
}

// GetBackground - set the background colour for the view
func (v *View) GetBackground() Color {
	return v.background
}

// SetAttributes - set attributes for the view
func (v *View) SetAttributes(atts ...Attribute) {
	a := Attribute(0)
	for _, att := range atts {
		a = a | att
	}
	v.attribute = a
}

// OnDraw - specify a draw funtion
func (v *View) OnDraw(f func(c Canvas)) {
	v.drawHandler = f
}

// OnKeyEvent - specify a key event handler function
func (v *View) OnKeyEvent(f func(*KeyEvent)) {
	v.keyEventHandler = f
}

// OnMouseEvent - specify a mouse event handler function
func (v *View) OnMouseEvent(f func(*MouseEvent)) {
	v.mouseEventHandler = f
}

// OnBoundsSet - specify a bound set handler function
func (v *View) OnBoundsSet(f func(old, new Bounds)) {
	v.boundsUpdateHandler = f
}

// SetZIndex - set the view's z-index
func (v *View) SetZIndex(zindex int) {
	v.zindex = zindex
}

// GetZIndex - get the view's z-index
func (v *View) GetZIndex() int {
	return v.zindex
}
