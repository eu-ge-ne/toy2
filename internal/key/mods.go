package key

type Mods byte

const (
	Shift    Mods = 1
	Alt      Mods = 2
	Ctrl     Mods = 4
	Super    Mods = 8
	Hyper    Mods = 16
	Meta     Mods = 32
	CapsLock Mods = 64
	NumLock  Mods = 128
)
