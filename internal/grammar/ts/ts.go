package ts

import (
	_ "embed"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/grammar/js"
)

//go:embed highlights.scm
var ScmHighlight string

var TS GrammarTS

type GrammarTS struct {
	lang  *treeSitter.Language
	query *treeSitter.Query
}

func (grm GrammarTS) Lang() *treeSitter.Language {
	return grm.lang
}

func (grm GrammarTS) Query() *treeSitter.Query {
	return grm.query
}

func init() {
	lang := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())

	query, err0 := treeSitter.NewQuery(lang, js.ScmHighlight+ScmHighlight)
	if err0 != nil {
		panic(err0)
	}

	TS = GrammarTS{lang, query}
}
