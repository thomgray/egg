package egg

import (
	"github.com/gdamore/tcell"
)

type (
	// Modifier - for key
	Modifier uint8
	// Key - key
	Key uint16
)

// Mod constants
const (
	ModAlt = Modifier(tcell.ModAlt)
)

// Key constants, see Event.Key field.
const (
	KeyF1         = Key(tcell.KeyF1)
	KeyF2         = Key(tcell.KeyF2)
	KeyF3         = Key(tcell.KeyF3)
	KeyF4         = Key(tcell.KeyF4)
	KeyF5         = Key(tcell.KeyF5)
	KeyF6         = Key(tcell.KeyF6)
	KeyF7         = Key(tcell.KeyF7)
	KeyF8         = Key(tcell.KeyF8)
	KeyF9         = Key(tcell.KeyF9)
	KeyF10        = Key(tcell.KeyF10)
	KeyF11        = Key(tcell.KeyF11)
	KeyF12        = Key(tcell.KeyF12)
	KeyInsert     = Key(tcell.KeyInsert)
	KeyDelete     = Key(tcell.KeyDelete)
	KeyHome       = Key(tcell.KeyHome)
	KeyEnd        = Key(tcell.KeyEnd)
	KeyPgup       = Key(tcell.KeyPgUp)
	KeyPgdn       = Key(tcell.KeyPgDn)
	KeyArrowUp    = Key(tcell.KeyUp)
	KeyArrowDown  = Key(tcell.KeyDown)
	KeyArrowLeft  = Key(tcell.KeyLeft)
	KeyArrowRight = Key(tcell.KeyRight)
	// MouseLeft      = Key(tcell.Mouse)
	// MouseMiddle    = Key(tcell.MouseMiddle)
	// MouseRight     = Key(tcell.MouseRight)
	// MouseRelease   = Key(tcell.MouseRelease)
	// MouseWheelUp   = Key(tcell.MouseWheelUp)
	// MouseWheelDown = Key(tcell.MouseWheelDown)
)

// Key
const (
	// KeyCtrlTilde      = Key(tcell.KeyCtrlTilde)
	// KeyCtrl2          = Key(tcell.KeyCtrl2)
	KeyCtrlSpace = Key(tcell.KeyCtrlSpace)
	KeyCtrlA     = Key(tcell.KeyCtrlA)
	KeyCtrlB     = Key(tcell.KeyCtrlB)
	KeyCtrlC     = Key(tcell.KeyCtrlC)
	KeyCtrlD     = Key(tcell.KeyCtrlD)
	KeyCtrlE     = Key(tcell.KeyCtrlE)
	KeyCtrlF     = Key(tcell.KeyCtrlF)
	KeyCtrlG     = Key(tcell.KeyCtrlG)
	KeyBackspace = Key(tcell.KeyBackspace)
	KeyCtrlH     = Key(tcell.KeyCtrlH)
	KeyTab       = Key(tcell.KeyTab)
	KeyCtrlI     = Key(tcell.KeyCtrlI)
	KeyCtrlJ     = Key(tcell.KeyCtrlJ)
	KeyCtrlK     = Key(tcell.KeyCtrlK)
	KeyCtrlL     = Key(tcell.KeyCtrlL)
	KeyEnter     = Key(tcell.KeyEnter)
	KeyCtrlM     = Key(tcell.KeyCtrlM)
	KeyCtrlN     = Key(tcell.KeyCtrlN)
	KeyCtrlO     = Key(tcell.KeyCtrlO)
	KeyCtrlP     = Key(tcell.KeyCtrlP)
	KeyCtrlQ     = Key(tcell.KeyCtrlQ)
	KeyCtrlR     = Key(tcell.KeyCtrlR)
	KeyCtrlS     = Key(tcell.KeyCtrlS)
	KeyCtrlT     = Key(tcell.KeyCtrlT)
	KeyCtrlU     = Key(tcell.KeyCtrlU)
	KeyCtrlV     = Key(tcell.KeyCtrlV)
	KeyCtrlW     = Key(tcell.KeyCtrlW)
	KeyCtrlX     = Key(tcell.KeyCtrlX)
	KeyCtrlY     = Key(tcell.KeyCtrlY)
	KeyCtrlZ     = Key(tcell.KeyCtrlZ)
	KeyEsc       = Key(tcell.KeyEsc)
	// KeyCtrlLsqBracket = Key(tcell.KeyCtrlLsqBracket)
	// KeyCtrl3          = Key(tcell.KeyCtrl3)
	// KeyCtrl4          = Key(tcell.KeyCtrl4)
	KeyCtrlBackslash = Key(tcell.KeyCtrlBackslash)
	// KeyCtrl5          = Key(tcell.KeyCtrl5)
	// KeyCtrlRsqBracket = Key(tcell.KeyCtrlRsqBracket)
	// KeyCtrl6          = Key(tcell.KeyCtrl6)
	// KeyCtrl7          = Key(tcell.KeyCtrl7)
	// KeyCtrlSlash      = Key(tcell.KeyCtrlSlash)
	KeyCtrlUnderscore = Key(tcell.KeyCtrlUnderscore)
	// KeySpace          = Key(tcell.KeySpace)
	KeyBackspace2 = Key(tcell.KeyBackspace2)
	// KeyCtrl8          = Key(tcell.KeyCtrl8)
)
