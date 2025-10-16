package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Stone struct {
}

func (Stone) DangerBg() []byte {
	return vt.CharBg(Red_900)
}

func (Stone) MainBg() []byte {
	return vt.CharBg(stone_900)
}

func (Stone) MainFg() []byte {
	return vt.CharFg(stone_900)
}

func (Stone) Light2Bg() []byte {
	return vt.CharBg(stone_500)
}

func (Stone) Light2Fg() []byte {
	return vt.CharFg(stone_100)
}

func (Stone) Light1Bg() []byte {
	return vt.CharBg(stone_700)
}

func (Stone) Light1Fg() []byte {
	return vt.CharFg(stone_200)
}

func (Stone) Light0Bg() []byte {
	return vt.CharBg(stone_800)
}

func (Stone) Light0Fg() []byte {
	return vt.CharFg(stone_300)
}

func (Stone) Dark0Bg() []byte {
	return vt.CharBg(stone_950)
}

func (Stone) Dark0Fg() []byte {
	return vt.CharFg(stone_400)
}

func (Stone) Dark1Fg() []byte {
	return vt.CharFg(stone_600)
}

func (Stone) Dark2Fg() []byte {
	return vt.CharFg(stone_700)
}
