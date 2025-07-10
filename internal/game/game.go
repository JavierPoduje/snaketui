package game

import "math/rand/v2"

const (
	DefaultAppleX = 2
	DefaultAppleY = 2
)

type GameState int

const (
	Running = iota
	GameOver
	Paused
)

type Game struct {
	Apple    *Coord
	Canvas   *Canvas
	NextMove Direction
	Snake    *Snake
	State    GameState
	Stats    *Stats
}

func NewGame(width, height int) *Game {
	return &Game{
		Apple:    &Coord{X: DefaultAppleX, Y: DefaultAppleY},
		Canvas:   NewCanvas(width, height),
		NextMove: Up,
		Snake:    NewSnake(),
		State:    Running,
		Stats:    NewStats(),
	}
}

func (game *Game) Tick(direction Direction) {
	if !game.nextMoveIsValid(direction) {
		game.State = GameOver
		return
	}

	if !game.nextMoveIsValid(direction) {
		return
	}

	game.NextMove = direction
	err := game.Snake.Move(game.NextMove)
	if err != nil {
		panic(err)
	}

	if game.Snake.Body[0] == *game.Apple {
		game.eatApple()
	}
}

func (game Game) nextMoveIsValid(dir Direction) bool {
	nextHead := game.Snake.NextHead(dir)
	return game.Snake.IsValidMove(dir) && game.Canvas.InBounds(nextHead)
}

func (game *Game) eatApple() {
	game.Snake.Add()
	game.Stats.EatApple()
	game.Stats.UpdateScore(game.Snake.Speed)

	game.Apple = game.getRandApple()
	for game.Snake.Contains(*game.Apple) {
		game.Apple = game.getRandApple()
	}
}

func (game Game) getRandApple() *Coord {
	width := game.Canvas.Width
	height := game.Canvas.Height
	return &Coord{
		X: rand.IntN(height),
		Y: rand.IntN(width),
	}
}
