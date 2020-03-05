package eggc

import (
	"github.com/thomgray/egg"
	"github.com/thomgray/egg/model"
)

// TextView a view that contains editable text
type TextView struct {
	*egg.View
	document                *model.Document
	alignmentH              AlignmentHorizontal
	alignmentV              AlignmentVertical
	cursorVisible           bool
	showCursorWhenUnfocused bool
	label                   *LabelView
}

// MakeTextView - make a text view
func MakeTextView() *TextView {
	tv := TextView{
		View:                    egg.MakeView(),
		document:                model.MakeDocument(),
		cursorVisible:           true,
		showCursorWhenUnfocused: false,
	}
	label := MakeLabelView()
	label.SetLabel("")
	label.SetBounds(egg.MakeBounds(0, 10, 20, 1))
	tv.label = label
	tv.OnDraw(tv.draw)
	tv.OnKeyEvent(tv.keyEventHandler)
	tv.View.AddSubView(label.View)
	return &tv
}

// GetView - to comply with ViewController interface: returns the view
func (tv *TextView) GetView() *egg.View {
	return tv.View
}

// GetTextContent - get the text content of this view as a byte slice
func (tv *TextView) GetTextContent() []rune {
	return tv.document.GetTextContent()
}

// GetTextContentString - get the text content of this view as a string
func (tv *TextView) GetTextContentString() string {
	return string(tv.document.GetTextContent())
}

// SetTextContentString - set the text content of this view
func (tv *TextView) SetTextContentString(textContent string) {
	tv.document.SetTextContentString(textContent)
}

// SetCursorVisible ...
func (tv *TextView) SetCursorVisible(visible bool) {
	tv.cursorVisible = visible
}

// SetCursorVisibleWhenUnfocused - show the cursor even when the component isn't focused?
// False by default
func (tv *TextView) SetCursorVisibleWhenUnfocused(v bool) {
	tv.cursorVisible = v
}

func (tv *TextView) draw(c egg.Canvas) {
	// fmt.Printf("Drawing txt view")
	curX := 0
	curY := 0
	cursorDocX := tv.document.GetCursorX()
	if !tv.cursorVisible {
		c.DrawCursor(-1, -1)
	}
	for docx, r := range tv.document.GetTextContentNullTerminating() {
		if docx == cursorDocX && tv.cursorVisible && tv.IsFocused() {
			c.DrawCursor(curX, curY)
		}
		if r == '\n' {
			curX = 0
			curY++
		} else {
			c.DrawRune2(r, curX, curY)
			curX++
		}
	}
}

func (tv *TextView) keyEventHandler(e *egg.KeyEvent) {
	switch e.Key {
	case egg.KeyBackspace | egg.KeyBackspace2:
		if tv.document.GetCursorX() > 0 {
			tv.moveCursorBackwards()
			tv.deleteUnderCursor()
		}
	case egg.KeyDelete:
		tv.deleteUnderCursor()
	case egg.KeyArrowDown:
		tv.moveCursorDown()
	case egg.KeyArrowUp:
		tv.moveCursorUp()
	case egg.KeyArrowLeft:
		tv.moveCursorBackwards()
	case egg.KeyArrowRight:
		tv.moveCursorForward()
	case egg.KeyHome:
		tv.moveCursorToBeginningOfLine()
	case egg.KeyEnd:
		tv.moveCursorToEndOfLine()
	case egg.KeySpace:
		tv.insertRune(' ')
	case egg.KeyEnter:
		tv.insertRune('\n')
	default:
		if e.Char != 0 {
			tv.insertRune(e.Char)
		}
	}
	tv.ReDraw()
}

func (tv *TextView) insertRune(r rune) {
	tv.document.InsertRuneAtCursorIncrementing(r)
}

func (tv *TextView) moveCursorForward() {
	tv.document.SetCursorX(tv.document.GetCursorX() + 1)
}

func (tv *TextView) moveCursorBackwards() {
	tv.document.SetCursorX(tv.document.GetCursorX() - 1)
}

