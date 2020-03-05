package egg

import (
	"github.com/nsf/termbox-go"
)

// WindowWidth ...
func WindowWidth() int {
	w, _ := termbox.Size()
	return w
}

// WindowHeight ...
func WindowHeight() int {
	_, h := termbox.Size()
	return h
}

// WindowSize ...
func WindowSize() (width int, height int) {
	return termbox.Size()
}
