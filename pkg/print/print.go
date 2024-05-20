package print

import (
	"fmt"
	"gameoflife/pkg/board"
)

// PrintBoard prints the game board
func PrintBoard(g board.GameOfLife) {
	for _, row := range g.Board {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
