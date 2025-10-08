package std

import (
	"iter"
	"slices"
	"strings"
)

func IterToStr(itr iter.Seq[string]) string {
	return string(strings.Join(slices.Collect(itr), ""))
}
