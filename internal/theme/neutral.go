package theme

import (
	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Neutral struct {
}

func (Neutral) DangerBg() []byte {
	return vt.CharBg(color.Red900)
}

func (Neutral) MainBg() []byte {
	return vt.CharBg(color.Neutral900)
}

func (Neutral) MainFg() []byte {
	return vt.CharFg(color.Neutral900)
}

func (Neutral) Light2Bg() []byte {
	return vt.CharBg(color.Neutral500)
}

func (Neutral) Light2Fg() []byte {
	return vt.CharFg(color.Neutral100)
}

func (Neutral) Light1Bg() []byte {
	return vt.CharBg(color.Neutral700)
}

func (Neutral) Light1Fg() []byte {
	return vt.CharFg(color.Neutral200)
}

func (Neutral) Light0Bg() []byte {
	return vt.CharBg(color.Neutral800)
}

func (Neutral) Light0Fg() []byte {
	return vt.CharFg(color.Neutral300)
}

func (Neutral) Dark0Bg() []byte {
	return vt.CharBg(color.Neutral950)
}

func (Neutral) Dark0Fg() []byte {
	return vt.CharFg(color.Neutral400)
}

func (Neutral) Dark1Fg() []byte {
	return vt.CharFg(color.Neutral600)
}

func (Neutral) Dark2Fg() []byte {
	return vt.CharFg(color.Neutral700)
}
