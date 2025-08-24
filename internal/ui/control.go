package ui

type Control interface {
	IsEnabled() bool
	ToggleEnabled()
	SetEnabled(bool)
	Area() Area
	Layout(Area)
	Render()
}
