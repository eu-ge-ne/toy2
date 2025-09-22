package ui

type Control interface {
	Area() Area
	Render()
}
