package egg

// WindowWidth ...
func WindowWidth() int {
	w, _ := WindowSize()
	return w
}

// WindowHeight ...
func WindowHeight() int {
	_, h := WindowSize()
	return h
}

// WindowSize ...
func WindowSize() (width int, height int) {
	return _APP.screen.Size()
}
