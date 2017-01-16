// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

var setBalam = []Game{
	Game{
		Name: "B'ALAM 1",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 1},
					VerticalWall{0, 0, 1},
					HorizontalWall{4, 0, 1},
				},
			},
			Tiles: []Position{
				{4, 1}, {4, 3},
			},
		},
		Goal: []Position{
			{3, 3}, {3, 1},
		},
		MinSteps: 17,
	},
	Game{
		Name: "B'ALAM 2",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{3, 3, 3},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 0}, {2, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 3}, {6, 1},
		},
		MinSteps: 16,
	},
	Game{
		Name: "B'ALAM 3",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 5, 5},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 0, 3},
					HorizontalWall{3, 4, 4},
					HorizontalWall{4, 5, 5},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1}, {6, 0}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {0, 0}, {6, 3},
		},
		MinSteps: 9,
	},
	Game{
		Name: "B'ALAM 4",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{2, 0, 0},
					VerticalWall{5, 2, 4},
					VerticalWall{6, 2, 4},
				},
			},
			Tiles: []Position{
				{4, 2}, {1, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{5, 0}, {0, 4},
		},
		MinSteps: 12,
	},
	Game{
		Name: "B'ALAM 5",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 6},
					VerticalWall{6, 1, 2},
					HorizontalWall{4, 0, 6},
					VerticalWall{0, 2, 3},
				},
			},
			Tiles: []Position{
				{1, 2}, {5, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{3, 1}, {3, 3},
		},
		MinSteps: 19,
	},
	Game{
		Name: "B'ALAM 6",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 2},
					HorizontalWall{0, 4, 5},
					VerticalWall{6, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 2}, {6, 4},
		},
		MinSteps: 24,
	},
	Game{
		Name: "B'ALAM 7",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{3, 2}, {2, 1}, {3, 1}, {4, 0}, {5, 0}, {6, 0},
			},
		},
		Goal: []Position{
			{2, 2}, {1, 1}, {1, 2},
		},
		MinSteps: 20,
	},
	Game{
		Name: "B'ALAM 8",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 0}, {2, 2}, {5, 0}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 2}, {1, 0},
		},
		MinSteps: 23,
	},
	Game{
		Name: "B'ALAM 9",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 1},
					HorizontalWall{0, 6, 6},
					HorizontalWall{1, 0, 1},
					HorizontalWall{3, 0, 0},
				},
			},
			Tiles: []Position{
				{2, 4}, {3, 4},
			},
		},
		Goal: []Position{
			{3, 1}, {4, 4},
		},
		MinSteps: 27,
	},
	Game{
		Name: "B'ALAM 10",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 0},
				},
			},
			Tiles: []Position{
				{3, 2}, {3, 3}, {6, 0}, {0, 4},
			},
		},
		Goal: []Position{
			{3, 1}, {3, 4},
		},
		MinSteps: 20,
	},
	Game{
		Name: "B'ALAM 11",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{1, 1, 4},
					VerticalWall{3, 0, 3},
					VerticalWall{5, 1, 4},
				},
			},
			Tiles: []Position{
				{1, 0}, {3, 4}, {5, 0},
			},
		},
		Goal: []Position{
			{0, 0}, {2, 3}, {6, 2},
		},
		MinSteps: 15,
	},
	Game{
		Name: "B'ALAM 12",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{0, 0}, {2, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
			},
		},
		Goal: []Position{
			{5, 4}, {6, 0},
		},
		MinSteps: 14,
	},
	Game{
		Name: "B'ALAM 13",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 5, 5},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 0}, {3, 2},
			},
		},
		Goal: []Position{
			{0, 1}, {6, 2}, {2, 3},
		},
		MinSteps: 16,
	},
	Game{
		Name: "B'ALAM 14",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 2, 4},
					VerticalWall{3, 1, 3},
				},
			},
			Tiles: []Position{
				{5, 1}, {1, 3}, {5, 3}, {1, 1},
			},
		},
		Goal: []Position{
			{4, 1}, {2, 3}, {4, 3},
		},
		MinSteps: 12,
	},
	Game{
		Name: "B'ALAM 15",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{1, 0, 0},
					VerticalWall{1, 2, 2},
					VerticalWall{5, 1, 2},
					HorizontalWall{3, 1, 2},
					HorizontalWall{3, 4, 5},
				},
			},
			Tiles: []Position{
				{2, 2}, {3, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{3, 3}, {1, 1}, {5, 0},
		},
		MinSteps: 13,
	},
	Game{
		Name: "B'ALAM 16",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 0, 4},
					HorizontalWall{2, 0, 4},
					HorizontalWall{4, 3, 6},
				},
			},
			Tiles: []Position{
				{3, 3}, {4, 0}, {3, 0}, {2, 0},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		MinSteps: 17,
	},
	Game{
		Name: "B'ALAM 17",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 3},
					VerticalWall{3, 1, 4},
				},
			},
			Tiles: []Position{
				{5, 2}, {1, 2}, {2, 1}, {0, 3},
			},
		},
		Goal: []Position{
			{4, 2}, {2, 2},
		},
		MinSteps: 13,
	},
	Game{
		Name: "B'ALAM 18",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{3, 1, 4},
					VerticalWall{4, 2, 4},
					HorizontalWall{1, 6, 6},
				},
			},
			Tiles: []Position{
				{2, 0}, {3, 0}, {4, 0},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 2},
		},
		MinSteps: 19,
	},
	Game{
		Name: "B'ALAM 19",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{3, 0, 1},
				},
			},
			Tiles: []Position{
				{0, 3}, {0, 2}, {0, 1},
			},
		},
		Goal: []Position{
			{6, 1}, {6, 2}, {6, 3},
		},
		MinSteps: 17,
	},
	Game{
		Name: "B'ALAM 20",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					VerticalWall{6, 0, 4},
					HorizontalWall{1, 2, 4},
					HorizontalWall{3, 2, 4},
					VerticalWall{3, 1, 3},
				},
			},
			Tiles: []Position{
				{4, 0}, {2, 0}, {3, 0},
			},
		},
		Goal: []Position{
			{4, 2}, {2, 2},
		},
		MinSteps: 11,
	},
	Game{
		Name: "B'ALAM 21",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 4, 4},
					HorizontalWall{3, 4, 4},
				},
			},
			Tiles: []Position{
				{3, 2}, {4, 2}, {2, 2},
			},
		},
		Goal: []Position{
			{5, 1}, {2, 4}, {1, 0},
		},
		MinSteps: 13,
	},
	Game{
		Name: "B'ALAM 22",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 1, 5},
					HorizontalWall{3, 1, 2},
					VerticalWall{5, 3, 4},
				},
			},
			Tiles: []Position{
				{2, 0}, {1, 0}, {1, 1}, {3, 0}, {4, 0},
			},
		},
		Goal: []Position{
			{3, 3}, {4, 3}, {6, 4},
		},
		MinSteps: 28,
	},
	Game{
		Name: "B'ALAM 23",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 5},
					VerticalWall{1, 2, 3},
					HorizontalWall{3, 3, 5},
				},
			},
			Tiles: []Position{
				{4, 0}, {3, 0}, {2, 0},
			},
		},
		Goal: []Position{
			{2, 3}, {5, 2}, {0, 0},
		},
		MinSteps: 11,
	},
	Game{
		Name: "B'ALAM 24",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 2, 4},
					VerticalWall{1, 1, 4},
					VerticalWall{5, 0, 3},
					VerticalWall{6, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 1}, {3, 3}, {3, 2},
			},
		},
		Goal: []Position{
			{6, 4}, {0, 0},
		},
		MinSteps: 29,
	},
	Game{
		Name: "B'ALAM 25",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 2, 2},
					HorizontalWall{2, 3, 4},
					HorizontalWall{3, 2, 2},
					VerticalWall{4, 1, 3},
				},
			},
			Tiles: []Position{
				{0, 3}, {0, 2}, {0, 1}, {6, 2},
			},
		},
		Goal: []Position{
			{3, 3}, {2, 2}, {3, 1},
		},
		MinSteps: 18,
	},
	Game{
		Name: "B'ALAM 26",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{5, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 2}, {2, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		MinSteps: 16,
	},
	Game{
		Name: "B'ALAM 27",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 4},
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{2, 3}, {4, 3}, {2, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		MinSteps: 13,
	},
	Game{
		Name: "B'ALAM 28",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 0, 2},
					HorizontalWall{1, 4, 6},
					HorizontalWall{2, 2, 4},
					HorizontalWall{3, 1, 2},
					HorizontalWall{3, 4, 5},
					VerticalWall{2, 0, 3},
					VerticalWall{4, 0, 3},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 0}, {6, 2}, {0, 2}, {0, 4}, {6, 4},
			},
		},
		Goal: []Position{
			{3, 3}, {3, 1}, {5, 2}, {1, 2},
		},
		MinSteps: 9,
	},
	Game{
		Name: "B'ALAM 29",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 5},
					VerticalWall{1, 1, 4},
					VerticalWall{5, 1, 4},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{3, 2}, {0, 4}, {6, 0}, {2, 4}, {4, 4},
			},
		},
		Goal: []Position{
			{3, 1}, {0, 0}, {6, 4},
		},
		MinSteps: 15,
	},
	Game{
		Name: "B'ALAM 30",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 1, 2},
					VerticalWall{2, 2, 3},
					HorizontalWall{2, 4, 5},
					VerticalWall{4, 1, 2},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1}, {3, 2},
			},
		},
		Goal: []Position{
			{1, 3}, {5, 1},
		},
		MinSteps: 13,
	},
}
