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
	State State
	Goal  Goal
	Steps int
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
		if n > game.Steps {
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
			if result := f(state.Move(dir), append(solution, dir), k); result != nil {
				return result
			}
		}
		return nil
	}
	return f(game.State, make(Solution, 0), f)
}

type BreadthFirstSolver struct {
}

func (solver BreadthFirstSolver) Solve(game Game) Solution {
	cache := make(map[uint64]Solution)
	var current, next []uint64
	{
		code := game.State.Board.Encode(game.State)
		cache[code] = make(Solution, 0)
		current = []uint64{code}
	}
	depth := 0
	for len(current) > 0 {
		next = make([]uint64, 0)
		for _, code := range current {
			solution := cache[code]
			state := game.State.Board.Decode(code)
			if state.Match(game.Goal) {
				return solution
			}
			if depth >= game.Steps {
				continue
			}
			for i := Up; i <= Right; i++ {
				dir := Direction(i)
				newstate := state.Move(dir)
				newcode := game.State.Board.Encode(newstate)
				if _, exists := cache[newcode]; !exists {
					newsolution := make(Solution, len(solution))
					copy(newsolution, solution)
					newsolution = append(newsolution, dir)
					cache[newcode] = newsolution
					next = append(next, newcode)
				}
			}
		}
		depth++
		current = next
	}
	return nil
}
