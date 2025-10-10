package editor

import (
	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
	ed.syntax.Reset()
}

func (ed *Editor) GetText() string {
	return ed.buffer.All()
}

func (ed *Editor) LoadFromFile(filePath string) error {
	err := ed.buffer.LoadFromFile(filePath)
	if err != nil {
		return err
	}

	ed.syntax.Reset()

	return nil
}

func (ed *Editor) SaveToFile(filePath string) error {
	return ed.buffer.SaveToFile(filePath)
}

func (ed *Editor) deleteChar() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)

	ed.history.Push()
}

func (ed *Editor) deletePrevChar() {
	cur := ed.cursor

	if cur.Ln > 0 && cur.Col == 0 {
		l := 0
		for range ed.buffer.IterLine(cur.Ln, false) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			cur.Left(false)
		} else {
			cur.Left(false)
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		}
	} else {
		ed.buffer.Delete2(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		ed.syntax.Delete(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		cur.Left(false)
	}

	ed.history.Push()
}

func (ed *Editor) deleteSelection() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.cursor.Set(cur.StartLn, cur.StartCol, false)

	ed.history.Push()
}

func (ed *Editor) insertText(text string) {
	cur := ed.cursor

	if cur.Selecting {
		ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.StartLn, cur.StartCol, false)
	}

	ed.buffer.Insert2(cur.Ln, cur.Col, text)

	startLn := cur.Ln
	startCol := cur.Col

	dLn, dCol := grapheme.Graphemes.MeasureText(text)
	cur.Forward(dLn, dCol)

	endLn := cur.Ln
	endCol := cur.Col

	ed.syntax.Insert(startLn, startCol, endLn, endCol)

	ed.history.Push()
}
