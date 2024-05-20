# Game of life 
This project implements Conway's Game of Life in Go. The project includes functionality to initialize a game board with random values, update the board based on the game's rules, and print the board to the console.
Number of rows and columns are the inputs which by given by the user. 

# How to Run
1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the directory containing the program.
4. Run `go mod tidy` to ensure all the dependencies are in place.
5. Run the program using the command `go run main.go` inside the cmd/gameoflike folder or directly type `make run` in the root directory to run the program.
6. Follow the prompts to input the number of rows and columns

# How to Use
Run the program.
When prompted, enter the number of rows and then hit enter and enter the number of columns and again hit enter.
The program will first display the initial game by putting random values in the matrix and then its final game of life matrix .

# Testing
This project includes unit tests to ensure the functionality of the Game of Life implementation. This test cover the initialization of the game board
To run the tests, navigate to the project directory in your terminal and use the following command:go test ./.. or type make test in the root directory to run all the tests

# File overview
pkg file contains the main logic for the Game of Life.
GameOfLife: Struct that holds the game board and its dimensions.
NewGameOfLife: Initializes a new game with random values.
countLiveNeighbourCell: Counts the live neighbors of a cell.
UpdateBoard: Updates the game board based on the rules of the Game of Life.
PrintBoard: Prints the current state of the game board.

 # Requirements
This program is written in Go and requires a Go compiler to run.


