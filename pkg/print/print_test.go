package main

import (
	"bytes"
	"gameoflife/pkg/board"
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

	// Capture the output of PrintBoard
	var buf bytes.Buffer
	old := stdout
	stdout = &buf
	defer func() { stdout = old }()

	PrintBoard(game)

	expected := "1 0 1 \n0 1 0 \n1 0 1 \n"
	if buf.String() != expected {
		t.Errorf("PrintBoard() output = %v, want %v", buf.String(), expected)
	}
}

// Helper to capture stdout
var stdout = &bytes.Buffer{}

func init() {
	stdout = &bytes.Buffer{}
}
