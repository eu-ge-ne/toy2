package theme

type Tokens interface {
	DangerBg() []byte

	MainBg() []byte
	MainFg() []byte

	Light2Bg() []byte
	Light2Fg() []byte

	Light1Bg() []byte
	Light1Fg() []byte

	Light0Bg() []byte
	Light0Fg() []byte

	Dark0Bg() []byte
	Dark0Fg() []byte

	Dark1Fg() []byte
	Dark2Fg() []byte
}
