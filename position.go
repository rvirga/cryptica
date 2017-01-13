package cryptica

import "fmt"

type Position struct {
	X, Y int
}

func (pos1 Position) Equal(pos2 Position) bool {
	return pos1.X == pos2.X && pos1.Y == pos2.Y
}

func (pos Position) String() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}

const (
	Up    = 0
	Down  = 1
	Left  = 2
	Right = 3
)

type Direction int

func (dir *Direction) String() string {
	switch *dir {
	case Up:
		return "↑"
	case Down:
		return "↓"
	case Left:
		return "←"
	default:
		return "→"
	}
}

func (pos Position) Move(dir Direction) Position {
	switch dir {
	case Up:
		return Position{pos.X, pos.Y - 1}
	case Down:
		return Position{pos.X, pos.Y + 1}
	case Left:
		return Position{pos.X - 1, pos.Y}
	default:
		return Position{pos.X + 1, pos.Y}
	}
}
