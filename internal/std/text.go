package std

import (
	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func MeasureText(text string) (count, eolCount, lastEolIndex int) {
	gg := uniseg.NewGraphemes(text)

	for gg.Next() {
		g := grapheme.Graphemes.Get(gg.Str())

		if g.IsEol {
			eolCount += 1
			lastEolIndex = count
		}

		count += 1
	}

	return
}
