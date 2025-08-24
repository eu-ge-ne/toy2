package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/content"
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/tree"
)

type TextBuf struct {
	content content.Content
	tree    tree.Tree
}

func New(text string) *TextBuf {
	buf := TextBuf{
		content: content.Content{},
		tree:    tree.Tree{Root: node.NIL},
	}

	if len(text) > 0 {
		root := buf.content.Create(text)
		buf.tree.Root = &root
		buf.tree.Root.Red = false
	}

	return &buf
}

func (buf *TextBuf) Count() int {
	return buf.tree.Root.TotalLen
}

func (buf *TextBuf) LineCount() int {
	if buf.tree.Root.TotalLen == 0 {
		return 0
	} else {
		return buf.tree.Root.TotalEolsLen + 1
	}
}

func (buf *TextBuf) Reset(text string) {
	buf.DeleteToEnd(0)

	if len(text) > 0 {
		buf.Insert(0, text)
	}
}

func (buf *TextBuf) Validate() {
	buf.tree.Root.Validate()
}
