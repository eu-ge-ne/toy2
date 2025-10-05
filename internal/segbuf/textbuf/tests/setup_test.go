package textbuf_test

import (
	"iter"
	"slices"
	"strings"
)

func iterToStr(itr iter.Seq[string]) string {
	return strings.Join(slices.Collect(itr), "")
}
