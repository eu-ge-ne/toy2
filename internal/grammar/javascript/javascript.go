package javascript

import (
	_ "embed"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterJs "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
)

//go:embed highlights.scm
var ScmHighlight string

var Grammar JavaScript

type JavaScript struct {
	lang  *treeSitter.Language
	query *treeSitter.Query
}

func (grm JavaScript) Lang() *treeSitter.Language {
	return grm.lang
}

func (grm JavaScript) Query() *treeSitter.Query {
	return grm.query
}

func init() {
	lang := treeSitter.NewLanguage(treeSitterJs.Language())

	query, err0 := treeSitter.NewQuery(lang, ScmHighlight)
	if err0 != nil {
		panic(err0)
	}

	Grammar = JavaScript{lang, query}
}
