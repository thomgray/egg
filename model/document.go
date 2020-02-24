package model

// Document - a text model
type Document struct {
	textContent []rune
	curx        int
}

// MakeDocument - constructor
func MakeDocument() *Document {
	return &Document{
		textContent: make([]rune, 0),
		curx:        0,
	}
}

func (d *Document) GetTextContent() []rune {
	return d.textContent
}

// GetTextContentNullTerminating - return the text content with a null rune terminus.
// This can be useful for traversing the content of the document when you want to
// guarantee a rune in the array at the index of the cursor.
func (d *Document) GetTextContentNullTerminating() []rune {
	tc := d.GetTextContent()
	return runeSliceInsert(tc, len(tc), '\000')

}

func (d *Document) SetTextContentString(tc string) {
	d.textContent = []rune(tc)
	if d.curx > len(tc) {
		d.curx = len(tc)
	}
}

func (d *Document) InsertRuneAtCursor(r rune) {
	d.textContent = runeSliceInsert(d.textContent, d.curx, r)
}

func (d *Document) InsertRuneAtCursorIncrementing(r rune) {
	d.InsertRuneAtCursor(r)
	d.curx++
}

func (d *Document) RemoveRuneUnderCursor() {
	d.textContent = runeSliceRemove(d.textContent, d.curx)
}

func (d *Document) SetCursorX(x int) {
	if x < 0 {
		d.curx = 0
	} else if x > len(d.textContent) {
		d.curx = len(d.textContent)
	} else {
		d.curx = x
	}
}

func (d *Document) GetCursorX() int {
	return d.curx
}
