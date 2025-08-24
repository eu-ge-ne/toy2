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
	return vt.CharBg(neutral[900])
}

func (Neutral) MainFg() []byte {
	return vt.CharFg(neutral[900])
}

func (Neutral) Light2Bg() []byte {
	return vt.CharBg(neutral[500])
}

func (Neutral) Light2Fg() []byte {
	return vt.CharFg(neutral[100])
}

func (Neutral) Light1Bg() []byte {
	return vt.CharBg(neutral[700])
}

func (Neutral) Light1Fg() []byte {
	return vt.CharFg(neutral[200])
}

func (Neutral) Light0Bg() []byte {
	return vt.CharBg(neutral[800])
}

func (Neutral) Light0Fg() []byte {
	return vt.CharFg(neutral[300])
}

func (Neutral) Dark0Bg() []byte {
	return vt.CharBg(neutral[950])
}

func (Neutral) Dark0Fg() []byte {
	return vt.CharFg(neutral[400])
}

func (Neutral) Dark1Fg() []byte {
	return vt.CharFg(neutral[600])
}

func (Neutral) Dark2Fg() []byte {
	return vt.CharFg(neutral[700])
}
