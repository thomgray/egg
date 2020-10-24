package egg

import (
	"github.com/gdamore/tcell"
)

type (
	// Modifier - for key
	Modifier uint8
	// Key - key
	Key uint16
	// MouseButton button mask - mouse key
	MouseButton uint16
)

// Mod constants
const (
	ModAlt = Modifier(tcell.ModAlt)
)

// This is the list of named keys.  KeyRune is special however, in that it is
// a place holder key indicating that a printable character was sent.  The
// actual value of the rune will be transported in the Rune of the associated
// EventKey.
const (
	KeyRune      = Key(tcell.KeyRune)
	KeyUp        = Key(tcell.KeyUp)
	KeyDown      = Key(tcell.KeyDown)
	KeyRight     = Key(tcell.KeyRight)
	KeyLeft      = Key(tcell.KeyLeft)
	KeyUpLeft    = Key(tcell.KeyUpLeft)
	KeyUpRight   = Key(tcell.KeyUpRight)
	KeyDownLeft  = Key(tcell.KeyDownLeft)
	KeyDownRight = Key(tcell.KeyDownRight)
	KeyCenter    = Key(tcell.KeyCenter)
	KeyPgUp      = Key(tcell.KeyPgUp)
	KeyPgDn      = Key(tcell.KeyPgDn)
	KeyHome      = Key(tcell.KeyHome)
	KeyEnd       = Key(tcell.KeyEnd)
	KeyInsert    = Key(tcell.KeyInsert)
	KeyDelete    = Key(tcell.KeyDelete)
	KeyHelp      = Key(tcell.KeyHelp)
	KeyExit      = Key(tcell.KeyExit)
	KeyClear     = Key(tcell.KeyClear)
	KeyCancel    = Key(tcell.KeyCancel)
	KeyPrint     = Key(tcell.KeyPrint)
	KeyPause     = Key(tcell.KeyPause)
	KeyBacktab   = Key(tcell.KeyBacktab)
	KeyF1        = Key(tcell.KeyF1)
	KeyF2        = Key(tcell.KeyF2)
	KeyF3        = Key(tcell.KeyF3)
	KeyF4        = Key(tcell.KeyF4)
	KeyF5        = Key(tcell.KeyF5)
	KeyF6        = Key(tcell.KeyF6)
	KeyF7        = Key(tcell.KeyF7)
	KeyF8        = Key(tcell.KeyF8)
	KeyF9        = Key(tcell.KeyF9)
	KeyF10       = Key(tcell.KeyF10)
	KeyF11       = Key(tcell.KeyF11)
	KeyF12       = Key(tcell.KeyF12)
	KeyF13       = Key(tcell.KeyF13)
	KeyF14       = Key(tcell.KeyF14)
	KeyF15       = Key(tcell.KeyF15)
	KeyF16       = Key(tcell.KeyF16)
	KeyF17       = Key(tcell.KeyF17)
	KeyF18       = Key(tcell.KeyF18)
	KeyF19       = Key(tcell.KeyF19)
	KeyF20       = Key(tcell.KeyF20)
	KeyF21       = Key(tcell.KeyF21)
	KeyF22       = Key(tcell.KeyF22)
	KeyF23       = Key(tcell.KeyF23)
	KeyF24       = Key(tcell.KeyF24)
	KeyF25       = Key(tcell.KeyF25)
	KeyF26       = Key(tcell.KeyF26)
	KeyF27       = Key(tcell.KeyF27)
	KeyF28       = Key(tcell.KeyF28)
	KeyF29       = Key(tcell.KeyF29)
	KeyF30       = Key(tcell.KeyF30)
	KeyF31       = Key(tcell.KeyF31)
	KeyF32       = Key(tcell.KeyF32)
	KeyF33       = Key(tcell.KeyF33)
	KeyF34       = Key(tcell.KeyF34)
	KeyF35       = Key(tcell.KeyF35)
	KeyF36       = Key(tcell.KeyF36)
	KeyF37       = Key(tcell.KeyF37)
	KeyF38       = Key(tcell.KeyF38)
	KeyF39       = Key(tcell.KeyF39)
	KeyF40       = Key(tcell.KeyF40)
	KeyF41       = Key(tcell.KeyF41)
	KeyF42       = Key(tcell.KeyF42)
	KeyF43       = Key(tcell.KeyF43)
	KeyF44       = Key(tcell.KeyF44)
	KeyF45       = Key(tcell.KeyF45)
	KeyF46       = Key(tcell.KeyF46)
	KeyF47       = Key(tcell.KeyF47)
	KeyF48       = Key(tcell.KeyF48)
	KeyF49       = Key(tcell.KeyF49)
	KeyF50       = Key(tcell.KeyF50)
	KeyF51       = Key(tcell.KeyF51)
	KeyF52       = Key(tcell.KeyF52)
	KeyF53       = Key(tcell.KeyF53)
	KeyF54       = Key(tcell.KeyF54)
	KeyF55       = Key(tcell.KeyF55)
	KeyF56       = Key(tcell.KeyF56)
	KeyF57       = Key(tcell.KeyF57)
	KeyF58       = Key(tcell.KeyF58)
	KeyF59       = Key(tcell.KeyF59)
	KeyF60       = Key(tcell.KeyF60)
	KeyF61       = Key(tcell.KeyF61)
	KeyF62       = Key(tcell.KeyF62)
	KeyF63       = Key(tcell.KeyF63)
	KeyF64       = Key(tcell.KeyF64)
)

