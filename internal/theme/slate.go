package theme

import (
	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Slate struct {
}

func (Slate) DangerBg() []byte {
	return vt.CharBg(color.Red900)
}

func (Slate) MainBg() []byte {
	return vt.CharBg(color.Slate900)
}

func (Slate) MainFg() []byte {
	return vt.CharFg(color.Slate900)
}

func (Slate) Light2Bg() []byte {
	return vt.CharBg(color.Slate500)
}

func (Slate) Light2Fg() []byte {
	return vt.CharFg(color.Slate100)
}

func (Slate) Light1Bg() []byte {
	return vt.CharBg(color.Slate700)
}

func (Slate) Light1Fg() []byte {
	return vt.CharFg(color.Slate200)
}

func (Slate) Light0Bg() []byte {
	return vt.CharBg(color.Slate800)
}

func (Slate) Light0Fg() []byte {
	return vt.CharFg(color.Slate300)
}

func (Slate) Dark0Bg() []byte {
	return vt.CharBg(color.Slate950)
}

func (Slate) Dark0Fg() []byte {
	return vt.CharFg(color.Slate400)
}

func (Slate) Dark1Fg() []byte {
	return vt.CharFg(color.Slate600)
}

func (Slate) Dark2Fg() []byte {
	return vt.CharFg(color.Slate700)
}
