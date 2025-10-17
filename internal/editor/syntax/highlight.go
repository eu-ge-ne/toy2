package syntax

import (
	"fmt"
	"os"
	"time"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

type Highlight struct {
	buffer *textbuf.TextBuf
	tree   *treeSitter.Tree
	rang   treeSitter.Range

	text  []byte
	query *treeSitter.Query
	qc    *treeSitter.QueryCursor
	capts treeSitter.QueryCaptures

	started time.Time
	f       *os.File
}

func newHighlight(buffer *textbuf.TextBuf, tree *treeSitter.Tree, rang treeSitter.Range) *Highlight {
	started := time.Now()

	f, err := os.OpenFile("tmp/highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	start := int(rang.StartByte)
	end := int(rang.EndByte)
	text := make([]byte, buffer.Count())
	chunk := std.IterToStr(buffer.Read(start, end))
	copy(text[start:end], chunk)

	query, err0 := treeSitter.NewQuery(tree.Language(), scmJsHighlights+scmTsHighlights)
	if err0 != nil {
		panic(err0)
	}

	qc := treeSitter.NewQueryCursor()
	qc.SetPointRange(rang.StartPoint, rang.EndPoint)

	capts := qc.Captures(query, tree.RootNode(), text)

	return &Highlight{
		buffer: buffer,
		tree:   tree,
		rang:   rang,

		text:  text,
		query: query,
		qc:    qc,
		capts: capts,

		started: started,
		f:       f,
	}
}

func (h *Highlight) Close() {
	if h == nil {
		return
	}

	h.qc.Close()

	fmt.Fprintf(h.f, "Elapsed %v\n", time.Since(h.started))
	h.f.Close()
}

type ColorSpan struct {
	start int
	end   int
	color CharFgColor
}

func (h *Highlight) Next() (ColorSpan, bool) {
	match, captIdx := h.capts.Next()
	if match == nil {
		return ColorSpan{}, false
	}

	capt := match.Captures[captIdx]
	node := capt.Node

	span := ColorSpan{
		start: int(node.StartByte()),
		end:   int(node.EndByte()),
	}

	switch capt.Index {
	case 0:
		span.color = CharFgColorVariable
	case 18:
		span.color = CharFgColorKeyword
	default:
		span.color = CharFgColorUndefined
	}

	fmt.Fprintf(h.f,
		"%v:%v|%v:%v %s (%s, %d, %d)\n",
		node.StartByte(),
		node.EndByte(),
		node.StartPosition(),
		node.EndPosition(),
		node.Utf8Text(h.text),
		h.query.CaptureNames()[capt.Index],
		match.PatternIndex,
		capt.Index,
	)

	return span, true
}
