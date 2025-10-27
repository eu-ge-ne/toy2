package syntax

import (
	"os"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Syntax struct {
	parser

	grammar grammar.Grammar

	log  *os.File
	text []byte
	//started time.Time
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

func (s *Syntax) Highlight(buf *textbuf.TextBuf, startLn, endLn int) *Highlight {
	if s.grammar == nil {
		return nil
	}

	//s.started = time.Now()

	startPos, _ := buf.Pos(startLn, 0)
	endPos := buf.EndPos(endLn, 0)

	//fmt.Fprintf(s.log, "[%v] reset %v:%v\n", time.Since(s.started), startPos, endPos)

	hl := &Highlight{
		spans: make(chan span, 1024*2),
		span:  span{startIdx: -1, endIdx: -1},
		idx:   startPos.Idx,
	}

	go func() {
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
				hl.spans <- spn
				spn = span{startIdx, endIdx, name}
			} else {
				spn.name = name
			}
		}

		hl.spans <- spn

		close(hl.spans)

		//fmt.Fprintf(s.log, "[%v] done\n", time.Since(s.started))
	}()

	return hl
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
