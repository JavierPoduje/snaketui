package game

type Game struct {
	Snake    *Snake
	Canvas   *Canvas
	NextMove Direction
}

func NewGame(width, height int) *Game {
	return &Game{
		Canvas:   NewCanvas(width, height),
		Snake:    NewSnake(),
		NextMove: Up,
	}
}

func (game *Game) Tick(direction Direction) {
	if !game.nextMoveIsValid(direction) {
		return
	}

	game.NextMove = direction
	err := game.Snake.Move(game.NextMove)
	if err != nil {
		panic(err)
	}
}

func (game Game) nextMoveIsValid(dir Direction) bool {
	nextHead := game.Snake.NextHead(dir)
	return game.Snake.IsValidMove(dir) && game.Canvas.InBounds(nextHead)
}
