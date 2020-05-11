package egg

import (
	"github.com/gdamore/tcell"
)

type (
	// Attribute - a format attribute
	Attribute int
	// Color - color
	Color int16
)

// Colour
const (
	ColorDefault = Color(tcell.ColorDefault)
	ColorBlack   = Color(tcell.ColorBlack)
	ColorRed     = Color(tcell.ColorRed)
	ColorGreen   = Color(tcell.ColorGreen)
	ColorYellow  = Color(tcell.ColorYellow)
	ColorBlue    = Color(tcell.ColorBlue)
	ColorMagenta = Color(tcell.ColorFuchsia)
	ColorCyan    = Color(tcell.ColorAqua)
	ColorWhite   = Color(tcell.ColorWhite)
)

// Attribute
const (
	AttrNormal    = Attribute(0)
	AttrBold      = Attribute(tcell.AttrBold)
	AttrUnderline = Attribute(tcell.AttrUnderline)
	AttrReverse   = Attribute(tcell.AttrReverse)
)

// ColorAnsi - convert an ansi colour value 0-255 to a Color
func ColorAnsi(i int) Color {
	return Color(i)
}
