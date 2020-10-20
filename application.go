package egg

import (
	"sync"

	"github.com/gdamore/tcell"
)

var _APP *Application

// Application ...
type Application struct {
	screen             tcell.Screen
	view               *applicationView
	exitOnSigInt       bool
	eventDelegate      func(Event)
	keyEventHandler    func(*KeyEvent)
	mouseEventHandler  func(*MouseEvent)
	resizeEventHandler func(*ResizeEvent)
	running            bool
	focusedView        *View
	redrawDebouncer    *Debouncer
	mux                sync.Mutex
	state              interface{}
}

// Stop ...
func (app *Application) Stop() {
	app.running = false
}

func (app *Application) WindowSize() (w, h int) {
	return app.screen.Size()
}

// Start ...
func (app *Application) Start() {
	app.ReDraw()
	defer app.screen.Fini()
mainloop:
	for {
		if !app.running {
			break mainloop
		} else {
			e := pollEvent(app.screen)
			app.handleEvent(e)
		}
	}
}

func (app *Application) handleEvent(e Event) {
	if e == nil {
		return
	}
	if app.eventDelegate != nil {
		app.eventDelegate(e)
	}

	// even if propagation stopped, always resize the main app view
	// if e.Resize != nil {
	// 	app.view.SetBounds(MakeBounds(0, 0, e.Resize.Width, e.Resize.Height))
	// }
	if !e.ShouldPropagate() {
		return
	}

	switch e := e.(type) {
	case *KeyEvent:
		app.handleKeyEvent(e)
	case *MouseEvent:
		app.handleMouseEvent(e)
	case *ResizeEvent:
		app.handleResizeEvent(e)
		// error?
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
		if app.focusedView != nil && ke.ShouldPropagate() {
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
}

func (app *Application) redrawBebounced(b []interface{}) {
	app.mux.Lock()
	app.screen.Clear()
	app.view.redraw()
	app.screen.Show()
	app.mux.Unlock()
}
