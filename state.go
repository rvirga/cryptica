// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import (
	"bytes"
	"fmt"
)

type Wall interface {
	Contains(pos Position) bool
}

type HorizontalWall struct {
	Y, Start, End int
}

func (wall HorizontalWall) Contains(pos Position) bool {
	return pos.Y == wall.Y && pos.X >= wall.Start && pos.X <= wall.End
}

type VerticalWall struct {
	X, Start, End int
}

func (wall VerticalWall) Contains(pos Position) bool {
	return pos.X == wall.X && pos.Y >= wall.Start && pos.Y <= wall.End
}

type Board struct {
	W, H  int
	Walls []Wall
}

func (board Board) Encode(state State) (n uint64) {
	n = 0
	power, s := uint64(1), uint64(board.W)*uint64(board.H)+1
	for _, tile := range state.Tiles {
		n += (uint64(tile.Y)*uint64(board.W) + uint64(tile.X) + 1) * power
		power *= s
	}
	return n
}

func (board Board) Decode(n uint64) (state State) {
	s := uint64(board.W)*uint64(board.H) + 1
	state = State{board, make([]Position, 0)}
	for n > 0 {
		m := n%s - 1
		x, y := int(m%uint64(board.W)), int(m/uint64(board.W))
		state.Tiles = append(state.Tiles, Position{x, y})
		n /= s
	}
	return state
}

type State struct {
	Board Board
	Tiles []Position
}

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

func (state State) String() string {
	var buffer bytes.Buffer
	for y := 0; y < state.Board.H; y++ {
		for x := 0; x < state.Board.W; x++ {
			item := state.ItemAt(Position{x, y})
			if tile, ok := item.(int); ok {
				buffer.WriteString(fmt.Sprintf("%X", tile))
			} else if empty := item.(bool); empty {
				buffer.WriteString(".")
			} else {
				buffer.WriteString("#")
			}

		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (state State) move(pos Position, dir Direction) (Position, bool) {
	newpos := pos.Move(dir)
	if newpos.X < 0 || newpos.X > state.Board.W ||
		newpos.Y < 0 || newpos.Y > state.Board.H {
		return pos, false
	}
	item := state.ItemAt(newpos)
	if tile, ok := item.(int); ok {
		_, ok := state.move(state.Tiles[tile], dir)
		return newpos, ok
	} else if empty := item.(bool); empty {
		return newpos, true
	} else {
		return pos, false
	}
}

func (state State) Move(dir Direction) State {
	tiles := make([]Position, len(state.Tiles))
	copy(tiles, state.Tiles)
	for tile, pos := range state.Tiles {
		if newpos, ok := state.move(pos, dir); ok {
			tiles[tile] = newpos
		}
	}
	return State{state.Board, tiles}
}

type Goal []Position

func (state State) Match(goal Goal) bool {
	for i, pos := range goal {
		if !pos.Equal(state.Tiles[i]) {
			return false
		}
	}
	return true
}
