package theme

import (
	"github.com/eu-ge-ne/toy2/internal/colors"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Slate struct {
}

func (Slate) DangerBg() []byte {
	return vt.CharBg(colors.Red900)
}

func (Slate) MainBg() []byte {
	return vt.CharBg(colors.Slate900)
}

func (Slate) MainFg() []byte {
	return vt.CharFg(colors.Slate900)
}

func (Slate) Light2Bg() []byte {
	return vt.CharBg(colors.Slate500)
}

func (Slate) Light2Fg() []byte {
	return vt.CharFg(colors.Slate100)
}

func (Slate) Light1Bg() []byte {
	return vt.CharBg(colors.Slate700)
}

func (Slate) Light1Fg() []byte {
	return vt.CharFg(colors.Slate200)
}

func (Slate) Light0Bg() []byte {
	return vt.CharBg(colors.Slate800)
}

func (Slate) Light0Fg() []byte {
	return vt.CharFg(colors.Slate300)
}

func (Slate) Dark0Bg() []byte {
	return vt.CharBg(colors.Slate950)
}

func (Slate) Dark0Fg() []byte {
	return vt.CharFg(colors.Slate400)
}

func (Slate) Dark1Fg() []byte {
	return vt.CharFg(colors.Slate600)
}

func (Slate) Dark2Fg() []byte {
	return vt.CharFg(colors.Slate700)
}
