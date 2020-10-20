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

	ColorBlack   = Color(0)
	ColorRed     = Color(1)
	ColorGreen   = Color(2)
	ColorYellow  = Color(3)
	ColorBlue    = Color(4)
	ColorMagenta = Color(5)
	ColorCyan    = Color(6)
	ColorWhite   = Color(7)

	ColorBrightBlack   = Color(8)
	ColorBrightRed     = Color(9)
	ColorBrightGreen   = Color(10)
	ColorBrightYellow  = Color(11)
	ColorBrightBlue    = Color(12)
	ColorBrightMagenta = Color(13)
	ColorBrightCyan    = Color(14)
	ColorBrightWhite   = Color(15)
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
