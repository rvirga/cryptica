package cryptica

var Chan = []Game{
	// CHAN 1
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 0, 2},
					HorizontalWall{2, 5, 6},
					VerticalWall{2, 2, 3},
					VerticalWall{4, 1, 2},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 0}, {3, 1}, {3, 2}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 20,
	},
	// CHAN 2
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{4, 0, 0},
				},
			},
			Tiles: []Position{
				{2, 0}, {2, 1}, {2, 2}, {2, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3},
			},
		},
		Goal: []Position{
			{0, 3}, {0, 2}, {0, 1}, {0, 0},
		},
		Steps: 29,
	},
	// CHAN 3
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{4, 3, 3},
					VerticalWall{6, 3, 4},
				},
			},
			Tiles: []Position{
				{4, 4}, {5, 4}, {5, 3}, {6, 0}, {6, 1}, {6, 2},
			},
		},
		Goal: []Position{
			{1, 4}, {3, 2}, {5, 0},
		},
		Steps: 27,
	},
	// CHAN 4
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					VerticalWall{6, 0, 4},
					VerticalWall{2, 3, 4},
					HorizontalWall{4, 1, 1},
				},
			},
			Tiles: []Position{
				{3, 4}, {5, 0}, {4, 0}, {5, 1}, {3, 3},
			},
		},
		Goal: []Position{
			{5, 4}, {1, 0},
		},
		Steps: 15,
	},
	// CHAN 5
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 3},
					VerticalWall{1, 0, 3},
					HorizontalWall{0, 5, 6},
					HorizontalWall{1, 5, 6},
					HorizontalWall{3, 5, 6},
					HorizontalWall{4, 5, 6},
				},
			},
			Tiles: []Position{
				{2, 3}, {3, 2}, {4, 1}, {4, 3},
			},
		},
		Goal: []Position{
			{1, 4}, {5, 2}, {3, 0},
		},
		Steps: 13,
	},
	// CHAN 6
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 4},
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{2, 1}, {3, 1}, {4, 1}, {5, 0},
			},
		},
		Goal: []Position{
			{0, 0}, {6, 0}, {3, 4},
		},
		Steps: 20,
	},
	// CHAN 7
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 2},
					HorizontalWall{0, 5, 5},
					HorizontalWall{4, 1, 3},
				},
			},
			Tiles: []Position{
				{3, 2}, {1, 2}, {5, 2}, {2, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{2, 1}, {6, 0}, {0, 4},
		},
		Steps: 30,
	},
	// CHAN 8
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 6},
					HorizontalWall{1, 4, 6},
					HorizontalWall{3, 0, 2},
					HorizontalWall{4, 0, 3},
				},
			},
			Tiles: []Position{
				{1, 1}, {5, 3}, {2, 0}, {1, 2}, {5, 2}, {4, 4},
			},
		},
		Goal: []Position{
			{0, 2}, {6, 2},
		},
		Steps: 20,
	},
	// CHAN 9
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 2, 3},
					HorizontalWall{2, 6, 6},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{1, 1}, {1, 3}, {5, 3}, {5, 1}, {3, 1}, {3, 2}, {3, 3},
			},
		},
		Goal: []Position{
			{1, 4}, {1, 0}, {5, 0}, {5, 4},
		},
		Steps: 24,
	},
	// CHAN 10
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 2},
					HorizontalWall{0, 4, 5},
					VerticalWall{6, 0, 1},
					VerticalWall{0, 3, 4},
					HorizontalWall{4, 1, 6},
				},
			},
			Tiles: []Position{
				{3, 1}, {4, 2}, {2, 2}, {1, 1}, {5, 1}, {1, 3}, {5, 3},
			},
		},
		Goal: []Position{
			{3, 0}, {6, 2}, {0, 1},
		},
		Steps: 15,
	},
	// CHAN 11
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 6},
					VerticalWall{0, 0, 3},
					VerticalWall{1, 0, 2},
					VerticalWall{2, 0, 1},
					VerticalWall{4, 0, 1},
					VerticalWall{5, 0, 2},
					VerticalWall{6, 0, 3},
				},
			},
			Tiles: []Position{
				{0, 4}, {6, 4}, {3, 2}, {2, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{5, 4}, {1, 4},
		},
		Steps: 20,
	},
	// CHAN 12
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 6, 6},
					HorizontalWall{1, 0, 4},
					VerticalWall{2, 1, 2},
					VerticalWall{4, 1, 2},
					HorizontalWall{4, 0, 1},
					HorizontalWall{4, 3, 4},
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{0, 0}, {1, 3}, {1, 0}, {0, 3}, {5, 1}, {5, 2}, {5, 3},
			},
		},
		Goal: []Position{
			{6, 2}, {0, 2}, {5, 0}, {2, 4},
		},
		Steps: 22,
	},
	// CHAN 13
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 1, 5},
					VerticalWall{1, 2, 3},
					VerticalWall{2, 1, 2},
					VerticalWall{3, 2, 3},
					VerticalWall{4, 1, 2},
					VerticalWall{5, 2, 3},
				},
			},
			Tiles: []Position{
				{0, 4}, {6, 4}, {0, 0}, {6, 0}, {2, 0}, {4, 0},
			},
		},
		Goal: []Position{
			{2, 3}, {4, 3}, {1, 1}, {5, 1},
		},
		Steps: 19,
	},
	// CHAN 14
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 2, 2},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 1, 1},
					HorizontalWall{2, 5, 5},
					HorizontalWall{3, 0, 0},
					HorizontalWall{3, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1}, {3, 2}, {3, 4}, {0, 0}, {3, 0}, {6, 0},
			},
		},
		Goal: []Position{
			{4, 4}, {5, 0}, {1, 0}, {2, 4},
		},
		Steps: 26,
	},
	// CHAN 15
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 1},
					VerticalWall{3, 0, 2},
					HorizontalWall{1, 6, 6},
					HorizontalWall{3, 2, 2},
					VerticalWall{5, 3, 4},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{0, 2}, {0, 3}, {6, 2}, {6, 3},
			},
		},
		Goal: []Position{
			{2, 1}, {0, 4}, {5, 1}, {4, 3},
		},
		Steps: 17,
	},
	// CHAN 16
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 4, 4},
					HorizontalWall{1, 0, 0},
					HorizontalWall{2, 6, 6},
					HorizontalWall{4, 1, 1},
					HorizontalWall{4, 4, 4},
				},
			},
			Tiles: []Position{
				{1, 2}, {2, 2}, {5, 2}, {3, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {5, 4}, {6, 0},
		},
		Steps: 23,
	},
	// CHAN 17
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					HorizontalWall{0, 5, 5},
					HorizontalWall{2, 2, 2},
					HorizontalWall{3, 0, 0},
				},
			},
			Tiles: []Position{
				{4, 3}, {5, 2}, {5, 3}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {3, 1}, {6, 0}, {1, 0},
		},
		Steps: 20,
	},
	// CHAN 18
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					HorizontalWall{1, 2, 2},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 1, 1},
					HorizontalWall{2, 5, 5},
					HorizontalWall{3, 2, 2},
					HorizontalWall{3, 4, 4},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{3, 1}, {2, 4}, {4, 4}, {0, 0}, {4, 2}, {3, 3}, {0, 4},
			},
		},
		Goal: []Position{
			{3, 2}, {2, 0}, {4, 0},
		},
		Steps: 25,
	},
	// CHAN 19
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 6},
					HorizontalWall{1, 0, 1},
					HorizontalWall{1, 3, 6},
					HorizontalWall{3, 0, 2},
					HorizontalWall{3, 4, 4},
					HorizontalWall{3, 6, 6},
					HorizontalWall{4, 0, 6},
				},
			},
			Tiles: []Position{
				{3, 2}, {4, 2}, {5, 2}, {6, 2},
			},
		},
		Goal: []Position{
			{2, 1}, {3, 3}, {0, 2}, {5, 3},
		},
		Steps: 11,
	},
	// CHAN 20
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 1},
					HorizontalWall{0, 5, 5},
					HorizontalWall{1, 2, 2},
					HorizontalWall{2, 3, 3},
					HorizontalWall{3, 4, 4},
					HorizontalWall{4, 1, 1},
					HorizontalWall{4, 5, 5},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 2}, {3, 3}, {3, 1}, {4, 1}, {2, 3},
			},
		},
		Goal: []Position{
			{0, 2}, {6, 2}, {3, 4}, {2, 0},
		},
		Steps: 14,
	},
	// CHAN 21
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 4, 6},
					HorizontalWall{1, 5, 6},
					HorizontalWall{2, 6, 6},
					HorizontalWall{2, 0, 0},
					HorizontalWall{3, 0, 1},
					HorizontalWall{4, 0, 2},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 2}, {2, 1}, {4, 1}, {3, 2}, {2, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{1, 0}, {5, 4},
		},
		Steps: 16,
	},
	// CHAN 22
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 1},
					HorizontalWall{0, 5, 5},
					HorizontalWall{2, 0, 0},
					HorizontalWall{2, 6, 6},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{4, 0}, {2, 0}, {4, 1}, {2, 1}, {4, 2}, {2, 2},
			},
		},
		Goal: []Position{
			{6, 0}, {5, 1}, {6, 4}, {0, 4},
		},
		Steps: 18,
	},
	// CHAN 23
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					HorizontalWall{0, 6, 6},
					HorizontalWall{2, 3, 3},
					VerticalWall{0, 2, 4},
					VerticalWall{6, 3, 4},
				},
			},
			Tiles: []Position{
				{5, 0}, {1, 0}, {3, 1}, {1, 1}, {2, 1}, {4, 1}, {5, 1},
			},
		},
		Goal: []Position{
			{6, 2}, {0, 0}, {3, 4},
		},
		Steps: 16,
	},
	// CHAN 24
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 2},
					HorizontalWall{4, 4, 4},
				},
			},
			Tiles: []Position{
				{5, 4}, {3, 0}, {0, 0}, {4, 0}, {6, 4},
			},
		},
		Goal: []Position{
			{6, 0}, {1, 4},
		},
		Steps: 26,
	},
	// CHAN 25
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 3, 3},
					HorizontalWall{1, 6, 6},
					HorizontalWall{3, 0, 0},
					HorizontalWall{3, 3, 3},
					HorizontalWall{4, 5, 5},
				},
			},
			Tiles: []Position{
				{3, 0}, {3, 4}, {4, 1}, {2, 3},
			},
		},
		Goal: []Position{
			{0, 2}, {6, 2},
		},
		Steps: 18,
	},
	// CHAN 26
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 6, 6},
					HorizontalWall{1, 5, 5},
					HorizontalWall{2, 4, 4},
					HorizontalWall{2, 2, 2},
					HorizontalWall{3, 1, 1},
					HorizontalWall{4, 0, 0},
				},
			},
			Tiles: []Position{
				{3, 1}, {6, 4}, {3, 3}, {0, 0},
			},
		},
		Goal: []Position{
			{1, 2}, {5, 2}, {1, 4},
		},
		Steps: 20,
	},
	// CHAN 27
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 0},
					HorizontalWall{0, 3, 3},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 3, 3},
					HorizontalWall{3, 2, 2},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{0, 4}, {6, 0}, {4, 0}, {2, 4},
			},
		},
		Goal: []Position{
			{2, 2}, {3, 3},
		},
		Steps: 20,
	},
	// CHAN 28
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{3, 0, 1},
					HorizontalWall{1, 1, 1},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 0, 6},
					HorizontalWall{3, 3, 3},
					HorizontalWall{4, 1, 1},
					HorizontalWall{4, 3, 4},
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{2, 0}, {4, 0}, {2, 4}, {5, 3},
			},
		},
		Goal: []Position{
			{0, 1}, {6, 0}, {2, 3}, {5, 4},
		},
		Steps: 10,
	},
	// CHAN 29
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 2},
					HorizontalWall{4, 4, 5},
					VerticalWall{0, 0, 4},
					VerticalWall{6, 0, 4},
					VerticalWall{1, 3, 4},
					VerticalWall{5, 0, 1},
				},
			},
			Tiles: []Position{
				{3, 1}, {3, 3}, {2, 1}, {4, 1}, {2, 3}, {4, 3},
			},
		},
		Goal: []Position{
			{1, 2}, {5, 2},
		},
		Steps: 14,
	},
	// CHAN 30
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 1},
					VerticalWall{6, 0, 1},
					VerticalWall{1, 1, 2},
					HorizontalWall{2, 5, 5},
				},
			},
			Tiles: []Position{
				{6, 2}, {6, 3}, {6, 4}, {0, 2}, {0, 3}, {0, 4},
			},
		},
		Goal: []Position{
			{1, 0}, {5, 0}, {3, 4},
		},
		Steps: 32,
	},
}
