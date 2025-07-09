package game

type Direction int

const (
	Up = iota
	Right
	Down
	Left
)

func (d Direction) IsOpposite(other Direction) bool {
	switch d {
	case Up:
		return other == Down
	case Down:
		return other == Up
	case Left:
		return other == Right
	case Right:
		return other == Left
	default:
		panic("unknown direction")
	}
}
