package theme

import (
	"github.com/eu-ge-ne/toy2/internal/colors"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Neutral struct {
}

func (Neutral) DangerBg() []byte {
	return vt.CharBg(colors.Red900)
}

func (Neutral) MainBg() []byte {
	return vt.CharBg(colors.Neutral900)
}

func (Neutral) MainFg() []byte {
	return vt.CharFg(colors.Neutral900)
}

func (Neutral) Light2Bg() []byte {
	return vt.CharBg(colors.Neutral500)
}

func (Neutral) Light2Fg() []byte {
	return vt.CharFg(colors.Neutral100)
}

func (Neutral) Light1Bg() []byte {
	return vt.CharBg(colors.Neutral700)
}

func (Neutral) Light1Fg() []byte {
	return vt.CharFg(colors.Neutral200)
}

func (Neutral) Light0Bg() []byte {
	return vt.CharBg(colors.Neutral800)
}

func (Neutral) Light0Fg() []byte {
	return vt.CharFg(colors.Neutral300)
}

func (Neutral) Dark0Bg() []byte {
	return vt.CharBg(colors.Neutral950)
}

func (Neutral) Dark0Fg() []byte {
	return vt.CharFg(colors.Neutral400)
}

func (Neutral) Dark1Fg() []byte {
	return vt.CharFg(colors.Neutral600)
}

func (Neutral) Dark2Fg() []byte {
	return vt.CharFg(colors.Neutral700)
}
