// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

import "fmt"

func ExampleSolution() {
	var solution Solution = []Direction{
		Up, Up, Down, Left, Right, Up,
	}
	fmt.Printf("%v\n", solution)
	// Output: ↑ ↑ ↓ ← → ↑
}

func ExampleGame() {
	var game = Game{
		Name: "AJAW 2",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{4, 1, 4},
					VerticalWall{5, 1, 4},
				},
			},
			Tiles: []Position{
				{1, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{5, 0}, {6, 4},
		},
		MinSteps: 10,
	}
	fmt.Print(game.Start)
	// Output:
	// .......
	// ....##.
	// .0.1##.
	// ....##.
	// ....##.
}

func ExampleSolver() {
	var game Game = setChan[27]
	var solver Solver = new(DepthFirstSolver)
	if solution := solver.Solve(game); solution != nil {
		fmt.Printf("Solution of %v: %v\n", game.Name, solution)
	}
	// Output: Solution of CHAN 28: → ← ↓ ← ↑ → ← ↓ → →
}
