package cryptica

var Kuk = []Game{
	Game{
		Name: "K'UK 1",
		Start: State{
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
		MinSteps: 17,
	},
	Game{
		Name: "K'UK 2",
		Start: State{
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
		MinSteps: 19,
	},
	Game{
		Name: "K'UK 3",
		Start: State{
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
		MinSteps: 29,
	},
	Game{
		Name: "K'UK 4",
		Start: State{
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
		MinSteps: 17,
	},
	Game{
		Name: "K'UK 5",
		Start: State{
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
		MinSteps: 20,
	},
	Game{
		Name: "K'UK 6",
		Start: State{
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
		MinSteps: 20,
	},
	Game{
		Name: "K'UK 7",
		Start: State{
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
		MinSteps: 13,
	},
	Game{
		Name: "K'UK 8",
		Start: State{
			Board: Board{
				W: 7, H: 5,
				Walls: []Wall{
					VerticalWall{0, 0, 3},
					VerticalWall{4, 1, 3},
					VerticalWall{6, 0, 4},
					HorizontalWall{1, 2, 6},
					HorizontalWall{2, 4, 6},
					HorizontalWall{3, 0, 2},
				},
			},
			Tiles: []Position{
				{3, 3}, {1, 0}, {2, 0}, {1, 1}, {1, 2}, {2, 2}, {3, 2},
			},
		},
		Goal: []Position{
			{3, 0},
		},
		MinSteps: 25,
	},
	Game{
		Name: "K'UK 9",
		Start: State{
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
		MinSteps: 10,
	},
	Game{
		Name: "K'UK 10",
		Start: State{
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
		MinSteps: 20,
	},
	Game{
		Name: "K'UK 11",
		Start: State{
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
		MinSteps: 21,
	},
	Game{
		Name: "K'UK 12",
		Start: State{
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
		MinSteps: 23,
	},
	Game{
		Name: "K'UK 13",
		Start: State{
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
		MinSteps: 20,
	},
	Game{
		Name: "K'UK 14",
		Start: State{
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
		MinSteps: 15,
	},
	Game{
		Name: "K'UK 15",
		Start: State{
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
		MinSteps: 13,
	},
	Game{
		Name: "K'UK 16",
		Start: State{
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
		MinSteps: 18,
	},
	Game{
		Name: "K'UK 17",
		Start: State{
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
		MinSteps: 19,
	},
	Game{
		Name: "K'UK 18",
		Start: State{
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
			{1, 0}, {5, 0},
		},
		MinSteps: 20,
	},
	Game{
		Name: "K'UK 19",
		Start: State{
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
		MinSteps: 19,
	},
	Game{
		Name: "K'UK 20",
		Start: State{
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
		MinSteps: 22,
	},
	Game{
		Name: "K'UK 21",
		Start: State{
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
		MinSteps: 21,
	},
	Game{
		Name: "K'UK 22",
		Start: State{
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
		MinSteps: 14,
	},
	Game{
		Name: "K'UK 23",
		Start: State{
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
		MinSteps: 18,
	},
	Game{
		Name: "K'UK 24",
		Start: State{
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
		MinSteps: 22,
	},
	Game{
		Name: "K'UK 25",
		Start: State{
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
		MinSteps: 17,
	},
	Game{
		Name: "K'UK 26",
		Start: State{
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
		MinSteps: 15,
	},
	Game{
		Name: "K'UK 27",
		Start: State{
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
		MinSteps: 18,
	},
	Game{
		Name: "K'UK 28",
		Start: State{
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
		MinSteps: 17,
	},
	Game{
		Name: "K'UK 29",
		Start: State{
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
		MinSteps: 27,
	},
	Game{
		Name: "K'UK 30",
		Start: State{
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
		MinSteps: 25,
	},
}
