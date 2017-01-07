package cryptica

var Balam = []Game{
	// B'ALAM 1
	Game{
		State: State{
			Board: Board{
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
		Steps: 17,
	},
	// B'ALAM 2
	Game{
		State: State{
			Board: Board{
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
		Steps: 16,
	},
	// B'ALAM 3
	Game{
		State: State{
			Board: Board{
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
		Steps: 9,
	},
	// B'ALAM 4
	Game{
		State: State{
			Board: Board{
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
		Steps: 12,
	},
	// B'ALAM 5
	Game{
		State: State{
			Board: Board{
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
		Steps: 19,
	},
	// B'ALAM 6
	Game{
		State: State{
			Board: Board{
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
		Steps: 24,
	},
	// B'ALAM 7
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{3, 2}, {2, 1}, {3, 1}, {4, 0}, {5, 0}, {6, 0},
			},
		},
		Goal: []Position{
			{2, 2}, {1, 1}, {1, 2},
		},
		Steps: 20,
	},
	// B'ALAM 8
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 9
	Game{
		State: State{
			Board: Board{
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
		Steps: 27,
	},
	// B'ALAM 10
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 11
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 12
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
			},
			Tiles: []Position{
				{0, 0}, {2, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
			},
		},
		Goal: []Position{
			{5, 4}, {6, 0},
		},
		Steps: 99,
	},
	// B'ALAM 13
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 14
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 15
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 16
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 17
	Game{
		State: State{
			Board: Board{
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
		Steps: 99,
	},
	// B'ALAM 18
	Game{
		State: State{
			Board: Board{
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
		Steps: 19,
	},
	// B'ALAM 19
	Game{
		State: State{
			Board: Board{
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
		Steps: 17,
	},
	// B'ALAM 20
	Game{
		State: State{
			Board: Board{
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
		Steps: 11,
	},
	// B'ALAM 21
	Game{
		State: State{
			Board: Board{
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
		Steps: 13,
	},
	// B'ALAM 22
	Game{
		State: State{
			Board: Board{
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
		Steps: 22,
	},
	// B'ALAM 23
	Game{
		State: State{
			Board: Board{
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
		Steps: 11,
	},
	// B'ALAM 24
	Game{
		State: State{
			Board: Board{
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
		Steps: 29,
	},
	// B'ALAM 25
	Game{
		State: State{
			Board: Board{
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
		Steps: 18,
	},
	// B'ALAM 26
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{5, 0, 2},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 2}, {2, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 16,
	},
	// B'ALAM 27
	Game{
		State: State{
			Board: Board{
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
		Steps: 13,
	},
	// B'ALAM 28
	Game{
		State: State{
			Board: Board{
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
		Steps: 9,
	},
	// B'ALAM 29
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 2, 5},
					VerticalWall{1, 1, 4},
					VerticalWall{5, 1, 4},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{3, 2}, {0, 4}, {6, 0}, {4, 2}, {4, 4},
			},
		},
		Goal: []Position{
			{3, 1}, {0, 0}, {6, 4},
		},
		Steps: 15,
	},
	// B'ALAM 30
	Game{
		State: State{
			Board: Board{
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
		Steps: 13,
	},
}
