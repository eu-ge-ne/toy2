package data

type Editor interface {
	Backspace() bool
	Bottom(bool) bool
	Copy() bool
	Cut() bool
	Delete() bool
	Down(bool) bool
	End(bool) bool
	Enter() bool
	Home(bool) bool
	Insert(string) bool
	Left(bool) bool
	PageDown(bool) bool
	PageUp(bool) bool
	Paste() bool
	Redo() bool
	Right(bool) bool
	SelectAll() bool
	Top(bool) bool
	Undo() bool
	Up(bool) bool
}
