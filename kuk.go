package cryptica

var Kuk = []Game{
	// K'UK 1
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 3, 3},
					HorizontalWall{2, 0, 0},
				},
			},
			Tiles: []Position{
				{0, 4}, {1, 4}, {0, 3},
			},
		},
		Goal: []Position{
			{2, 0}, {6, 4},
		},
		Steps: 17,
	},
	// K'UK 2
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{1, 1, 4},
					VerticalWall{3, 1, 3},
					VerticalWall{5, 0, 1},
					HorizontalWall{2, 4, 5},
				},
			},
			Tiles: []Position{
				{2, 4}, {2, 1}, {2, 2}, {2, 3},
			},
		},
		Goal: []Position{
			{6, 0}, {0, 0}, {4, 1}, {4, 3},
		},
		Steps: 19,
	},
	// K'UK 3
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 2},
					HorizontalWall{0, 4, 6},
					HorizontalWall{2, 1, 4},
					HorizontalWall{4, 0, 0},
					VerticalWall{4, 3, 4},
				},
			},
			Tiles: []Position{
				{1, 1}, {3, 0}, {5, 1}, {2, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{2, 3}, {3, 3}, {5, 4},
		},
		Steps: 29,
	},
	// K'UK 4
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					HorizontalWall{1, 5, 5},
					HorizontalWall{3, 1, 1},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 2}, {3, 1}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 17,
	},
	// K'UK 5
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 3, 3},
				},
			},
			Tiles: []Position{
				{2, 0}, {4, 0}, {3, 4}, {2, 2}, {4, 2}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 4}, {3, 0},
		},
		Steps: 20,
	},
	// K'UK 6
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 0, 1},
					VerticalWall{2, 3, 4},
				},
			},
			Tiles: []Position{
				{0, 1}, {1, 1}, {1, 4}, {0, 0}, {1, 0},
			},
		},
		Goal: []Position{
			{4, 1}, {6, 1}, {0, 4},
		},
		Steps: 20,
	},
	// K'UK 7
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 2, 4},
					VerticalWall{2, 0, 2},
					VerticalWall{4, 2, 4},
					VerticalWall{6, 0, 1},
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{6, 2}, {0, 1}, {3, 1}, {3, 2}, {3, 3},
			},
		},
		Goal: []Position{
			{6, 3}, {0, 0},
		},
		Steps: 13,
	},
	// K'UK 8
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 3},
					VerticalWall{4, 1, 3},
					VerticalWall{6, 0, 4},
					HorizontalWall{1, 2, 6},
					HorizontalWall{2, 4, 6},
				},
			},
			Tiles: []Position{
				{3, 3}, {1, 0}, {2, 0}, {1, 1}, {1, 2}, {2, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{3, 0},
		},
		Steps: 25,
	},
	// K'UK 9
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 3, 3},
					HorizontalWall{1, 0, 0},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 1, 1},
					HorizontalWall{2, 5, 5},
					HorizontalWall{3, 2, 2},
					HorizontalWall{3, 6, 6},
					HorizontalWall{4, 3, 3},
				},
			},
			Tiles: []Position{
				{0, 4}, {6, 0}, {3, 2}, {4, 0}, {5, 1}, {1, 3},
			},
		},
		Goal: []Position{
			{0, 3}, {6, 2}, {1, 0},
		},
		Steps: 10,
	},
	// K'UK 10
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{2, 0, 1},
					HorizontalWall{2, 2, 4},
					VerticalWall{4, 2, 4},
				},
			},
			Tiles: []Position{
				{0, 0}, {6, 0}, {6, 2}, {6, 3}, {6, 4},
			},
		},
		Goal: []Position{
			{3, 3}, {3, 1},
		},
		Steps: 20,
	},
	// K'UK 11
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 1},
					HorizontalWall{1, 3, 3},
					HorizontalWall{1, 5, 5},
					HorizontalWall{3, 1, 1},
					HorizontalWall{3, 3, 3},
					HorizontalWall{3, 5, 5},
				},
			},
			Tiles: []Position{
				{0, 2}, {2, 2}, {4, 2}, {6, 2},
			},
		},
		Goal: []Position{
			{1, 4}, {5, 4}, {1, 0}, {5, 0},
		},
		Steps: 21,
	},
	// K'UK 12
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 1},
				},
			},
			Tiles: []Position{
				{2, 2}, {3, 2}, {4, 2}, {2, 1}, {3, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0}, {6, 4},
		},
		Steps: 23,
	},
	// K'UK 13
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 6},
					HorizontalWall{1, 0, 1},
					HorizontalWall{1, 3, 6},
					VerticalWall{0, 3, 4},
					HorizontalWall{4, 1, 6},
				},
			},
			Tiles: []Position{
				{4, 3}, {4, 2}, {2, 2}, {3, 2}, {2, 3}, {3, 3},
			},
		},
		Goal: []Position{
			{0, 2}, {5, 2},
		},
		Steps: 20,
	},
	// K'UK 14
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 2, 2},
					HorizontalWall{3, 4, 4},
				},
			},
			Tiles: []Position{
				{0, 0}, {6, 0}, {6, 4}, {0, 4},
			},
		},
		Goal: []Position{
			{3, 0}, {5, 2}, {3, 4}, {1, 2},
		},
		Steps: 15,
	},
	// K'UK 15
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 4},
					VerticalWall{1, 1, 4},
					VerticalWall{2, 2, 4},
					VerticalWall{3, 3, 4},
					HorizontalWall{4, 4, 4},
				},
			},
			Tiles: []Position{
				{4, 0}, {6, 2}, {5, 0}, {5, 1}, {6, 1},
			},
		},
		Goal: []Position{
			{2, 1}, {5, 4},
		},
		Steps: 13,
	},
	// K'UK 16
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{6, 0, 1},
					VerticalWall{6, 3, 4},
				},
			},
			Tiles: []Position{
				{5, 1}, {5, 2}, {5, 3}, {1, 0}, {1, 1}, {1, 3}, {1, 4},
			},
		},
		Goal: []Position{
			{0, 3}, {0, 2}, {0, 1},
		},
		Steps: 18,
	},
	// K'UK 17
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 1},
					HorizontalWall{3, 2, 2},
				},
			},
			Tiles: []Position{
				{2, 2}, {1, 2}, {0, 2},
			},
		},
		Goal: []Position{
			{2, 4}, {3, 0}, {4, 4},
		},
		Steps: 19,
	},
	// K'UK 18
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{4, 0, 0},
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{3, 4}, {3, 0}, {3, 2}, {1, 4}, {5, 4},
			},
		},
		Goal: []Position{
			{5, 3}, {1, 1},
		},
		Steps: 20,
	},
	// K'UK 19
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 1, 6},
					HorizontalWall{1, 4, 6},
				},
			},
			Tiles: []Position{
				{3, 3}, {4, 3}, {0, 4},
			},
		},
		Goal: []Position{
			{0, 1}, {6, 2},
		},
		Steps: 19,
	},
	// K'UK 20
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{3, 0, 1},
					HorizontalWall{1, 1, 2},
					VerticalWall{3, 3, 4},
					HorizontalWall{3, 4, 5},
				},
			},
			Tiles: []Position{
				{2, 2}, {4, 2}, {0, 2}, {6, 2},
			},
		},
		Goal: []Position{
			{0, 0}, {6, 4},
		},
		Steps: 22,
	},
	// K'UK 21
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 0},
					HorizontalWall{0, 5, 6},
					HorizontalWall{1, 4, 4},
					HorizontalWall{2, 3, 3},
					HorizontalWall{3, 2, 2},
					HorizontalWall{4, 1, 1},
					HorizontalWall{4, 6, 6},
				},
			},
			Tiles: []Position{
				{5, 3}, {1, 1}, {4, 0}, {0, 4},
			},
		},
		Goal: []Position{
			{4, 2}, {2, 2},
		},
		Steps: 21,
	},
	// K'UK 22
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 1, 1},
					HorizontalWall{3, 1, 1},
				},
			},
			Tiles: []Position{
				{2, 1}, {4, 1}, {4, 3}, {2, 3},
			},
		},
		Goal: []Position{
			{6, 0}, {5, 2}, {0, 3}, {0, 1},
		},
		Steps: 14,
	},
	// K'UK 23
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{2, 1, 2},
				},
			},
			Tiles: []Position{
				{3, 2}, {4, 1}, {4, 2}, {3, 1},
			},
		},
		Goal: []Position{
			{0, 0}, {2, 3}, {6, 2},
		},
		Steps: 18,
	},
	// K'UK 24
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 5, 5},
				},
			},
			Tiles: []Position{
				{2, 1}, {4, 1}, {3, 1}, {1, 1},
			},
		},
		Goal: []Position{
			{0, 4}, {3, 4}, {6, 4},
		},
		Steps: 22,
	},
	// K'UK 25
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 0},
					HorizontalWall{0, 5, 5},
					HorizontalWall{2, 1, 1},
					HorizontalWall{2, 5, 5},
					HorizontalWall{4, 1, 1},
					HorizontalWall{4, 5, 6},
				},
			},
			Tiles: []Position{
				{2, 4}, {4, 4}, {2, 0}, {3, 0}, {4, 0},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 17,
	},
	// K'UK 26
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{1, 0, 2},
					HorizontalWall{3, 4, 6},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 15,
	},
	// K'UK 27
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{2, 2, 4},
				},
			},
			Tiles: []Position{
				{3, 2}, {4, 2}, {5, 2}, {3, 3}, {5, 3},
			},
		},
		Goal: []Position{
			{0, 4}, {2, 0}, {6, 4},
		},
		Steps: 18,
	},
	// K'UK 28
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 1},
					HorizontalWall{1, 2, 2},
					HorizontalWall{1, 4, 4},
					HorizontalWall{3, 2, 2},
					HorizontalWall{3, 4, 4},
				},
			},
			Tiles: []Position{
				{3, 3}, {3, 1}, {2, 2}, {3, 2}, {4, 2},
			},
		},
		Goal: []Position{
			{0, 4}, {6, 0},
		},
		Steps: 17,
	},
	// K'UK 29
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{0, 0, 0},
					HorizontalWall{0, 6, 6},
				},
			},
			Tiles: []Position{
				{1, 1}, {3, 4}, {5, 1}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
			},
		},
		Goal: []Position{
			{1, 4}, {5, 4}, {3, 1},
		},
		Steps: 27,
	},
	// K'UK 30
	Game{
		State: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					HorizontalWall{2, 1, 5},
					VerticalWall{3, 3, 4},
					HorizontalWall{4, 0, 0},
				},
			},
			Tiles: []Position{
				{2, 0}, {4, 0}, {3, 0}, {2, 1}, {4, 1},
			},
		},
		Goal: []Position{
			{2, 3}, {4, 3}, {3, 1},
		},
		Steps: 25,
	},
}
