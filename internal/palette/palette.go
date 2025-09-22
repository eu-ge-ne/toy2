package palette

import (
	"slices"
	"strings"

	"github.com/eu-ge-ne/toy2/internal/editor"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

const maxListSize = 10

type Palette struct {
	area            ui.Area
	enabled         bool
	parent          ui.Control
	options         []*Option
	filteredOptions []*Option
	editor          *editor.Editor
	listSize        int
	selectedIndex   int
	scrollIndex     int

	colorBackground     []byte
	colorOption         []byte
	colorSelectedOption []byte
}

func New(parent ui.Control, options []*Option) *Palette {
	return &Palette{
		parent:  parent,
		options: options,
		editor:  editor.New(false),
	}
}

func (d *Palette) SetColors(t theme.Tokens) {
	d.colorBackground = t.Light1Bg()
	d.colorOption = append(t.Light1Bg(), t.Light1Fg()...)
	d.colorSelectedOption = append(t.Light2Bg(), t.Light1Fg()...)

	d.editor.SetColors(t)
}

func (*Palette) Layout(ui.Area) {}

func (p *Palette) Open(done chan<- *Option) {
	p.enabled = true
	p.editor.Enabled = true

	p.editor.Buffer.Reset("")
	p.editor.Reset(false)

	p.filter()
	p.parent.Render()

	result := p.processInput()

	p.enabled = false
	p.editor.Enabled = false

	done <- result
}

func (p *Palette) filter() {
	p.selectedIndex = 0

	text := strings.ToUpper(
		strings.Join(
			slices.Collect(p.editor.Buffer.ReadToEnd(0)),
			"",
		),
	)

	if len(text) == 0 {
		p.filteredOptions = p.options
		return
	}

	p.filteredOptions = []*Option{}
	for _, o := range p.options {
		d := strings.ToUpper(o.Description)
		i := strings.Index(d, text)
		if i >= 0 {
			p.filteredOptions = append(p.filteredOptions, o)
		}
	}
}

func (p *Palette) processInput() *Option {
	for {
		for key := range vt.Read() {
			switch key.Name {
			case "ESC":
				return nil
			case "ENTER":
				return p.filteredOptions[p.selectedIndex]
			case "UP":
				if len(p.filteredOptions) > 0 {
					p.selectedIndex = max(p.selectedIndex-1, 0)
					p.parent.Render()
				}
				continue
			case "DOWN":
				if len(p.filteredOptions) > 0 {
					p.selectedIndex = min(p.selectedIndex+1, len(p.filteredOptions)-1)
					p.parent.Render()
				}
				continue
			}

			if p.editor.HandleKey(key) {
				p.filter()
				p.parent.Render()
			}
		}
	}
}
