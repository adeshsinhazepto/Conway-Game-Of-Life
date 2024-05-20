package board

import (
	"math/rand"
)

type GameOfLife struct {
	Rows, Cols int
	Board      [][]int
}

// NewGameOfLife initializes a new Game of Life with random values
func NewGameOfLife(rows, cols int) *GameOfLife {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
		for j := range board[i] {
			board[i][j] = rand.Intn(2)
		}
	}
	return &GameOfLife{Rows: rows, Cols: cols, Board: board}
}