func (tv *TextView) moveCursorToBeginningOfLine() {
	text := tv.document.GetTextContentNullTerminating()
	cursorx := tv.document.GetCursorX()

	newLineIndexes := runeSliceIndexes(text, '\n')
	nlBeforeCursor, _ := splitIntSliceBetween(newLineIndexes, cursorx)

	if len(nlBeforeCursor) > 0 {
		tv.document.SetCursorX(nlBeforeCursor[len(nlBeforeCursor)-1] + 1)
	} else {
		tv.document.SetCursorX(0)
	}
}

func (tv *TextView) moveCursorToEndOfLine() {
	text := tv.document.GetTextContent()
	cursorx := tv.document.GetCursorX()

	newLineIndexes := runeSliceIndexes(text, '\n')
	_, nlAfterCursor := splitIntSliceBetween(newLineIndexes, cursorx)

	if len(nlAfterCursor) > 0 {
		tv.document.SetCursorX(nlAfterCursor[0])
	} else {
		tv.document.SetCursorX(len(text))
	}
}

func (tv *TextView) moveCursorDown() {
	text := tv.document.GetTextContentNullTerminating()
	cursorx := tv.document.GetCursorX()
	newLineIndexes := runeSliceIndexes(text, '\n')
	if len(newLineIndexes) < 1 {
		// can't move across lines where there is only one!
		return
	}
	nlBeforeCursor, nlAfterCursor := splitIntSliceBetween(newLineIndexes, cursorx)

	if len(nlAfterCursor) == 0 {
		// cursor is already on last line
		return
	}
	beginningOfLineAfterCursor := nlAfterCursor[0] + 1
	beginningOfLineBeforeCursor := 0
	if len(nlBeforeCursor) > 0 {
		beginningOfLineBeforeCursor = nlBeforeCursor[len(nlBeforeCursor)-1] + 1
	}

	nCharsBeforeCursorOnThisLine := cursorx - beginningOfLineBeforeCursor
	lengthOfNextLine := len(text) - (nlAfterCursor[0] + 1)
	if len(nlAfterCursor) > 1 {
		lengthOfNextLine = nlAfterCursor[1] - nlAfterCursor[0]
	}

	cursorXOnLine := nCharsBeforeCursorOnThisLine
	if cursorXOnLine >= lengthOfNextLine {
		cursorXOnLine = lengthOfNextLine - 1
	}
	newCursorX := beginningOfLineAfterCursor + cursorXOnLine

	tv.document.SetCursorX(newCursorX)
}

func (tv *TextView) moveCursorUp() {
	text := tv.document.GetTextContentNullTerminating()
	cursorx := tv.document.GetCursorX()
	newLineIndexes := runeSliceIndexes(text, '\n')
	if len(newLineIndexes) < 1 {
		// can't move across lines where there is only one!
		return
	}
	nlBeforeCursor, _ := splitIntSliceBetween(newLineIndexes, cursorx)
	if len(nlBeforeCursor) == 0 {
		// cursor is already on first line
		return
	}
	beginningOfThisLine := nlBeforeCursor[len(nlBeforeCursor)-1] + 1
	beginningOfLastLine := 0
	if len(nlBeforeCursor) > 1 {
		beginningOfLastLine = nlBeforeCursor[len(nlBeforeCursor)-2] + 1
	}
	cursorXOnThisLine := cursorx - beginningOfThisLine
	lengthOfPrevLine := beginningOfThisLine - beginningOfLastLine

	xoff := cursorXOnThisLine
	if cursorXOnThisLine > lengthOfPrevLine {
		xoff = lengthOfPrevLine - 1
	}
	newx := beginningOfLastLine + xoff
	tv.document.SetCursorX(newx)
}

func (tv *TextView) deleteUnderCursor() {
	tv.document.RemoveRuneUnderCursor()
}

func runeSliceIndexes(runes []rune, delimiter rune) []int {
	result := make([]int, 0)
	for i, r := range runes {
		if r == delimiter {
			result = append(result, i)
		}
	}
	return result
}

// assumes the slice is ordered
// note if an item in the array == the between integer, that one goes in the left slice
func splitIntSliceBetween(ints []int, between int) ([]int, []int) {
	before := ints
	after := ints[len(ints):]

	for i, integer := range ints {
		if integer >= between {
			before = ints[:i]
			after = ints[i:]
			break
		}
	}
	return before, after
}
