package theme

import (
	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Stone struct {
}

func (Stone) DangerBg() []byte {
	return vt.CharBg(color.Red900)
}

func (Stone) MainBg() []byte {
	return vt.CharBg(color.Stone900)
}

func (Stone) MainFg() []byte {
	return vt.CharFg(color.Stone900)
}

func (Stone) Light2Bg() []byte {
	return vt.CharBg(color.Stone500)
}

func (Stone) Light2Fg() []byte {
	return vt.CharFg(color.Stone100)
}

func (Stone) Light1Bg() []byte {
	return vt.CharBg(color.Stone700)
}

func (Stone) Light1Fg() []byte {
	return vt.CharFg(color.Stone200)
}

func (Stone) Light0Bg() []byte {
	return vt.CharBg(color.Stone800)
}

func (Stone) Light0Fg() []byte {
	return vt.CharFg(color.Stone300)
}

func (Stone) Dark0Bg() []byte {
	return vt.CharBg(color.Stone950)
}

func (Stone) Dark0Fg() []byte {
	return vt.CharFg(color.Stone400)
}

func (Stone) Dark1Fg() []byte {
	return vt.CharFg(color.Stone600)
}

func (Stone) Dark2Fg() []byte {
	return vt.CharFg(color.Stone700)
}
