package data

import (
	"io"
	"math"
	"os"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Data struct {
	multiLine bool
	buffer    *textbuf.TextBuf
	cursor    *cursor.Cursor
	history   *history.History
	syntax    *syntax.Syntax

	Handlers []Handler

	enabled   bool
	pageSize  int
	clipboard string
}

func New(multiLine bool, buffer *textbuf.TextBuf, cursor *cursor.Cursor, history *history.History) *Data {
	d := &Data{
		multiLine: multiLine,
		buffer:    buffer,
		cursor:    cursor,
		history:   history,
	}

	d.Handlers = append(d.Handlers,
		&Insert{d},
		&Backspace{d},
		&Bottom{d},
		&Copy{d},
		&Cut{d},
		&Delete{d},
		&GoDown{d},
		&GoEnd{d},
		&GoHome{d},
		&Enter{d},
		&Left{d},
		&PageDown{d},
		&PageUp{d},
		&Paste{d},
		&Redo{d},
		&Right{d},
		&SelectAll{d},
		&Top{d},
		&Undo{d},
		&Up{d},
	)

	return d
}

func (d *Data) SetSyntax() {
	d.syntax = syntax.New(d.buffer)
	d.syntax.Reset()
}

func (d *Data) SetEnabled(enabled bool) {
	d.enabled = enabled
}

func (d *Data) SetPageSize(pageSize int) {
	d.pageSize = pageSize
}

func (d *Data) HasChanges() bool {
	return !d.history.IsEmpty()
}

func (d *Data) CursorStatus() (int, int, int) {
	return d.cursor.Ln, d.cursor.Col, d.buffer.LineCount()
}

func (d *Data) SetText(text string) {
	d.buffer.Reset(text)
	d.syntax.Reset()
}

func (d *Data) GetText() string {
	return d.buffer.All()
}

func (d *Data) Load(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	buf := make([]byte, 1024*1024*64)

	for {
		bytesRead, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunk := buf[:bytesRead]

		if !utf8.Valid(chunk) {
			panic("invalid utf8 chunk")
		}

		d.buffer.Append(string(chunk))
	}

	d.syntax.Reset()

	return nil
}

func (d *Data) Save(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	for text := range d.buffer.Iter() {
		_, err := f.WriteString(text)
		if err != nil {
			return err
		}
	}

	return nil
}

// --

func (d *Data) Backspace() bool {
	if d.cursor.Selecting {
		d.deleteSelection()
	} else {
		d.deletePrevChar()
	}

	return true
}

func (d *Data) Bottom(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Bottom(sel)
}

func (d *Data) Copy() bool {
	if !d.enabled {
		return false
	}

	cur := d.cursor

	if cur.Selecting {
		d.clipboard = d.buffer.Read2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.Ln, cur.Col, false)
	} else {
		d.clipboard = d.buffer.Read2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	}

	vt.CopyToClipboard(vt.Sync, d.clipboard)

	return false
}

func (d *Data) Cut() bool {
	if !d.enabled {
		return false
	}

	cur := d.cursor

	if cur.Selecting {
		d.clipboard = d.buffer.Read2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		d.deleteSelection()
	} else {
		d.clipboard = d.buffer.Read2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		d.deleteChar()
	}

	vt.CopyToClipboard(vt.Sync, d.clipboard)

	return true
}

func (d *Data) Delete() bool {
	if d.cursor.Selecting {
		d.deleteSelection()
	} else {
		d.deleteChar()
	}

	return true
}

func (d *Data) Enter() bool {
	if !d.multiLine {
		return false
	}

	return d.Insert("\n")
}

func (d *Data) Insert(text string) bool {
	d.insertText(text)

	return true
}

func (d *Data) Left(sel bool) bool {
	return d.cursor.Left(sel)
}

func (d *Data) PageDown(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Down(d.pageSize, sel)
}

func (d *Data) PageUp(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Up(d.pageSize, sel)
}

func (d *Data) Paste() bool {
	if !d.enabled {
		return false
	}

	if len(d.clipboard) == 0 {
		return false
	}

	return d.Insert(d.clipboard)
}

func (d *Data) Redo() bool {
	if !d.enabled {
		return false
	}

	return d.history.Redo()
}

func (d *Data) Right(sel bool) bool {
	return d.cursor.Right(sel)
}

func (d *Data) SelectAll() bool {
	if !d.enabled {
		return false
	}

	d.cursor.Set(0, 0, false)
	d.cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}

func (d *Data) Top(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Top(sel)
}

func (d *Data) Undo() bool {
	if !d.enabled {
		return false
	}

	return d.history.Undo()
}

func (d *Data) Up(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Up(1, sel)
}

func (d *Data) TopHome(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Set(0, 0, false)
}

func (d *Data) deleteChar() {
	cur := d.cursor

	d.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	d.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)

	d.history.Push()
}

func (d *Data) deletePrevChar() {
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

func (d *Data) deleteSelection() {
	cur := d.cursor

	d.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	d.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	d.cursor.Set(cur.StartLn, cur.StartCol, false)

	d.history.Push()
}

func (d *Data) insertText(text string) {
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
