package egg

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func pollEvent() *Event {
	ev := termbox.PollEvent()

	e := Event{}
	e.Error = ev.Err

	switch ev.Type {
	case termbox.EventMouse:
		e.Mouse = &MouseEvent{
			MouseX: ev.MouseX,
			MouseY: ev.MouseY,
		}
	case termbox.EventKey:
		e.Key = &KeyEvent{
			Char: ev.Ch,
			Mod:  Modifier(ev.Mod),
			Key:  Key(ev.Key),
		}
	case termbox.EventResize:
		e.Resize = &ResizeEvent{
			Width:  ev.Width,
			Height: ev.Height,
		}
	case termbox.EventInterrupt | termbox.EventError:
		e.Error = fmt.Errorf("Event poll failed to to %d", ev.Type)
	}

	return &e
}

// BasicEvent - fields common to all event types
type BasicEvent struct {
	StopPropagation bool
}

// Event - contains a specific event or error
type Event struct {
	BasicEvent
	Mouse  *MouseEvent
	Key    *KeyEvent
	Resize *ResizeEvent
	Error  error
}

// MouseEvent - a mouse event
type MouseEvent struct {
	BasicEvent
	MouseX int
	MouseY int
}

// KeyEvent - a key event
type KeyEvent struct {
	BasicEvent
	Char rune
	Mod  Modifier
	Key  Key
}

// ResizeEvent - a resize event
type ResizeEvent struct {
	BasicEvent
	Width  int
	Height int
}
