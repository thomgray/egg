package eggc

import (
	"github.com/mattn/go-runewidth"
	"github.com/thomgray/egg"
)

type InputView struct {
	*egg.View
	document          *Document
	suggest           func(string) (string, bool)
	commandHandler    func(string)
	currentSuggestion *string
}

func MakeInputView() *InputView {
	iv := InputView{}
	iv.View = egg.MakeView()
	iv.document = MakeDocument()
	iv.OnKeyEvent(iv.handleKeyEvent)
	iv.OnDraw(iv.draw)

	return &iv
}

func (iv *InputView) GetView() *egg.View {
	return iv.View
}

// OnCommandSent - function called with the input string when user hits enter
func (iv *InputView) OnCommandSent(f func(string)) {
	iv.commandHandler = f
}

// Suggest - given the parameter input, set a tab-completion suggestion
func (iv *InputView) Suggest(f func(string) (string, bool)) {
	iv.suggest = f
}

func (iv *InputView) handleKeyEvent(e *egg.KeyEvent) {
	switch e.Key {
	case egg.KeyBackspace | egg.KeyBackspace2:
		if iv.document.GetCursorX() > 0 {
			iv.moveCursorBackwards()
			iv.deleteUnderCursor()
		}
	case egg.KeyDelete:
		iv.deleteUnderCursor()
	case egg.KeyArrowLeft:
		iv.moveCursorBackwards()
	case egg.KeyArrowRight:
		iv.moveCursorForward()
	case egg.KeyHome:
		iv.moveCursorToBeginningOfLine()
	case egg.KeyEnd:
		iv.moveCursorToEndOfLine()
	case egg.KeySpace:
		iv.insertRune(' ')
	case egg.KeyEnter:
		if iv.commandHandler != nil {
			iv.commandHandler(string(iv.document.GetTextContent()))
		}
	case egg.KeyTab:
		iv.handleTab()
	default:
		if e.Char != 0 {
			iv.insertRune(e.Char)
		}
	}

	iv.evaluateSuggestion()
	iv.ReDraw()
}

func (iv *InputView) handleTab() {
	text := iv.document.GetTextContent()
	if iv.currentSuggestion != nil {
		newTxt := string(text) + *iv.currentSuggestion
		iv.document.SetTextContentString(newTxt)
		iv.document.SetCursorX(runewidth.StringWidth(newTxt))
		iv.currentSuggestion = nil
		return
	}
}

func (iv *InputView) deleteUnderCursor() {
	iv.document.RemoveRuneUnderCursor()
}

func (iv *InputView) moveCursorToBeginningOfLine() {
	iv.document.SetCursorX(0)
}

func (iv *InputView) moveCursorToEndOfLine() {
	iv.document.SetCursorX(len(iv.document.GetTextContent()))
}

func (iv *InputView) insertRune(r rune) {
	iv.document.InsertRuneAtCursorIncrementing(r)
}

func (iv *InputView) moveCursorForward() {
	iv.document.SetCursorX(iv.document.GetCursorX() + 1)
}

func (iv *InputView) moveCursorBackwards() {
	iv.document.SetCursorX(iv.document.GetCursorX() - 1)
}

func (iv *InputView) draw(c egg.Canvas) {
	c.DrawCursor(iv.document.GetCursorX(), 0)
	content := iv.document.GetTextContent()
	c.DrawString2(string(content), 0, 0)
	if iv.currentSuggestion != nil {
		w := len(content)
		c.DrawString(*iv.currentSuggestion, w, 0, egg.ColorBlue, c.Background, c.Attribute)
	}
}

// GetTextContent - get the text content of this view as a byte slice
func (iv *InputView) GetTextContent() []rune {
	return iv.document.GetTextContent()
}

// GetTextContentString - get the text content of this view as a string
func (iv *InputView) GetTextContentString() string {
	return string(iv.document.GetTextContent())
}

// SetTextContentString - set the text content of this view
func (iv *InputView) SetTextContentString(textContent string) {
	iv.document.SetTextContentString(textContent)
	iv.evaluateSuggestion()
}

func (iv *InputView) evaluateSuggestion() {
	str := string(iv.document.GetTextContent())
	iv.currentSuggestion = nil
	if iv.suggest != nil {
		if compl, doit := iv.suggest(str); doit {
			iv.currentSuggestion = &compl
		}
	}
}
