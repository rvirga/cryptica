package cryptica

var Ajaw = []Game{
	// AJAW 1
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{4, 4}, {2, 4},
			},
		},
		Goal: []Position{
			{6, 0}, {5, 0},
		},
		Steps: 7,
	},
	// AJAW 2
	Game{
		State: State{
			Board: Board{
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
		Steps: 10,
	},
	// AJAW 3
	Game{
		State: State{
			Board: Board{
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
		Steps: 6,
	},
	// AJAW 4
	Game{
		State: State{
			Board: Board{
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
		Steps: 8,
	},
	// AJAW 5
	Game{
		State: State{
			Board: Board{
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
		Steps: 12,
	},
	// AJAW 6
	Game{
		State: State{
			Board: Board{
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
		Steps: 13,
	},
	// AJAW 7
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{3, 2}, {3, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{6, 4},
		},
		Steps: 21,
	},
	// AJAW 8
	Game{
		State: State{
			Board: Board{
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
		Steps: 11,
	},
	// AJAW 9
	Game{
		State: State{
			Board: Board{
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
		Steps: 17,
	},
	// AJAW 10
	Game{
		State: State{
			Board: Board{
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
		Steps: 15,
	},
	// AJAW 11
	Game{
		State: State{
			Board: Board{
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
		Steps: 10,
	},
	// AJAW 12
	Game{
		State: State{
			Board: Board{
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
		Steps: 12,
	},
	// AJAW 13
	Game{
		State: State{
			Board: Board{
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
		Steps: 16,
	},
	// AJAW 14
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
			},
		},
		Goal: []Position{
			{3, 4},
		},
		Steps: 21,
	},
	// AJAW 15
	Game{
		State: State{
			Board: Board{
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
		Steps: 16,
	},
	// AJAW 16
	Game{
		State: State{
			Board: Board{
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
		Steps: 10,
	},
	// AJAW 17
	Game{
		State: State{
			Board: Board{
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
		Steps: 22,
	},
	// AJAW 18
	Game{
		State: State{
			Board: Board{
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
		Steps: 12,
	},
	// AJAW 19
	Game{
		State: State{
			Board: Board{
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
			{2, 4}, {4, 4},
		},
		Steps: 12,
	},
	// AJAW 20
	Game{
		State: State{
			Board: Board{
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
		Steps: 10,
	},
	// AJAW 21
	Game{
		State: State{
			Board: Board{
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
		Steps: 11,
	},
	// AJAW 22
	Game{
		State: State{
			Board: Board{
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
		Steps: 10,
	},
	// AJAW 23
	Game{
		State: State{
			Board: Board{
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
			{1, 4}, {5, 3},
		},
		Steps: 13,
	},
	// AJAW 24
	Game{
		State: State{
			Board: Board{
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
		Steps: 22,
	},
	// AJAW 25
	Game{
		State: State{
			Board: Board{
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
		Steps: 9,
	},
	// AJAW 26
	Game{
		State: State{
			Board: Board{
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
		Steps: 14,
	},
	// AJAW 27
	Game{
		State: State{
			Board: Board{
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
		Steps: 16,
	},
	// AJAW 28
	Game{
		State: State{
			Board: Board{
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
		Steps: 13,
	},
	// AJAW 29
	Game{
		State: State{
			Board: Board{
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
		Steps: 7,
	},
	// AJAW 30
	Game{
		State: State{
			Board: Board{
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
		Steps: 14,
	},
}
