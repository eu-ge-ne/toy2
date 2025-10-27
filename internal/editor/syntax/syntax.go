package syntax

import (
	"fmt"
	"os"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Syntax struct {
	parser

	grammar grammar.Grammar

	spans chan span
	span  span
	idx   int

	log     *os.File
	text    []byte
	started time.Time
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func New() *Syntax {
	s := Syntax{}

	s.initLogger()

	return &s
}

func (s *Syntax) SetGrammar(grm grammar.Grammar) {
	s.grammar = grm

	s.closeParser()

	if grm != nil {
		s.initParser(grm)
	}
}

func (s *Syntax) Highlight(buf *textbuf.TextBuf, startLn, endLn int) {
	if s.grammar == nil {
		return
	}

	s.started = time.Now()

	startPos, _ := buf.Pos(startLn, 0)
	endPos := buf.EndPos(endLn, 0)

	fmt.Fprintf(s.log, "[%v] reset %v:%v\n", time.Since(s.started), startPos, endPos)

	s.spans = make(chan span, 1024*2)
	s.span = span{startIdx: -1, endIdx: -1}
	s.idx = startPos.Idx

	go s.highlight(buf, startPos, endPos)
}

func (s *Syntax) Next(l int) string {
	defer func() { s.idx += l }()

	if s.grammar == nil {
		return ""
	}

	if s.idx >= s.span.endIdx {
		if spn, ok := <-s.spans; ok {
			s.span = spn
		}
	}

	if s.idx >= s.span.startIdx && s.idx < s.span.endIdx {
		return s.span.name
	}

	return ""
}

func (s *Syntax) highlight(buf *textbuf.TextBuf, startPos, endPos textbuf.Pos) {
	query := s.grammar.Query()

	s.parse(buf, startPos, endPos)
	s.prepareText(buf, startPos, endPos)

	qc := treeSitter.NewQueryCursor()
	qc.SetByteRange(uint(startPos.Idx), uint(endPos.Idx))
	defer qc.Close()

	var spn span

	capts := qc.Captures(query, s.tree.RootNode(), s.text)

	match, captIdx := capts.Next()
	if match != nil {
		capt := match.Captures[captIdx]
		spn = span{int(capt.Node.StartByte()), int(capt.Node.EndByte()), query.CaptureNames()[capt.Index]}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := query.CaptureNames()[capt.Index]
		startIdx := int(capt.Node.StartByte())
		endIdx := int(capt.Node.EndByte())

		//fmt.Fprintf(s.log, "highlight: %v:%v %s (%s)\n", capt.Node.StartPosition(), capt.Node.EndPosition(), capt.Node.Utf8Text(s.text), name /*match.PatternIndex,*/ /*capt.Index,*/)

		if spn.startIdx != startIdx || spn.endIdx != endIdx {
			s.spans <- spn
			spn = span{startIdx, endIdx, name}
		} else {
			spn.name = name
		}
	}

	s.spans <- spn

	close(s.spans)

	fmt.Fprintf(s.log, "[%v] done\n", time.Since(s.started))
}

func (s *Syntax) prepareText(buf *textbuf.TextBuf, startPos, endPos textbuf.Pos) {
	if buf.Count() > len(s.text) {
		s.text = make([]byte, buf.Count())
	}

	copy(
		s.text[startPos.Idx:endPos.Idx],
		std.IterToStr(buf.Slice(startPos.Idx, endPos.Idx)),
	)
}

func (s *Syntax) initLogger() {
	log, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	/*
		i := 0

		s.parser.SetLogger(func(t treeSitter.LogType, msg string) {
			var tp string

			switch t {
			case treeSitter.LogTypeParse:
				tp = "parse"
			case treeSitter.LogTypeLex:
				tp = "lex"
			}

			fmt.Fprintf(log, "%d: %s: %s\n", i, tp, msg)

			i += 1
		})
	*/

	s.log = log
}
