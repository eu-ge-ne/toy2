package command

type App interface {
	Copy()
	Cut()
	Debug()
	Exit()
	Palette()
	Paste()
	Redo()
	Save()
	SelectAll()
	ThemeBase16()
	ThemeGray()
	ThemeNeutral()
	ThemeSlate()
	ThemeStone()
	ThemeZinc()
	Undo()
	Whitespace()
	Wrap()
	Zen()
}
