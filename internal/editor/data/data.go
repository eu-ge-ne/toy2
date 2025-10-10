package data

import (
	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Data struct {
	buffer  *textbuf.TextBuf
	cursor  *cursor.Cursor
	history *history.History
	syntax  *syntax.Syntax
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor, history *history.History) Data {
	return Data{
		buffer:  buffer,
		cursor:  cursor,
		history: history,
	}
}

func (d *Data) SetSyntax(syntax *syntax.Syntax) {
	d.syntax = syntax
}

func (d *Data) DeleteChar() {
	cur := d.cursor

	d.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	d.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)

	d.history.Push()
}

func (d *Data) DeletePrevChar() {
	cur := d.cursor

	if cur.Ln > 0 && cur.Col == 0 {
		l := 0
		for range d.buffer.IterLine(cur.Ln, false) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			d.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			d.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			cur.Left(false)
		} else {
			cur.Left(false)
			d.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			d.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		}
	} else {
		d.buffer.Delete2(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		d.syntax.Delete(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		cur.Left(false)
	}

	d.history.Push()
}

func (d *Data) DeleteSelection() {
	cur := d.cursor

	d.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	d.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	d.cursor.Set(cur.StartLn, cur.StartCol, false)

	d.history.Push()
}

func (d *Data) InsertText(text string) {
	cur := d.cursor

	if cur.Selecting {
		d.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		d.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.StartLn, cur.StartCol, false)
	}

	d.buffer.Insert2(cur.Ln, cur.Col, text)

	startLn := cur.Ln
	startCol := cur.Col

	dLn, dCol := grapheme.Graphemes.MeasureText(text)
	cur.Forward(dLn, dCol)

	endLn := cur.Ln
	endCol := cur.Col

	d.syntax.Insert(startLn, startCol, endLn, endCol)

	d.history.Push()
}
