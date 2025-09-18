package palette

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (*Palette) Layout(ui.Area) {
}

func (p *Palette) Render() {
	if !p.enabled {
		return
	}

	p.resize()
	p.scroll()

	vt.Sync.Bsu()

	vt.Buf.Write(
		vt.HideCursor,
		p.colorBackground,
	)
	vt.ClearArea(vt.Buf, p.area.Y, p.area.X, p.area.W, p.area.H)

	if len(p.filteredOptions) == 0 {
		p.renderEmpty()
	} else {
		p.renderOptions()
	}

	p.editor.Render()

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (p *Palette) resize() {
	a := p.parent.Area()

	p.listSize = min(len(p.filteredOptions), maxListSize)

	p.area.W = min(60, a.W)

	p.area.H = 3 + max(p.listSize, 1)
	if p.area.H > a.H {
		p.area.H = a.H
		if p.listSize > 0 {
			p.listSize = p.area.H - 3
		}
	}

	p.area.Y = a.Y + ((a.H - p.area.H) / 2)
	p.area.X = a.X + ((a.W - p.area.W) / 2)

	p.editor.Layout(ui.Area{
		Y: p.area.Y + 1,
		X: p.area.X + 2,
		W: p.area.W - 4,
		H: 1,
	})
}

func (p *Palette) scroll() {
	delta := p.selectedIndex - p.scrollIndex

	if delta < 0 {
		p.scrollIndex = p.selectedIndex
	} else if delta >= p.listSize {
		p.scrollIndex = p.selectedIndex - p.listSize + 1
	}
}

func (p *Palette) renderEmpty() {
	vt.SetCursor(vt.Buf, p.area.Y+2, p.area.X+2)
	vt.Buf.Write(
		p.colorOption,
		[]byte("No matching commands"),
	)
}

func (p *Palette) renderOptions() {
	i := 0
	y := p.area.Y + 2

	for {
		if i == p.listSize {
			break
		}

		index := p.scrollIndex + i
		option := p.filteredOptions[index]
		if option == nil {
			break
		}

		if y == p.area.Y+p.area.H {
			break
		}

		span := p.area.W - 4

		if index == p.selectedIndex {
			vt.Buf.Write(p.colorSelectedOption)
		} else {
			vt.Buf.Write(p.colorOption)
		}

		vt.SetCursor(vt.Buf, y, p.area.X+2)
		vt.WriteText(vt.Buf, &span, option.description)
		vt.WriteText(vt.Buf, &span, fmt.Sprintf("%*s", span, option.shortcuts))

		i += 1
		y += 1
	}
}
