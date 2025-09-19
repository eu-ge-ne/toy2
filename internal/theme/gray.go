package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Gray struct {
}

func (Gray) DangerBg() []byte {
	return vt.CharBg(red_900)
}

func (Gray) MainBg() []byte {
	return vt.CharBg(gray[900])
}

func (Gray) MainFg() []byte {
	return vt.CharFg(gray[900])
}

func (Gray) Light2Bg() []byte {
	return vt.CharBg(gray[500])
}

func (Gray) Light2Fg() []byte {
	return vt.CharFg(gray[100])
}

func (Gray) Light1Bg() []byte {
	return vt.CharBg(gray[700])
}

func (Gray) Light1Fg() []byte {
	return vt.CharFg(gray[200])
}

func (Gray) Light0Bg() []byte {
	return vt.CharBg(gray[800])
}

func (Gray) Light0Fg() []byte {
	return vt.CharFg(gray[300])
}

func (Gray) Dark0Bg() []byte {
	return vt.CharBg(gray[950])
}

func (Gray) Dark0Fg() []byte {
	return vt.CharFg(gray[400])
}

func (Gray) Dark1Fg() []byte {
	return vt.CharFg(gray[600])
}

func (Gray) Dark2Fg() []byte {
	return vt.CharFg(gray[700])
}
