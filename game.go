// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import (
	"bytes"
	"math"
)

// Solution represents a tentative partial solution to the puzzle, and consists
// of a sequence of steps.
type Solution []Direction

// A solution is printed as a space-separated sequence of direction steps.
func (solution Solution) String() string {
	var buffer bytes.Buffer
	for i, dir := range solution {
		if i > 0 {
			buffer.WriteRune(' ')
		}
		buffer.WriteString(dir.String())
	}
	return buffer.String()
}

// Game represents single puzzle game. This comprises a name (e.g. "K'UK 32"),
// for ease of identification, an initial configuration, a goal configuration,
// and the minimum number of steps required to solve it.
//
// NOTE: in case this last number is unknown, we recommend you set it to an
// estimated upper bound, and then try to solve using breadth-first search.
// Breadth-first will always return the shortest solution (if there's one),
// allowing you to determine the correct value for this field.
type Game struct {
	Name     string
	Start    State
	Goal     Goal
	MinSteps int
}

// This method determines if the size of the board and the number of tiles
// are such that the requirements for this package to work (that is, that
// any game configuration can be uniquely encoded by a 64-bit number) are
// met.
func (game Game) Encodeable() bool {
	size := float64(game.Start.Board.W * game.Start.Board.H)
	maxTiles := int(math.Floor(64 / math.Log2(size+1)))
	return len(game.Start.Tiles) <= maxTiles
}

// The Solver interface is a generalization of the solving process. Both
// depth-first and breadth-first solvers implement this. Given a game,
// the Solve method should return the fist solution of length <= game.MinSteps,
// or nil if it cannot find one.
type Solver interface {
	Solve(game Game) Solution
}

// The DepthFirstSolver searches for solutions in a depth-first manner.
// It constructs a tentative solution of length game.MinSteps, backtracking
// if necessary.
// Although this uses a lot less memory than breadth-first search, it speeds
// up search by memorizing (as 64-bit numbers) the states it has already
// examined, and therefore it will consume memory as it progresses.
type DepthFirstSolver struct {
}

func (solver DepthFirstSolver) Solve(game Game) Solution {
	// sanity check
	if !game.Encodeable() {
		return nil
	}
	cache := make(map[uint64]int)
	f := func(state State, solution Solution, k interface{}) Solution {
		n := len(solution)
		if n > game.MinSteps {
			return nil
		}
		code := state.Board.Encode(state)
		if m, exists := cache[code]; exists && m <= n {
			return nil
		}
		cache[code] = n
		if state.Match(game.Goal) {
			return solution
		}
		f := k.(func(State, Solution, interface{}) Solution)
		for dir := Up; dir <= Right; dir++ {
			if newstate, changed := state.Move(dir); changed {
				if result := f(newstate, append(solution, dir), k); result != nil {
					return result
				}
			}
		}
		return nil
	}
	return f(game.Start, make(Solution, 0), f)
}

// BreadthFirstSolver will search for solutions in a breadth-first manner.
// First, all solutions of length 1 will be tested, then all those of length
// 2, and so on. If no solution of length <= game.MinSteps is found, the
// search process will stop.
// Although this method is usually slightly faster than depth-first search,
// and moreover is always guaranteed to produce the shortest solution, it
// might require massive amounts of memory to solve the most difficult games.
type BreadthFirstSolver struct {
}

func (solver BreadthFirstSolver) Solve(game Game) Solution {
	// sanity check
	if !game.Encodeable() {
		return nil
	}
	cache := make(map[uint64]Solution)
	var current, next []uint64
	{
		code := game.Start.Board.Encode(game.Start)
		cache[code] = make(Solution, 0)
		current = []uint64{code}
	}
	depth := 0
	for len(current) > 0 {
		next = make([]uint64, 0)
		for _, code := range current {
			solution := cache[code]
			state := game.Start.Board.Decode(code)
			if state.Match(game.Goal) {
				return solution
			}
			if depth >= game.MinSteps {
				continue
			}
			for dir := Up; dir <= Right; dir++ {
				if newstate, changed := state.Move(dir); changed {
					newcode := game.Start.Board.Encode(newstate)
					if _, exists := cache[newcode]; !exists {
						newsolution := make(Solution, len(solution))
						copy(newsolution, solution)
						newsolution = append(newsolution, dir)
						cache[newcode] = newsolution
						next = append(next, newcode)
					}
				}
			}
			cache[code] = nil
		}
		depth++
		current = next
	}
	return nil
}
