package editor

import (
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type Colors struct {
	background []byte
	index      []byte
	void       []byte
	char       map[charColorEnum][]byte
}

type charColorEnum int

const (
	charColorUndefined charColorEnum = iota
	charColorVisible
	charColorWhitespace
	charColorEmpty
	charColorVisibleSelected
	charColorWhitespaceSelected
	charColorEmptySelected
)

func NewColors(t theme.Tokens) Colors {
	return Colors{
		background: t.MainBg(),
		index:      append(t.Light0Bg(), t.Dark0Fg()...),
		void:       t.Dark0Bg(),
		char: map[charColorEnum][]byte{
			charColorVisible:            append(t.MainBg(), t.Light1Fg()...),
			charColorWhitespace:         append(t.MainBg(), t.Dark0Fg()...),
			charColorEmpty:              append(t.MainBg(), t.MainFg()...),
			charColorVisibleSelected:    append(t.Light2Bg(), t.Light1Fg()...),
			charColorWhitespaceSelected: append(t.Light2Bg(), t.Dark1Fg()...),
			charColorEmptySelected:      append(t.Light2Bg(), t.Dark1Fg()...),
		},
	}
}

func newCharColor(isSelected, isVisible, whitespaceEnabled bool) charColorEnum {
	if isSelected {
		if isVisible {
			return charColorVisibleSelected
		} else if whitespaceEnabled {
			return charColorWhitespaceSelected
		} else {
			return charColorEmptySelected
		}
	} else {
		if isVisible {
			return charColorVisible
		} else if whitespaceEnabled {
			return charColorWhitespace
		} else {
			return charColorEmpty
		}
	}
}
