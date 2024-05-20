package countneighbours

import (
	"gameoflife/pkg/board"
)

func CountLiveNeighbourCell(r int, c int, g board.GameOfLife) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i == r && j == c) || i < 0 || j < 0 || i >= g.Rows || j >= g.Cols {
				continue
			}
			if g.Board[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
