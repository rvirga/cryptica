// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import (
	"bytes"
	"fmt"
)

// A wall is an unmovable barrier. Its shape is defined by the Contains
// function.
type Wall interface {
	Contains(pos Position) bool
}

// HorizontalWall is an instance of a wall. It occupies a single row Y,
// from column Start to column End (both included).
type HorizontalWall struct {
	Y, Start, End int
}

func (wall HorizontalWall) Contains(pos Position) bool {
	return pos.Y == wall.Y && pos.X >= wall.Start && pos.X <= wall.End
}

// VerticalWall is an instance of a wall. It occupies a single column X,
// from row Start to row End (both included).
type VerticalWall struct {
	X, Start, End int
}

func (wall VerticalWall) Contains(pos Position) bool {
	return pos.X == wall.X && pos.Y >= wall.Start && pos.Y <= wall.End
}

// Board represent the barebone game board, without any tiles. As such, it
// needs only to specify the size of the board (width W and height H), and
// its walls (if any).
type Board struct {
	W, H  int
	Walls []Wall
}

// A game configuration is encoded into a single 64-bits unsigned integer.
// Specifically, in a WxH board, the position of each tile is encoded by a
// number 1 <= n <= W*H. The position of the last tile occupies the most
// significant bits of this encoding, the first the least significant ones.
//
// IMPORTANT: use of this encoding is critical in the implementation of this
// package, and therefore limits the size/complexity of the games that it
// can tackle. More precisely, given a WxH board, we cannot solve games
// involving more than 64 * log(2) / log(W*H+1) tiles. This works out OK
// for all Cryptica puzzles, but if you use this package on puzzles of your
// own creation, you should keep in mind this limitation.
func (board Board) Encode(state State) (n uint64) {
	n = 0
	power, s := uint64(1), uint64(board.W)*uint64(board.H)+1
	for _, tile := range state.Tiles {
		n += (uint64(tile.Y)*uint64(board.W) + uint64(tile.X) + 1) * power
		power *= s
	}
	return
}

// Decodes the given 64-bit unsigned integer into a newly created game
// configuration. In creating the new state, we copy the board by reference
// to save time and memory.
func (board *Board) Decode(n uint64) (state State) {
	s := uint64(board.W)*uint64(board.H) + 1
	state = State{board, make([]Position, 0)}
	for n > 0 {
		m := n%s - 1
		x, y := int(m%uint64(board.W)), int(m/uint64(board.W))
		state.Tiles = append(state.Tiles, Position{x, y})
		n /= s
	}
	return
}

// State represent a specific board configuration. It needs therefore to
// specify the underlying board, as well as the position of all its tiles.
// Since a step in the solution doesn't modify the board (its size, or the
// position of its walls), this is stored as a pointer and shared by all
// states of a puzzle.
type State struct {
	Board *Board
	Tiles []Position
}

// This method returns the item occupying the given position. Specifically:
//  - if there is a tile occupying the position, it will return its index (int)
//  - if it is invalid or occupied by a wall, it will return false (bool)
//  - if it is valid, and unoccupied, it will return true (bool)
func (state State) ItemAt(pos Position) interface{} {
	if pos.X < 0 || pos.X >= state.Board.W {
		return false
	}
	if pos.Y < 0 || pos.Y >= state.Board.H {
		return false
	}
	for _, wall := range state.Board.Walls {
		if wall.Contains(pos) {
			return false
		}
	}
	for i, tile := range state.Tiles {
		if tile.Equal(pos) {
			return i
		}
	}
	return true
}

// Printable representation of a board configuration as a WxH matrix of
// characters. Tiles are represented by their index as a hexadecimal
// number. Walls are represented by '#', and empty slots by '.'.
func (state State) String() string {
	var buffer bytes.Buffer
	for y := 0; y < state.Board.H; y++ {
		for x := 0; x < state.Board.W; x++ {
			item := state.ItemAt(Position{x, y})
			if tile, ok := item.(int); ok {
				buffer.WriteString(fmt.Sprintf("%X", tile))
			} else if empty := item.(bool); empty {
				buffer.WriteRune('.')
			} else {
				buffer.WriteRune('#')
			}

		}
		buffer.WriteRune('\n')
	}
	return buffer.String()
}

// Move the given position in the specified direction. All due checks are
// performed, and if the move is not allowed, false is returned together
// with the old position. Otherwise, the new position and true is returned.
func (state State) move(pos Position, dir Direction) (Position, bool) {
	newpos := pos.Move(dir)
	if newpos.X < 0 || newpos.X > state.Board.W ||
		newpos.Y < 0 || newpos.Y > state.Board.H {
		return pos, false
	}
	item := state.ItemAt(newpos)
	if tile, ok := item.(int); ok {
		if _, changed := state.move(state.Tiles[tile], dir); changed {
			return newpos, true
		} else {
			return pos, false
		}
	} else if empty := item.(bool); empty {
		return newpos, true
	} else {
		return pos, false
	}
}

// Constructs a new state by moving the given state in the specified direction.
// If no tiles changed position as a result of this move, and  the new state is
// the same as the old one, we return false. Otherwise, we return true.
func (state State) Move(dir Direction) (State, bool) {
	changed := false
	tiles := make([]Position, len(state.Tiles))
	for tile, pos := range state.Tiles {
		newpos, poschanged := state.move(pos, dir)
		changed = changed || poschanged
		tiles[tile] = newpos
	}
	return State{state.Board, tiles}, changed
}

// Goal specifies the goal positions of the first k of the n tiles (k <= n).
// For tiles k, k+1, n-1, any position will do.
type Goal []Position

// Checks if a state matches the goal, by verifying that all the tiles we care
// about are in the desired positions.
func (state State) Match(goal Goal) bool {
	for i, pos := range goal {
		if !pos.Equal(state.Tiles[i]) {
			return false
		}
	}
	return true
}
