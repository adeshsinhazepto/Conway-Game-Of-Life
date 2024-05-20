package countliveneighbours

import (
	"testing"
)

func TestCountLiveNeighbours(t *testing.T) {
	t.Run("should be able to count live neighbours", func(t *testing.T) {
		board := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}

		if count := CountLiveNeighbours(board, 1, 1); count != 0 {
			t.Errorf("Expected count 0, got %d", count)
		}

		board = [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}

		if count := CountLiveNeighbours(board, 1, 1); count != 8 {
			t.Errorf("Expected count 8, got %d", count)
		}
	})
}
