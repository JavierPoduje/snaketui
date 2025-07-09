package game

import "slices"

const (
	DefaultSnakeX = 4
	DefaultSnakeY = 4
)

type Snake struct {
	Body []Coord // head is at index 0
}

func NewSnake() *Snake {
	snakeBody := []Coord{
		{DefaultSnakeX, DefaultSnakeY},
	}

	return &Snake{
		Body: snakeBody,
	}
}

func (snake *Snake) Contains(coordinate Coord) bool {
	return slices.Contains(snake.Body, coordinate)
}

func (snake Snake) Head() Coord {
	return snake.Body[0]
}

func (snake *Snake) IsHead(coordinate Coord) bool {
	return snake.Head().Equals(coordinate)
}
