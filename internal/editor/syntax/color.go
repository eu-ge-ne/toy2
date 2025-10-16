package syntax

import (
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Colors struct {
	Char map[CharColor][]byte
}

type CharColor int

const (
	CharColorUndefined CharColor = iota
	CharColorVisible
	CharColorWhitespace
	CharColorEmpty
	CharColorDelimiter
)

func NewColors(t theme.Tokens) Colors {
	return Colors{
		Char: map[CharColor][]byte{
			CharColorVisible:    append(t.MainBg(), t.Light1Fg()...),
			CharColorWhitespace: append(t.MainBg(), t.Dark0Fg()...),
			CharColorEmpty:      append(t.MainBg(), t.MainFg()...),
			CharColorDelimiter:  append(t.MainBg(), vt.CharFg(theme.Red_900)...),
		},
	}
}

func WhitespaceCharColor(isVisible, whitespaceEnabled bool) CharColor {
	if isVisible {
		return CharColorVisible
	} else if whitespaceEnabled {
		return CharColorWhitespace
	} else {
		return CharColorEmpty
	}
}
