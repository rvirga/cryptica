// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import (
	"sync"
	"testing"
)

func partest(t *testing.T, solver Solver, set []Game) {
	var group sync.WaitGroup
	solutions := make([]Solution, len(set))
	for i, game := range set {
		group.Add(1)
		go func(i int, game Game) {
			solutions[i] = solver.Solve(game)
			group.Done()
		}(i, game)
	}
	group.Wait()
	for i, solution := range solutions {
		if solution == nil {
			t.Errorf("%v unsolved\n", set[i].Name)
		} else if len(solution) != set[i].MinSteps {
			t.Errorf("%v solved in %d steps (expected: %d)\n", set[i].Name,
				len(solution), set[i].MinSteps)
		}
	}
}

func seqtest(t *testing.T, solver Solver, set []Game) {
	for _, game := range set {
		solution := solver.Solve(game)
		if solution == nil {
			t.Errorf("%v unsolved\n", game.Name)
		} else if len(solution) != game.MinSteps {
			t.Errorf("%v solved in %d steps (expected: %d)\n", game.Name,
				len(solution), game.MinSteps)
		}
	}
}

func TestAjawDepthFirst(t *testing.T) {
	partest(t, new(DepthFirstSolver), setAjaw)
}

func TestBalamDepthFirst(t *testing.T) {
	partest(t, new(DepthFirstSolver), setBalam)
}

func TestKukDepthFirst(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	partest(t, new(DepthFirstSolver), setKuk)
}

func TestAjawBreadthFirst(t *testing.T) {
	partest(t, new(BreadthFirstSolver), setAjaw)
}

func TestBalamBreadthFirst(t *testing.T) {
	partest(t, new(BreadthFirstSolver), setBalam)
}

func TestKukBreadthFirst(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// test sequentially to save memory
	seqtest(t, new(BreadthFirstSolver), setKuk)
}
