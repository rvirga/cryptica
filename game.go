// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import "bytes"

type Solution []Direction

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

type Game struct {
	Name     string
	Start    State
	Goal     Goal
	MinSteps int
}

type Solver interface {
	Solve(game Game) Solution
}

type DepthFirstSolver struct {
}

func (solver DepthFirstSolver) Solve(game Game) Solution {
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
		for i := Up; i <= Right; i++ {
			dir := Direction(i)
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

type BreadthFirstSolver struct {
}

func (solver BreadthFirstSolver) Solve(game Game) Solution {
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
			for i := Up; i <= Right; i++ {
				dir := Direction(i)
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
