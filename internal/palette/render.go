package palette

import (
	"fmt"
	"io"

	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (p *Palette) Render() {
	if !p.enabled {
		return
	}

	area := p.resize2()

	p.editor.Layout(ui.Area{
		Y: area.Y + 1,
		X: area.X + 2,
		W: area.W - 4,
		H: 1,
	})

	p.scroll()

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(p.colorBackground)
	vt.ClearArea(vt.Buf, area)

	if len(p.filteredOptions) == 0 {
		p.renderEmpty(area)
	} else {
		p.renderOptions(area)
	}

	p.editor.Render()

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (p *Palette) resize2() ui.Area {
	p.listSize = min(len(p.filteredOptions), maxListSize)

	area := ui.Area{}

	area.W = min(60, p.area.W)

	area.H = 3 + max(p.listSize, 1)
	if area.H > p.area.H {
		area.H = p.area.H
		if p.listSize > 0 {
			p.listSize = area.H - 3
		}
	}

	area.Y = p.area.Y + ((p.area.H - area.H) / 2)
	area.X = p.area.X + ((p.area.W - area.W) / 2)

	return area
}

func (p *Palette) scroll() {
	delta := p.selectedIndex - p.scrollIndex

	if delta < 0 {
		p.scrollIndex = p.selectedIndex
	} else if delta >= p.listSize {
		p.scrollIndex = p.selectedIndex - p.listSize + 1
	}
}

func (p *Palette) renderEmpty(area ui.Area) {
	vt.SetCursor(vt.Buf, area.Y+2, area.X+2)
	vt.Buf.Write(p.colorOption)
	io.WriteString(vt.Buf, "No matching commands")
}

func (p *Palette) renderOptions(area ui.Area) {
	i := 0
	y := area.Y + 2

	for {
		if i == p.listSize {
			break
		}

		index := p.scrollIndex + i
		option := p.filteredOptions[index]
		if option == nil {
			break
		}

		if y == area.Y+area.H {
			break
		}

		span := area.W - 4

		if index == p.selectedIndex {
			vt.Buf.Write(p.colorSelectedOption)
		} else {
			vt.Buf.Write(p.colorOption)
		}

		vt.SetCursor(vt.Buf, y, area.X+2)
		vt.WriteText(vt.Buf, &span, option.Description)
		vt.WriteText(vt.Buf, &span, fmt.Sprintf("%*s", span, option.shortcuts))

		i += 1
		y += 1
	}
}
