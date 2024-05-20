package print

import (
	"gameoflife/pkg/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintBoard(t *testing.T) {
	game := board.GameOfLife{
		Rows: 3,
		Cols: 3,
		Board: [][]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
		},
	}
	assert.NotPanics(t, func() { PrintBoard(game) })

}
