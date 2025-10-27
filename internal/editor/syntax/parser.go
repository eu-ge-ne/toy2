package syntax

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type parser struct {
	parser *treeSitter.Parser
	tree   *treeSitter.Tree
}

func (p *parser) Delete(change textbuf.Change) {
	if p.tree == nil {
		return
	}

	e := treeSitter.InputEdit{
		StartByte:  uint(change.Start.Idx),
		OldEndByte: uint(change.End.Idx),

		StartPosition:  treeSitter.NewPoint(uint(change.Start.Ln), uint(change.Start.ColIdx)),
		OldEndPosition: treeSitter.NewPoint(uint(change.End.Ln), uint(change.End.ColIdx)),
	}

	e.NewEndByte = e.StartByte
	e.NewEndPosition = e.StartPosition

	p.tree.Edit(&e)

	//fmt.Fprintf(s.log, "delete: change %+v\n", change)
	//fmt.Fprintf(s.log, "delete: e %+v\n", e)
}

func (p *parser) Insert(change textbuf.Change) {
	if p.tree == nil {
		return
	}

	e := treeSitter.InputEdit{
		StartByte:  uint(change.Start.Idx),
		NewEndByte: uint(change.End.Idx),

		StartPosition:  treeSitter.NewPoint(uint(change.Start.Ln), uint(change.Start.ColIdx)),
		NewEndPosition: treeSitter.NewPoint(uint(change.End.Ln), uint(change.End.ColIdx)),
	}

	e.OldEndByte = e.StartByte
	e.OldEndPosition = e.StartPosition

	p.tree.Edit(&e)

	//fmt.Fprintf(s.log, "insert: change %+v\n", change)
	//fmt.Fprintf(s.log, "insert: e %+v\n", e)
}

func (p *parser) initParser(grm grammar.Grammar) {
	p.parser = treeSitter.NewParser()

	err := p.parser.SetLanguage(grm.Lang())
	if err != nil {
		panic(err)
	}
}

func (p *parser) closeParser() {
	if p.tree != nil {
		p.tree.Close()
		p.tree = nil
	}

	if p.parser != nil {
		p.parser.Close()
		p.parser = nil
	}
}

const maxChunkLen = 1024 * 4

func (p *parser) parse(buf *textbuf.TextBuf, start, endPos textbuf.Pos) {
	startPos, _ := buf.Pos(max(0, start.Ln-1_000), 0)

	p.parser.SetIncludedRanges([]treeSitter.Range{{
		StartByte:  uint(startPos.Idx),
		EndByte:    uint(endPos.Idx),
		StartPoint: treeSitter.NewPoint(uint(startPos.Ln), uint(startPos.ColIdx)),
		EndPoint:   treeSitter.NewPoint(uint(endPos.Ln), uint(endPos.ColIdx)),
	}})

	oldTree := p.tree

	p.tree = p.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := buf.Chunk(i)

		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}

		//fmt.Fprintf(s.log, "[%v] reading chunk %d, %+v, %d\n", time.Since(s.started), i, p, len(text))

		return []byte(text)
	}, oldTree, nil)

	//fmt.Fprintf(s.log, "[%v] parsed\n", time.Since(s.started))
}
