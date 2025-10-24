package js

import (
	_ "embed"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterJs "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
)

//go:embed highlights.scm
var ScmHighlight string

var JS GrammarJS

type GrammarJS struct {
	lang  *treeSitter.Language
	query *treeSitter.Query
}

func (grm GrammarJS) Lang() *treeSitter.Language {
	return grm.lang
}

func (grm GrammarJS) Query() *treeSitter.Query {
	return grm.query
}

func init() {
	lang := treeSitter.NewLanguage(treeSitterJs.Language())

	query, err0 := treeSitter.NewQuery(lang, ScmHighlight)
	if err0 != nil {
		panic(err0)
	}

	JS = GrammarJS{lang, query}
}
