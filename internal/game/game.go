package game

type Game struct {
	Snake  *Snake
	Canvas *Canvas
}

func NewGame(width, height int) *Game {
	return &Game{
		Canvas: NewCanvas(width, height),
		Snake:  NewSnake(),
	}
}
