package theme

import (
	"github.com/eu-ge-ne/toy2/internal/colors"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Zinc struct {
}

func (Zinc) DangerBg() []byte {
	return vt.CharBg(colors.Red900)
}

func (Zinc) MainBg() []byte {
	return vt.CharBg(colors.Zinc900)
}

func (Zinc) MainFg() []byte {
	return vt.CharFg(colors.Zinc900)
}

func (Zinc) Light2Bg() []byte {
	return vt.CharBg(colors.Zinc500)
}

func (Zinc) Light2Fg() []byte {
	return vt.CharFg(colors.Zinc100)
}

func (Zinc) Light1Bg() []byte {
	return vt.CharBg(colors.Zinc700)
}

func (Zinc) Light1Fg() []byte {
	return vt.CharFg(colors.Zinc200)
}

func (Zinc) Light0Bg() []byte {
	return vt.CharBg(colors.Zinc800)
}

func (Zinc) Light0Fg() []byte {
	return vt.CharFg(colors.Zinc300)
}

func (Zinc) Dark0Bg() []byte {
	return vt.CharBg(colors.Zinc950)
}

func (Zinc) Dark0Fg() []byte {
	return vt.CharFg(colors.Zinc400)
}

func (Zinc) Dark1Fg() []byte {
	return vt.CharFg(colors.Zinc600)
}

func (Zinc) Dark2Fg() []byte {
	return vt.CharFg(colors.Zinc700)
}