// These are the control keys.  Note that they overlap with other keys,
// perhaps.  For example, KeyCtrlH is the same as KeyBackspace.
const (
	KeyCtrlSpace      = Key(tcell.KeyCtrlSpace)
	KeyCtrlA          = Key(tcell.KeyCtrlA)
	KeyCtrlB          = Key(tcell.KeyCtrlB)
	KeyCtrlC          = Key(tcell.KeyCtrlC)
	KeyCtrlD          = Key(tcell.KeyCtrlD)
	KeyCtrlE          = Key(tcell.KeyCtrlE)
	KeyCtrlF          = Key(tcell.KeyCtrlF)
	KeyCtrlG          = Key(tcell.KeyCtrlG)
	KeyCtrlH          = Key(tcell.KeyCtrlH)
	KeyCtrlI          = Key(tcell.KeyCtrlI)
	KeyCtrlJ          = Key(tcell.KeyCtrlJ)
	KeyCtrlK          = Key(tcell.KeyCtrlK)
	KeyCtrlL          = Key(tcell.KeyCtrlL)
	KeyCtrlM          = Key(tcell.KeyCtrlM)
	KeyCtrlN          = Key(tcell.KeyCtrlN)
	KeyCtrlO          = Key(tcell.KeyCtrlO)
	KeyCtrlP          = Key(tcell.KeyCtrlP)
	KeyCtrlQ          = Key(tcell.KeyCtrlQ)
	KeyCtrlR          = Key(tcell.KeyCtrlR)
	KeyCtrlS          = Key(tcell.KeyCtrlS)
	KeyCtrlT          = Key(tcell.KeyCtrlT)
	KeyCtrlU          = Key(tcell.KeyCtrlU)
	KeyCtrlV          = Key(tcell.KeyCtrlV)
	KeyCtrlW          = Key(tcell.KeyCtrlW)
	KeyCtrlX          = Key(tcell.KeyCtrlX)
	KeyCtrlY          = Key(tcell.KeyCtrlY)
	KeyCtrlZ          = Key(tcell.KeyCtrlZ)
	KeyCtrlLeftSq     = Key(tcell.KeyCtrlLeftSq)
	KeyCtrlBackslash  = Key(tcell.KeyCtrlBackslash)
	KeyCtrlRightSq    = Key(tcell.KeyCtrlRightSq)
	KeyCtrlCarat      = Key(tcell.KeyCtrlCarat)
	KeyCtrlUnderscore = Key(tcell.KeyCtrlUnderscore)
)

// These are the defined ASCII values for key codes.  They generally match
// with KeyCtrl values.
const (
	KeyNUL = Key(tcell.KeyNUL)
	KeySOH = Key(tcell.KeySOH)
	KeySTX = Key(tcell.KeySTX)
	KeyETX = Key(tcell.KeyETX)
	KeyEOT = Key(tcell.KeyEOT)
	KeyENQ = Key(tcell.KeyENQ)
	KeyACK = Key(tcell.KeyACK)
	KeyBEL = Key(tcell.KeyBEL)
	KeyBS  = Key(tcell.KeyBS)
	KeyTAB = Key(tcell.KeyTAB)
	KeyLF  = Key(tcell.KeyLF)
	KeyVT  = Key(tcell.KeyVT)
	KeyFF  = Key(tcell.KeyFF)
	KeyCR  = Key(tcell.KeyCR)
	KeySO  = Key(tcell.KeySO)
	KeySI  = Key(tcell.KeySI)
	KeyDLE = Key(tcell.KeyDLE)
	KeyDC1 = Key(tcell.KeyDC1)
	KeyDC2 = Key(tcell.KeyDC2)
	KeyDC3 = Key(tcell.KeyDC3)
	KeyDC4 = Key(tcell.KeyDC4)
	KeyNAK = Key(tcell.KeyNAK)
	KeySYN = Key(tcell.KeySYN)
	KeyETB = Key(tcell.KeyETB)
	KeyCAN = Key(tcell.KeyCAN)
	KeyEM  = Key(tcell.KeyEM)
	KeySUB = Key(tcell.KeySUB)
	KeyESC = Key(tcell.KeyESC)
	KeyFS  = Key(tcell.KeyFS)
	KeyGS  = Key(tcell.KeyGS)
	KeyRS  = Key(tcell.KeyRS)
	KeyUS  = Key(tcell.KeyUS)
	KeyDEL = Key(tcell.KeyDEL)
)

// These keys are aliases for other names.
const (
	KeyBackspace  = Key(tcell.KeyBackspace)
	KeyTab        = Key(tcell.KeyTab)
	KeyEsc        = Key(tcell.KeyEsc)
	KeyEscape     = Key(tcell.KeyEscape)
	KeyEnter      = Key(tcell.KeyEnter)
	KeyBackspace2 = Key(tcell.KeyBackspace2)
)

// Mouse buttons
const (
	MouseButton1    = MouseButton(tcell.Button1) // Usually left mouse button.
	MouseButton2    = MouseButton(tcell.Button2) // Usually the middle mouse button.
	MouseButton3    = MouseButton(tcell.Button3) // Usually the right mouse button.
	MouseButton4    = MouseButton(tcell.Button4) // Often a side button (thumb/next).
	MouseButton5    = MouseButton(tcell.Button5) // Often a side button (thumb/prev).
	MouseButton6    = MouseButton(tcell.Button6)
	MouseButton7    = MouseButton(tcell.Button7)
	MouseButton8    = MouseButton(tcell.Button8)
	MouseWheelUp    = MouseButton(tcell.WheelUp)    // Wheel motion up/away from user.
	MouseWheelDown  = MouseButton(tcell.WheelDown)  // Wheel motion down/towards user.
	MouseWheelLeft  = MouseButton(tcell.WheelLeft)  // Wheel motion to left.
	MouseWheelRight = MouseButton(tcell.WheelRight) // Wheel motion to right.
	MouseButtonNone = MouseButton(tcell.ButtonNone) // No button or wheel events.
)
