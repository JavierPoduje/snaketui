package game

type Canvas struct {
	Width  int
	Height int
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
	}
}
