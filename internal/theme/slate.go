package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Slate struct {
}

func (Slate) DangerBg() []byte {
	return vt.CharBg(Red_900)
}

func (Slate) MainBg() []byte {
	return vt.CharBg(slate_900)
}

func (Slate) MainFg() []byte {
	return vt.CharFg(slate_900)
}

func (Slate) Light2Bg() []byte {
	return vt.CharBg(slate_500)
}

func (Slate) Light2Fg() []byte {
	return vt.CharFg(slate_100)
}

func (Slate) Light1Bg() []byte {
	return vt.CharBg(slate_700)
}

func (Slate) Light1Fg() []byte {
	return vt.CharFg(slate_200)
}

func (Slate) Light0Bg() []byte {
	return vt.CharBg(slate_800)
}

func (Slate) Light0Fg() []byte {
	return vt.CharFg(slate_300)
}

func (Slate) Dark0Bg() []byte {
	return vt.CharBg(slate_950)
}

func (Slate) Dark0Fg() []byte {
	return vt.CharFg(slate_400)
}

func (Slate) Dark1Fg() []byte {
	return vt.CharFg(slate_600)
}

func (Slate) Dark2Fg() []byte {
	return vt.CharFg(slate_700)
}
