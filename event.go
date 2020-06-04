package egg

import (
	"github.com/gdamore/tcell"
)

// func pollEvent() *Event {
// 	ev := termbox.PollEvent()

// 	e := Event{}
// 	e.Error = ev.Err

// 	switch ev.Type {
// 	case termbox.EventMouse:
// 		e.Mouse = &MouseEvent{
// 			MouseX: ev.MouseX,
// 			MouseY: ev.MouseY,
// 		}
// 	case termbox.EventKey:
// 		e.Key = &KeyEvent{
// 			Char: ev.Ch,
// 			Mod:  Modifier(ev.Mod),
// 			Key:  Key(ev.Key),
// 		}
// 	case termbox.EventResize:
// 		e.Resize = &ResizeEvent{
// 			Width:  ev.Width,
// 			Height: ev.Height,
// 		}
// 	case termbox.EventInterrupt | termbox.EventError:
// 		e.Error = fmt.Errorf("Event poll failed to to %d", ev.Type)
// 	}

// 	return &e
// }

func pollEvent(s tcell.Screen) Event {
	ev := s.PollEvent()
	var e Event
	// e := Event{}

	switch ev := ev.(type) {
	case *tcell.EventKey:
		e = &KeyEvent{
			propagate: true,
			Char:      ev.Rune(),
			Mod:       Modifier(ev.Modifiers()),
			Key:       Key(uint16(ev.Key())),
		}
	case *tcell.EventMouse:
		x, y := ev.Position()
		e = &MouseEvent{
			propagate: true,
			MouseX:    x,
			MouseY:    y,
		}
	case *tcell.EventResize:
		w, h := ev.Size()
		e = &ResizeEvent{
			propagate: true,
			Width:     w,
			Height:    h,
		}
	}

	return e
}

// Event ...
type Event interface {
	SetPropagate(bool)
	ShouldPropagate() bool
}

// BasicEvent - fields common to all event types
type BasicEvent struct {
	stopPropagation bool
}

// Event - contains a specific event or error
// type Event struct {
// 	BasicEvent
// 	Mouse  *MouseEvent
// 	Key    *KeyEvent
// 	Resize *ResizeEvent
// 	Error  error
// }

// MouseEvent - a mouse event
type MouseEvent struct {
	propagate bool
	MouseX    int
	MouseY    int
}

// KeyEvent - a key event
type KeyEvent struct {
	propagate bool
	BasicEvent
	Char rune
	Mod  Modifier
	Key  Key
}

// ResizeEvent - a resize event
type ResizeEvent struct {
	propagate bool
	BasicEvent
	Width  int
	Height int
}

func (ke *KeyEvent) SetPropagate(t bool) {
	ke.propagate = t
}

func (ke *KeyEvent) ShouldPropagate() bool {
	return ke.propagate
}

func (me *MouseEvent) SetPropagate(t bool) {
	me.propagate = t
}

func (me *MouseEvent) ShouldPropagate() bool {
	return me.propagate
}

func (re *ResizeEvent) SetPropagate(t bool) {
	re.propagate = t
}

func (re *ResizeEvent) ShouldPropagate() bool {
	return re.propagate
}
