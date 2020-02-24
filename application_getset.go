package egg

// ExitOnSigInt ...
func (app *Application) ExitOnSigInt(b bool) {
	app.exitOnSigInt = b
}

// OnKeyEvent - set the key event handler for the application
func (app *Application) OnKeyEvent(handler func(*KeyEvent)) {
	if handler != nil {
		app.keyEventHandler = handler
	}
}

// OnMouseEvent - set the mouse event handler for the application
func (app *Application) OnMouseEvent(handler func(*MouseEvent)) {
	if handler != nil {
		app.mouseEventHandler = handler
	}
}

// OnResizeEvent - set the resize event handler for the application
func (app *Application) OnResizeEvent(handler func(*ResizeEvent)) {
	if handler != nil {
		app.resizeEventHandler = handler
	}
}

// SetForeground ...
func (app *Application) SetForeground(fg Color) {
	app.view.background = fg
}

// SetBackground ...
func (app *Application) SetBackground(bg Color) {
	app.view.background = bg
}

// SetAttribute ...
func (app *Application) SetAttribute(atts ...Attribute) {
	a := Attribute(0)
	for _, att := range atts {
		a |= att
	}
	app.view.attribute = a
}
