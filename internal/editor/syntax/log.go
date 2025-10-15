package syntax

import (
	_ "embed"
	"fmt"
	"os"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

func Log(parser *treeSitter.Parser) {
	f, err := os.OpenFile("tmp/syntax-trace.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	i := 0

	parser.SetLogger(func(t treeSitter.LogType, msg string) {
		var tp string

		switch t {
		case treeSitter.LogTypeParse:
			tp = "Parse"
		case treeSitter.LogTypeLex:
			tp = "Lex"
		}

		fmt.Fprintf(f, "%d: %s: %s\n", i, tp, msg)

		i += 1
	})
}
