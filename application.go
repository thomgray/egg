package egg

import (
	"sync"

	"github.com/nsf/termbox-go"
)

var _APP *Application

// Application ...
type Application struct {
	view               *applicationView
	exitOnSigInt       bool
	eventDelegate      func(*Event)
	keyEventHandler    func(*KeyEvent)
	mouseEventHandler  func(*MouseEvent)
	resizeEventHandler func(*ResizeEvent)
	running            bool
	focusedView        *View
	redrawDebouncer    *Debouncer
	mux                sync.Mutex
}

// Stop ...
func (app *Application) Stop() {
	app.running = false
}

// Start ...
func (app *Application) Start() {
	app.ReDraw()
	defer termbox.Close()
mainloop:
	for {
		if !app.running {
			break mainloop
		} else {
			e := pollEvent()
			app.handleEvent(e)
		}
	}
}

func (app *Application) handleEvent(e *Event) {
	if app.eventDelegate != nil {
		app.eventDelegate(e)
	}
	// even if propagation stopped, always resize the main app view
	if e.Resize != nil {
		app.view.SetBounds(MakeBounds(0, 0, e.Resize.Width, e.Resize.Height))
	}
	if !e.StopPropagation {
		if e.Mouse != nil {
			app.handleMouseEvent(e.Mouse)
		} else if e.Key != nil {
			app.handleKeyEvent(e.Key)
		} else if e.Resize != nil {
			app.handleResizeEvent(e.Resize)
		} else if e.Error != nil {
			app.running = false
		}
	}
}

func (app *Application) handleMouseEvent(me *MouseEvent) {
	app.mouseEventHandler(me)
}

func (app *Application) handleKeyEvent(ke *KeyEvent) {
	if app.exitOnSigInt && (ke.Key == KeyCtrlC) {
		app.running = false
	} else {
		app.keyEventHandler(ke)
		if app.focusedView != nil && !ke.StopPropagation {
			app.focusedView.ReceiveKeyEvent(ke)
		}
	}
}

func (app *Application) handleResizeEvent(re *ResizeEvent) {
	app.resizeEventHandler(re)
}

// AddView ...
func (app *Application) AddView(view *View) {
	app.view.AddSubView(view)
}

// AddViewController ...
func (app *Application) AddViewController(vc ViewController) {
	app.AddView(vc.GetView())
}

// ReDraw ...
func (app *Application) ReDraw() {
	app.redrawDebouncer.Send(true)
	// termbox.Clear(termbox.Attribute(app.view.GetForeground()), termbox.Attribute(app.view.GetBackground()))
	// termbox.HideCursor()
	// app.view.redraw()
	// termbox.Flush()
}

func (app *Application) redrawBebounced(b []interface{}) {
	app.mux.Lock()
	termbox.Clear(termbox.Attribute(app.view.GetForeground()), termbox.Attribute(app.view.GetBackground()))
	termbox.HideCursor()
	app.view.redraw()
	termbox.Flush()
	app.mux.Unlock()
}
