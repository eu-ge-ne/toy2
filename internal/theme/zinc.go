package theme

import (
	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Zinc struct {
}

func (Zinc) DangerBg() []byte {
	return vt.CharBg(color.Red900)
}

func (Zinc) MainBg() []byte {
	return vt.CharBg(color.Zinc900)
}

func (Zinc) MainFg() []byte {
	return vt.CharFg(color.Zinc900)
}

func (Zinc) Light2Bg() []byte {
	return vt.CharBg(color.Zinc500)
}

func (Zinc) Light2Fg() []byte {
	return vt.CharFg(color.Zinc100)
}

func (Zinc) Light1Bg() []byte {
	return vt.CharBg(color.Zinc700)
}

func (Zinc) Light1Fg() []byte {
	return vt.CharFg(color.Zinc200)
}

func (Zinc) Light0Bg() []byte {
	return vt.CharBg(color.Zinc800)
}

func (Zinc) Light0Fg() []byte {
	return vt.CharFg(color.Zinc300)
}

func (Zinc) Dark0Bg() []byte {
	return vt.CharBg(color.Zinc950)
}

func (Zinc) Dark0Fg() []byte {
	return vt.CharFg(color.Zinc400)
}

func (Zinc) Dark1Fg() []byte {
	return vt.CharFg(color.Zinc600)
}

func (Zinc) Dark2Fg() []byte {
	return vt.CharFg(color.Zinc700)
}
