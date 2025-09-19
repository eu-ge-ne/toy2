package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Zinc struct {
}

func (Zinc) DangerBg() []byte {
	return vt.CharBg(red_900)
}

func (Zinc) MainBg() []byte {
	return vt.CharBg(zinc_900)
}

func (Zinc) MainFg() []byte {
	return vt.CharFg(zinc_900)
}

func (Zinc) Light2Bg() []byte {
	return vt.CharBg(zinc_500)
}

func (Zinc) Light2Fg() []byte {
	return vt.CharFg(zinc_100)
}

func (Zinc) Light1Bg() []byte {
	return vt.CharBg(zinc_700)
}

func (Zinc) Light1Fg() []byte {
	return vt.CharFg(zinc_200)
}

func (Zinc) Light0Bg() []byte {
	return vt.CharBg(zinc_800)
}

func (Zinc) Light0Fg() []byte {
	return vt.CharFg(zinc_300)
}

func (Zinc) Dark0Bg() []byte {
	return vt.CharBg(zinc_950)
}

func (Zinc) Dark0Fg() []byte {
	return vt.CharFg(zinc_400)
}

func (Zinc) Dark1Fg() []byte {
	return vt.CharFg(zinc_600)
}

func (Zinc) Dark2Fg() []byte {
	return vt.CharFg(zinc_700)
}
