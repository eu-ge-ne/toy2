package typescript

import (
	_ "embed"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/grammar/javascript"
)

//go:embed highlights.scm
var ScmHighlight string

var Grammar TypeScript

type TypeScript struct {
	lang  *treeSitter.Language
	query *treeSitter.Query
}

func (grm TypeScript) Lang() *treeSitter.Language {
	return grm.lang
}

func (grm TypeScript) Query() *treeSitter.Query {
	return grm.query
}

func init() {
	lang := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())

	query, err0 := treeSitter.NewQuery(lang, javascript.ScmHighlight+ScmHighlight)
	if err0 != nil {
		panic(err0)
	}

	Grammar = TypeScript{lang, query}
}
