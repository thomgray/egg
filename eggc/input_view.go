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
	prompt            string
	placeholder       string
}

func MakeInputView() *InputView {
	iv := InputView{}
	iv.View = egg.MakeView()
	iv.document = MakeDocument()
	iv.OnKeyEvent(iv.handleKeyEvent)
	iv.OnDraw(iv.draw)
	iv.prompt = ""

	return &iv
}

func (iv *InputView) GetView() *egg.View {
	return iv.View
}

// OnEnter - provide a function called when user hits enter.
// The function input string `cmd` is the text content of the input
func (iv *InputView) OnEnter(f func(cmd string)) {
	iv.commandHandler = f
}

// Suggest - given the parameter input, set a command suggestion.
// this suggestion will be set if the user hits tab
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
	case egg.KeyLeft:
		iv.moveCursorBackwards()
	case egg.KeyRight:
		iv.moveCursorForward()
	case egg.KeyHome:
		iv.moveCursorToBeginningOfLine()
	case egg.KeyEnd:
		iv.moveCursorToEndOfLine()
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
	promptLen := runewidth.StringWidth(iv.prompt)
	c.DrawCursor(iv.document.GetCursorX()+promptLen, 0)
	content := iv.document.GetTextContent()
	c.DrawString2(iv.prompt, 0, 0)
	c.DrawString2(string(content), promptLen, 0)
	if iv.currentSuggestion != nil {
		w := runewidth.StringWidth(string(content)) + promptLen
		c.DrawString(*iv.currentSuggestion, w, 0, egg.ColorBlue, c.Background, c.Attribute)
	} else if len(content) == 0 {
		// render placeholder
		w := promptLen
		c.DrawString(iv.placeholder, w, 0, egg.ColorYellow, c.Background, c.Attribute)
	}
}

// GetTextContent - get the text content of this view as a byte slice
func (iv *InputView) GetTextContent() []rune {
	return iv.document.GetTextContent()
}

// SetPrompt ...
func (iv *InputView) SetPrompt(s string) {
	iv.prompt = s
}

// SetPlaceholder ...
func (iv *InputView) SetPlaceholder(s string) {
	iv.placeholder = s
}

// Clear - clear the input
func (iv *InputView) Clear() {
	iv.document.SetCursorX(0)
	iv.document.SetTextContentString("")
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
