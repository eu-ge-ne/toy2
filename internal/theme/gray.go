package theme

import (
	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Gray struct {
}

func (Gray) DangerBg() []byte {
	return vt.CharBg(color.Red900)
}

func (Gray) MainBg() []byte {
	return vt.CharBg(color.Gray900)
}

func (Gray) MainFg() []byte {
	return vt.CharFg(color.Gray900)
}

func (Gray) Light2Bg() []byte {
	return vt.CharBg(color.Gray500)
}

func (Gray) Light2Fg() []byte {
	return vt.CharFg(color.Gray100)
}

func (Gray) Light1Bg() []byte {
	return vt.CharBg(color.Gray700)
}

func (Gray) Light1Fg() []byte {
	return vt.CharFg(color.Gray200)
}

func (Gray) Light0Bg() []byte {
	return vt.CharBg(color.Gray800)
}

func (Gray) Light0Fg() []byte {
	return vt.CharFg(color.Gray300)
}

func (Gray) Dark0Bg() []byte {
	return vt.CharBg(color.Gray950)
}

func (Gray) Dark0Fg() []byte {
	return vt.CharFg(color.Gray400)
}

func (Gray) Dark1Fg() []byte {
	return vt.CharFg(color.Gray600)
}

func (Gray) Dark2Fg() []byte {
	return vt.CharFg(color.Gray700)
}
