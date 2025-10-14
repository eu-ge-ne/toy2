package std

import (
	"bytes"
	"iter"
	"slices"
)

func IterToStr(itr iter.Seq[[]byte]) string {
	return string(bytes.Join(slices.Collect(itr), []byte{}))
}
