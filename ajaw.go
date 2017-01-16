// Copyright 2017 Roberto Virga. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptica

var setAjaw = []Game{
	Game{
		Name: "AJAW 1",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{4, 4}, {2, 4},
			},
		},
		Goal: []Position{
			{6, 0}, {5, 0},
		},
		MinSteps: 7,
	},
	Game{
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
	},
	Game{
		Name: "AJAW 3",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{2, 0, 1},
					VerticalWall{3, 0, 1},
				},
			},
			Tiles: []Position{
				{5, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{0, 2},
		},
		MinSteps: 6,
	},
	Game{
		Name: "AJAW 4",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{3, 0, 4},
					HorizontalWall{3, 3, 4},
				},
			},
			Tiles: []Position{
				{0, 2}, {6, 2},
			},
		},
		Goal: []Position{
			{0, 0}, {4, 4},
		},
		MinSteps: 8,
	},
	Game{
		Name: "AJAW 5",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 0, 6},
					VerticalWall{4, 1, 2},
				},
			},
			Tiles: []Position{
				{6, 0}, {6, 4}, {0, 0},
			},
		},
		Goal: []Position{
			{6, 1}, {0, 4},
		},
		MinSteps: 12,
	},
	Game{
		Name: "AJAW 6",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 4},
					HorizontalWall{1, 2, 4},
					HorizontalWall{3, 2, 4},
					HorizontalWall{4, 2, 4},
				},
			},
			Tiles: []Position{
				{5, 1}, {1, 3},
			},
		},
		Goal: []Position{
			{1, 1}, {5, 3},
		},
		MinSteps: 13,
	},
	Game{
		Name: "AJAW 7",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{3, 2}, {3, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{6, 4},
		},
		MinSteps: 21,
	},
	Game{
		Name: "AJAW 8",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{4, 0, 0},
				},
			},
			Tiles: []Position{
				{1, 2}, {5, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 4},
		},
		MinSteps: 11,
	},
	Game{
		Name: "AJAW 9",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 1, 4},
				},
			},
			Tiles: []Position{
				{4, 2}, {2, 2},
			},
		},
		Goal: []Position{
			{3, 4}, {3, 0},
		},
		MinSteps: 17,
	},
	Game{
		Name: "AJAW 10",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{1, 4, 4},
					HorizontalWall{3, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 0}, {3, 4},
			},
		},
		Goal: []Position{
			{0, 0}, {6, 4},
		},
		MinSteps: 15,
	},
	Game{
		Name: "AJAW 11",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 6},
					HorizontalWall{1, 2, 6},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 4}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 1}, {6, 3},
		},
		MinSteps: 10,
	},
	Game{
		Name: "AJAW 12",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 4},
					VerticalWall{2, 1, 2},
					VerticalWall{4, 1, 2},
				},
			},
			Tiles: []Position{
				{0, 4}, {6, 4},
			},
		},
		Goal: []Position{
			{3, 2}, {3, 4},
		},
		MinSteps: 12,
	},
	Game{
		Name: "AJAW 13",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 6},
					HorizontalWall{2, 2, 6},
					HorizontalWall{3, 2, 6},
					HorizontalWall{4, 2, 6},
					VerticalWall{0, 2, 4},
				},
			},
			Tiles: []Position{
				{6, 1}, {5, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{1, 4},
		},
		MinSteps: 16,
	},
	Game{
		Name: "AJAW 14",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
			},
		},
		Goal: []Position{
			{3, 4},
		},
		MinSteps: 21,
	},
	Game{
		Name: "AJAW 15",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					HorizontalWall{1, 0, 1},
					HorizontalWall{3, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1},
			},
		},
		Goal: []Position{
			{3, 0}, {3, 4},
		},
		MinSteps: 16,
	},
	Game{
		Name: "AJAW 16",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 2, 2},
					HorizontalWall{1, 4, 4},
				},
			},
			Tiles: []Position{
				{1, 1}, {5, 1}, {3, 1},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 4},
		},
		MinSteps: 10,
	},
	Game{
		Name: "AJAW 17",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{4, 0, 0},
				},
			},
			Tiles: []Position{
				{6, 2}, {6, 1}, {6, 3},
			},
		},
		Goal: []Position{
			{0, 0},
		},
		MinSteps: 22,
	},
	Game{
		Name: "AJAW 18",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{5, 1}, {5, 3},
			},
		},
		Goal: []Position{
			{1, 3}, {1, 1},
		},
		MinSteps: 12,
	},
	Game{
		Name: "AJAW 19",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 2}, {3, 0}, {3, 1},
			},
		},
		Goal: []Position{
			{4, 4}, {2, 4},
		},
		MinSteps: 12,
	},
	Game{
		Name: "AJAW 20",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{1, 3}, {5, 1},
		},
		MinSteps: 10,
	},
	Game{
		Name: "AJAW 21",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 6},
					HorizontalWall{2, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 2}, {3, 3},
			},
		},
		Goal: []Position{
			{1, 1}, {6, 4},
		},
		MinSteps: 11,
	},
	Game{
		Name: "AJAW 22",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 4, 6},
					VerticalWall{6, 3, 4},
				},
			},
			Tiles: []Position{
				{0, 1}, {0, 3}, {0, 2},
			},
		},
		Goal: []Position{
			{1, 0}, {3, 4},
		},
		MinSteps: 10,
	},
	Game{
		Name: "AJAW 23",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 1, 4},
					VerticalWall{3, 1, 4},
					VerticalWall{6, 1, 4},
					HorizontalWall{4, 4, 5},
				},
			},
			Tiles: []Position{
				{2, 0}, {4, 0},
			},
		},
		Goal: []Position{
			{5, 3}, {1, 4},
		},
		MinSteps: 13,
	},
	Game{
		Name: "AJAW 24",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 4},
					VerticalWall{6, 1, 3},
					HorizontalWall{3, 2, 5},
					VerticalWall{0, 2, 4},
				},
			},
			Tiles: []Position{
				{4, 2}, {2, 2},
			},
		},
		Goal: []Position{
			{6, 4}, {0, 0},
		},
		MinSteps: 22,
	},
	Game{
		Name: "AJAW 25",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{5, 3}, {5, 1}, {5, 2},
			},
		},
		Goal: []Position{
			{3, 4}, {6, 0},
		},
		MinSteps: 9,
	},
	Game{
		Name: "AJAW 26",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					VerticalWall{6, 0, 4},
					HorizontalWall{4, 1, 1},
					HorizontalWall{1, 2, 2},
					HorizontalWall{2, 3, 3},
					HorizontalWall{3, 2, 2},
					HorizontalWall{3, 4, 4},
				},
			},
			Tiles: []Position{
				{5, 1}, {5, 0},
			},
		},
		Goal: []Position{
			{3, 3}, {2, 2},
		},
		MinSteps: 14,
	},
	Game{
		Name: "AJAW 27",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 0, 5},
					VerticalWall{5, 2, 4},
					HorizontalWall{3, 0, 3},
					VerticalWall{3, 4, 4},
				},
			},
			Tiles: []Position{
				{2, 0}, {2, 2}, {2, 4},
			},
		},
		Goal: []Position{
			{6, 4}, {4, 3}, {0, 4},
		},
		MinSteps: 16,
	},
	Game{
		Name: "AJAW 28",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{6, 0, 1},
					HorizontalWall{4, 1, 6},
				},
			},
			Tiles: []Position{
				{4, 3}, {2, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{6, 3}, {0, 0},
		},
		MinSteps: 13,
	},
	Game{
		Name: "AJAW 29",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					VerticalWall{2, 0, 3},
					HorizontalWall{3, 2, 4},
					VerticalWall{3, 2, 4},
					HorizontalWall{4, 3, 6},
					HorizontalWall{0, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 0}, {1, 0}, {4, 0}, {5, 0},
			},
		},
		Goal: []Position{
			{6, 1}, {1, 4},
		},
		MinSteps: 7,
	},
	Game{
		Name: "AJAW 30",
		Start: State{
			Board: &Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{4, 0, 1},
					HorizontalWall{1, 4, 5},
					VerticalWall{4, 3, 4},
					HorizontalWall{3, 4, 5},
				},
			},
			Tiles: []Position{
				{3, 2}, {2, 2}, {1, 2},
			},
		},
		Goal: []Position{
			{5, 0}, {5, 4}, {0, 2},
		},
		MinSteps: 14,
	},
}
