package textbuf_test

import (
	"bytes"
	"iter"
	"slices"
)

func iterToStr(itr iter.Seq[[]byte]) string {
	return string(bytes.Join(slices.Collect(itr), []byte{}))
}
