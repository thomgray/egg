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

// SetSize - set the size for the view with the existing origin. Note this will trigger am OnBoundsSet event
func (v *View) SetSize(w, h int) {
	bounds := v.bounds
	bounds.Width = w
	bounds.Height = h
	v.SetBounds(bounds)
}

// SetOrigin - set the origin for the view with the existing size. Note this will trigger am OnBoundsSet event
func (v *View) SetOrigin(x, y int) {
	bounds := v.bounds
	bounds.Y = y
	bounds.X = x
	v.SetBounds(bounds)
}

// SetWidth - set the width of the view with the existing origin and height. Note this will trigger am OnBoundsSet event
func (v *View) SetWidth(w int) {
	bounds := v.bounds
	bounds.Width = w
	v.SetBounds(bounds)
}

// SetHeight - set the height of the view with the existing origin and width. Note this will trigger am OnBoundsSet event
func (v *View) SetHeight(h int) {
	bounds := v.bounds
	bounds.Height = h
	v.SetBounds(bounds)
}

// SetX - set the x origin of the view with the existing size and y origin. Note this will trigger am OnBoundsSet event
func (v *View) SetX(x int) {
	bounds := v.bounds
	bounds.X = x
	v.SetBounds(bounds)
}

// SetY - set the y origin of the view with the existing size and x origin. Note this will trigger am OnBoundsSet event
func (v *View) SetY(y int) {
	bounds := v.bounds
	bounds.Y = y
	v.SetBounds(bounds)
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
