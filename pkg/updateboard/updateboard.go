package updateboard

import (
	"gameoflife/pkg/board"
	"gameoflife/pkg/countneighbours"
)

func UpdateBoard(g board.GameOfLife) *board.GameOfLife {
	newBoard := make([][]int, g.Rows)
	for i := range newBoard {
		newBoard[i] = make([]int, g.Cols)
	}
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			neighbourLiveCell := countneighbours.CountLiveNeighbourCell(i, j, g)
			if g.Board[i][j] == 1 && (neighbourLiveCell == 2 || neighbourLiveCell == 3) {
				newBoard[i][j] = 1
			} else if g.Board[i][j] == 0 && neighbourLiveCell == 3 {
				newBoard[i][j] = 1
			} else {
				newBoard[i][j] = 0
			}
		}
	}
	g.Board = newBoard
	return &g
}
