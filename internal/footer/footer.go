package footer

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Footer struct {
	area         ui.Area
	cursorStatus string

	colorBackground []byte
	colorText       []byte
}

func New() *Footer {
	return &Footer{}
}

func (f *Footer) SetCursorStatus(ln0, col0, lnCount int) {
	ln := ln0 + 1
	col := col0 + 1

	pct := 0
	if lnCount != 0 {
		pct = int((float64(ln) / float64(lnCount)) * 100)
	}

	f.cursorStatus = fmt.Sprintf("%d %d  %d%% ", ln, col, pct)

	f.Render()
}

func (f *Footer) SetColors(t theme.Tokens) {
	f.colorBackground = t.Dark0Bg()
	f.colorText = append(t.Dark0Bg(), t.Dark0Fg()...)
}
