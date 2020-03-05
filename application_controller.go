package egg

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Init ...
func Init() (*Application, error) {
	if _APP != nil {
		panic("Application already initialized")
	}

	err := termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		return nil, err
	}
	baseView := &applicationView{
		*MakeView(),
	}

	baseView.SetFocusable(true)
	baseView.SetViewport(func(Bounds) *Bounds {
		return nil
	})
	w, h := WindowSize()
	baseView.SetBounds(MakeBounds(0, 0, w, h))
	baseView.attribute = 0
	baseView.foreground = ColorDefault
	baseView.background = ColorDefault

	_APP = &Application{
		view:         baseView,
		exitOnSigInt: true,
		running:      true,
		focusedView:  nil,
	}
	_APP.keyEventHandler = func(*KeyEvent) {}
	_APP.mouseEventHandler = func(*MouseEvent) {}
	_APP.resizeEventHandler = func(*ResizeEvent) { _APP.ReDraw() }
	_APP.eventDelegate = func(*Event) {}
	_APP.redrawDebouncer = MakeDebouncer()
	_APP.redrawDebouncer.Receive(_APP.redrawBebounced)

	return _APP, nil
}

// InitOrPanic - Initialise an application. Panic if an error occurs
func InitOrPanic() *Application {
	var app, err = Init()
	if err != nil {
		panic(fmt.Sprintf("There was an error: %s", err))
	} else {
		return app
	}
}

// GetApplication - returns the current application if initialised
func GetApplication() *Application {
	return _APP
}
