package debug

import (
	"fmt"
	"runtime"
	"time"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

const mib = 1024 * 1024

type Debug struct {
	area    ui.Area
	enabled bool

	inputTime  time.Duration
	renderTime time.Duration

	colorBackground []byte
	colorText       []byte
}

func New() *Debug {
	return &Debug{}
}

func (d *Debug) Enable(enable bool) {
	d.enabled = enable
}

func (d *Debug) ToggleEnabled() {
	d.enabled = !d.enabled
}

func (d *Debug) SetColors(t theme.Theme) {
	d.colorBackground = t.Light0Bg()
	d.colorText = append(t.Light0Bg(), t.Dark0Fg()...)
}

func (d *Debug) SetArea(a ui.Area) {
	w := std.Clamp(30, 0, a.W)
	h := std.Clamp(7, 0, a.H)

	d.area = ui.Area{
		Y: a.Y + a.H - h,
		X: a.X + a.W - w,
		W: w,
		H: h,
	}
}

func (d *Debug) Render() {
	if !d.enabled {
		return
	}

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(d.colorBackground)
	vt.ClearArea(vt.Buf, d.area)
	vt.Buf.Write(d.colorText)
	vt.SetCursor(vt.Buf, d.area.Y+1, d.area.X+1)
	fmt.Fprintf(vt.Buf, "Input  : %v", d.inputTime)
	vt.SetCursor(vt.Buf, d.area.Y+2, d.area.X+1)
	fmt.Fprintf(vt.Buf, "Render : %v", d.renderTime)
	vt.SetCursor(vt.Buf, d.area.Y+3, d.area.X+1)
	fmt.Fprintf(vt.Buf, "Alloc  : %v MiB", mem.Alloc/mib)

	vt.Buf.Write(vt.RestoreCursor)
	vt.Buf.Write(vt.ShowCursor)

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (d *Debug) SetInputTime(elapsed time.Duration) {
	d.inputTime = elapsed

	if d.enabled {
		d.Render()
	}
}

func (d *Debug) SetRenderTime(elapsed time.Duration) {
	d.renderTime = elapsed

	if d.enabled {
		d.Render()
	}
}
