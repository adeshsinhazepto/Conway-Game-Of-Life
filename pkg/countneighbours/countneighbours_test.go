package countneighbours

import (
	"gameoflife/pkg/board"
	"testing"
)

func TestCountLiveNeighbours(t *testing.T) {
	t.Run("should be able to count live neighbours", func(t *testing.T) {
		game := board.GameOfLife{
			Rows: 3,
			Cols: 3,
			Board: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
		}
		count := CountLiveNeighbourCell(1, 1, game)
		expected := 4
		if count != expected {
			t.Errorf("Expected %d live neighbors, got %d", expected, count)
		}
	})
}
