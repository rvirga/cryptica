// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

type Direction int

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

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
