package egg

import (
	"github.com/gdamore/tcell"
)

// Init ...
func Init() (*Application, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, e := tcell.NewScreen()
	if e != nil {
		return nil, e
	}
	if e = screen.Init(); e != nil {
		return nil, e
	}

	if _APP != nil {
		panic("Application already initialized")
	}

	_APP = &Application{
		screen:       screen,
		exitOnSigInt: true,
		running:      true,
		focusedView:  nil,
	}
	_APP.keyEventHandler = func(*KeyEvent) {}
	_APP.mouseEventHandler = func(*MouseEvent) {}
	_APP.resizeEventHandler = func(*ResizeEvent) { _APP.ReDraw() }
	_APP.eventDelegate = func(Event) {}
	_APP.redrawDebouncer = MakeDebouncer()
	_APP.redrawDebouncer.Receive(_APP.redrawBebounced)

	baseView := &applicationView{
		*MakeView(),
	}

	baseView.SetFocusable(true)
	baseView.SetViewport(func(Bounds) *Bounds {
		return nil
	})
	w, h := _APP.WindowSize()
	baseView.SetBounds(MakeBounds(0, 0, w, h))
	baseView.attribute = 0
	baseView.foreground = ColorDefault
	baseView.background = ColorDefault

	_APP.view = baseView

	return _APP, nil
}

// InitOrPanic - Initialise an application. Panic if an error occurs
func InitOrPanic() *Application {
	var app, err = Init()
	if err != nil {
		panic(err)
	} else {
		return app
	}
}

// GetApplication - returns the current application if initialised
func GetApplication() *Application {
	return _APP
}
