package grammar

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

type Grammar interface {
	Lang() *treeSitter.Language
	Query() *treeSitter.Query
}
