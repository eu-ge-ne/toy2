package std

import (
	"bytes"
	"iter"
	"slices"
)

func IterToStr(itr iter.Seq[[]byte]) string {
	return string(bytes.Join(slices.Collect(itr), []byte{}))
}

func IterToBytes(itr iter.Seq[[]byte]) []byte {
	return bytes.Join(slices.Collect(itr), []byte{})
}
