package updateboard

import (
	"gameoflife/pkg/board"
	"reflect"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	t.Run("should be able to  update board", func(t *testing.T) {
		game := board.GameOfLife{
			Rows: 3,
			Cols: 3,
			Board: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		}
		gameBoard := *UpdateBoard(game)
		gameBoardexpected := board.GameOfLife{
			Rows: 3,
			Cols: 3,
			Board: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		}
		if !reflect.DeepEqual(gameBoard, gameBoardexpected) {
			t.Errorf("Expected %v got %v", gameBoardexpected, gameBoard)
		}
	})
}
