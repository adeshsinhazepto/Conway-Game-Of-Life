package main

import (
	"fmt"
	"gameoflife/pkg/board"
	"gameoflife/pkg/print"
	"gameoflife/pkg/updateboard"
)

func main() {
	var numberOfRows, numberOfCols int
	fmt.Println("Enter the Number of rows:")
	fmt.Scanln(&numberOfRows)
	fmt.Println("Enter the Number of columns:")
	fmt.Scanln(&numberOfCols)
	newBoard := board.NewGameOfLife(numberOfRows, numberOfCols)
	fmt.Println("Initial Game: ")
	print.PrintBoard(*newBoard)
	updateBoard := updateboard.UpdateBoard(*newBoard)
	fmt.Println("Final Game: ")
	print.PrintBoard(*updateBoard)
}
