package egg

import "github.com/nsf/termbox-go"

type (
	// Attribute - a format attribute
	Attribute uint16
	// Color - color
	Color uint16
)

// Colour
const (
	ColorDefault = Color(termbox.ColorDefault)
	ColorBlack   = Color(termbox.ColorBlack)
	ColorRed     = Color(termbox.ColorRed)
	ColorGreen   = Color(termbox.ColorGreen)
	ColorYellow  = Color(termbox.ColorYellow)
	ColorBlue    = Color(termbox.ColorBlue)
	ColorMagenta = Color(termbox.ColorMagenta)
	ColorCyan    = Color(termbox.ColorCyan)
	ColorWhite   = Color(termbox.ColorWhite)
)

// Attribute
const (
	AttrNormal    = Attribute(0)
	AttrBold      = Attribute(termbox.AttrBold)
	AttrUnderline = Attribute(termbox.AttrUnderline)
	AttrReverse   = Attribute(termbox.AttrReverse)
)

// ColorAnsi - convert an ansi colour value 0-255 to a Color
func ColorAnsi(i int) Color {
	return Color(i)
}
