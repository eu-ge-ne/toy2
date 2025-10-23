package syntax

import (
	"fmt"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Highlighter struct {
	syntax        *Syntax
	startPos      textbuf.Pos
	startPosParse textbuf.Pos
	endPos        textbuf.Pos
	spans         chan span
	span          span
	idx           int

	text []byte
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func NewHighlighter(syntax *Syntax, startLn, endLn int) *Highlighter {
	startPos, _ := syntax.buffer.Pos(startLn, 0)
	endPos := syntax.buffer.EndPos(endLn, 0)
	startPosParse, _ := syntax.buffer.Pos(max(0, startLn-1000), 0)

	h := &Highlighter{
		syntax:        syntax,
		startPos:      startPos,
		startPosParse: startPosParse,
		endPos:        endPos,
		spans:         make(chan span, 1024),
		span:          span{-1, -1, ""},
		idx:           startPos.Idx,
	}

	go h.highlight()

	return h
}

func (h *Highlighter) Next(l int) string {
	var name string

	if h.idx >= h.span.endIdx {
		if s, ok := <-h.spans; ok {
			h.span = s
		}
	}

	if h.idx >= h.span.startIdx && h.idx < h.span.endIdx {
		name = h.span.name
	}

	h.idx += l

	return name
}

const maxChunkLen = 1024 * 64

func (h *Highlighter) highlight() {
	started := time.Now()

	fmt.Fprintln(h.syntax.log, "highlight: started")

	h.syntax.parse(h.startPosParse, h.endPos)

	fmt.Fprintf(h.syntax.log, "highlight: parsed %v\n", time.Since(started))

	h.prepareText()

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(h.startPos.Idx), uint(h.endPos.Idx))

	var spn span

	capts := qc.Captures(h.syntax.query, h.syntax.tree.RootNode(), h.text)

	match, captIdx := capts.Next()
	if match != nil {
		capt := match.Captures[captIdx]
		spn = span{
			int(capt.Node.StartByte()),
			int(capt.Node.EndByte()),
			h.syntax.query.CaptureNames()[capt.Index],
		}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := h.syntax.query.CaptureNames()[capt.Index]

		/*
			fmt.Fprintf(s.log,
				"highlight: %v:%v %s (%s)\n",
				capt.Node.StartPosition(),
				capt.Node.EndPosition(),
				capt.Node.Utf8Text(s.text),
				name,
				//match.PatternIndex,
				//capt.Index,
			)
		*/

		startIdx := int(capt.Node.StartByte())
		endIdx := int(capt.Node.EndByte())

		if spn.startIdx != startIdx || spn.endIdx != endIdx {
			h.spans <- spn
			spn = span{startIdx, endIdx, name}
		} else {
			spn.name = name
		}
	}

	h.spans <- spn

	close(h.spans)

	fmt.Fprintf(h.syntax.log, "highlight: elapsed %v\n", time.Since(started))
}

func (h *Highlighter) prepareText() {
	if h.syntax.buffer.Count() > len(h.text) {
		h.text = make([]byte, h.syntax.buffer.Count())
	}

	copy(
		h.text[h.startPos.Idx:h.endPos.Idx],
		std.IterToStr(h.syntax.buffer.Slice(h.startPos.Idx, h.endPos.Idx)),
	)
}
