package board

import (
	"testing"
)

func TestNewGameOfLife(t *testing.T) {
	t.Run("should be able to return something", func(t *testing.T) {
		board := NewGameOfLife(10, 10)
		if board == nil {
			t.Error("NewGameOfLife(10, 10) should not return nil")
		}
	})
	t.Run("should be able to return correct values", func(t *testing.T) {
		rows, cols := 5, 5
		game := NewGameOfLife(rows, cols)

		if game.Rows != rows || game.Cols != cols {
			t.Errorf("Expected rows %d and cols %d, got rows %d and cols %d", rows, cols, game.Rows, game.Cols)
		}

		if len(game.Board) != rows {
			t.Errorf("Expected board length %d, got %d", rows, len(game.Board))
		}

		for i := range game.Board {
			if len(game.Board[i]) != cols {
				t.Errorf("Expected board[%d] length %d, got %d", i, cols, len(game.Board[i]))
			}
		}
	})

}
