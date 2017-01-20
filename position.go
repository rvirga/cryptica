// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import "fmt"

// Position represents the cartesian coordinates of a location on the board.
// The origin (0,0) correspond to the top-left corner.
type Position struct {
	X, Y int
}

// Test for equality of two positions.
func (pos1 Position) Equal(pos2 Position) bool {
	return pos1.X == pos2.X && pos1.Y == pos2.Y
}

// Printable representation of a position as a pair of numbers (x,y).
func (pos Position) String() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}

// Direction represents a single step in one of the possible four directions.
type Direction int

// A direction is encoded as number between 0 and 4.
const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

// Printable representation of a direction step. This is one of the
// following: ↑, ↓, ←, or →.
func (dir Direction) String() string {
	switch dir {
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

// This method constructs a new position by moving the given one in the
// specified direction.
// No checks are performed here, so the new position might be invalid (i.e. it
// might fall outside the board).
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
