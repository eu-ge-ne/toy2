package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Neutral struct {
}

func (Neutral) DangerBg() []byte {
	return vt.CharBg(red_900)
}

func (Neutral) MainBg() []byte {
	return vt.CharBg(neutral_900)
}

func (Neutral) MainFg() []byte {
	return vt.CharFg(neutral_900)
}

func (Neutral) Light2Bg() []byte {
	return vt.CharBg(neutral_500)
}

func (Neutral) Light2Fg() []byte {
	return vt.CharFg(neutral_100)
}

func (Neutral) Light1Bg() []byte {
	return vt.CharBg(neutral_700)
}

func (Neutral) Light1Fg() []byte {
	return vt.CharFg(neutral_200)
}

func (Neutral) Light0Bg() []byte {
	return vt.CharBg(neutral_800)
}

func (Neutral) Light0Fg() []byte {
	return vt.CharFg(neutral_300)
}

func (Neutral) Dark0Bg() []byte {
	return vt.CharBg(neutral_950)
}

func (Neutral) Dark0Fg() []byte {
	return vt.CharFg(neutral_400)
}

func (Neutral) Dark1Fg() []byte {
	return vt.CharFg(neutral_600)
}

func (Neutral) Dark2Fg() []byte {
	return vt.CharFg(neutral_700)
}
