package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Gray struct {
}

func (Gray) DangerBg() []byte {
	return vt.CharBg(Red_900)
}

func (Gray) MainBg() []byte {
	return vt.CharBg(gray_900)
}

func (Gray) MainFg() []byte {
	return vt.CharFg(gray_900)
}

func (Gray) Light2Bg() []byte {
	return vt.CharBg(gray_500)
}

func (Gray) Light2Fg() []byte {
	return vt.CharFg(gray_100)
}

func (Gray) Light1Bg() []byte {
	return vt.CharBg(gray_700)
}

func (Gray) Light1Fg() []byte {
	return vt.CharFg(gray_200)
}

func (Gray) Light0Bg() []byte {
	return vt.CharBg(gray_800)
}

func (Gray) Light0Fg() []byte {
	return vt.CharFg(gray_300)
}

func (Gray) Dark0Bg() []byte {
	return vt.CharBg(gray_950)
}

func (Gray) Dark0Fg() []byte {
	return vt.CharFg(gray_400)
}

func (Gray) Dark1Fg() []byte {
	return vt.CharFg(gray_600)
}

func (Gray) Dark2Fg() []byte {
	return vt.CharFg(gray_700)
}
