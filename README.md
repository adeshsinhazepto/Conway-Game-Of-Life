# Game of life 
This project implements Conway's Game of Life in Go. The project includes functionality to initialize a game board with random values, update the board based on the game's rules, and print the board to the console.
Number of rows and columns are the inputs which by given by the user. 

# Testing
This project includes unit tests to ensure the functionality of the Game of Life implementation. This test cover the initialization of the game board
To run the tests, navigate to the project directory in your terminal and use the following command:go test ./..

# How to run
First, Install Dependencies
go mod tidy
Running the Game Navigate to the cmd/game directory and run the main program:
cd cmd/game go run main.go

# File overview
pkg file contains the main logic for the Game of Life.
GameOfLife: Struct that holds the game board and its dimensions.
NewGameOfLife: Initializes a new game with random values.
countLiveNeighbourCell: Counts the live neighbors of a cell.
UpdateBoard: Updates the game board based on the rules of the Game of Life.
PrintBoard: Prints the current state of the game board.

