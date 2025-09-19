package ui

type Control interface {
	Area() Area
	Layout(Area)
	Render()
}
