package theme

import (
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Base16 struct {
}

func (Base16) DangerBg() []byte {
	return vt.CharAttr(vt.BgRed)
}

func (Base16) MainBg() []byte {
	return vt.CharAttr(vt.BgBlack)
}

func (Base16) MainFg() []byte {
	return vt.CharAttr(vt.FgBlack)
}

func (Base16) Light2Bg() []byte {
	return vt.CharAttr(vt.BgCyan)
}

func (Base16) Light2Fg() []byte {
	return vt.CharAttr(vt.FgCyan)
}

func (Base16) Light1Bg() []byte {
	return vt.CharAttr(vt.BgBrightBlack)
}

func (Base16) Light1Fg() []byte {
	return vt.CharAttr(vt.FgBrightWhite)
}

func (Base16) Light0Bg() []byte {
	return vt.CharAttr(vt.BgBrightBlack)
}

func (Base16) Light0Fg() []byte {
	return vt.CharAttr(vt.FgBrightWhite)
}

func (Base16) Dark0Bg() []byte {
	return vt.CharAttr(vt.BgBlack)
}

func (Base16) Dark0Fg() []byte {
	return vt.CharAttr(vt.FgWhite)
}

func (Base16) Dark1Fg() []byte {
	return vt.CharAttr(vt.FgWhite)
}

func (Base16) Dark2Fg() []byte {
	return vt.CharAttr(vt.FgWhite)
}
