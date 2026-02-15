package main

import (
	"fmt"
	"os"
	sudoku "sudoku/pkg"
)

func main() {
	args := os.Args[1:]

	board, err := sudoku.Parse(args)
	if err != nil {
		fmt.Printf("Error: could not parse board: %s", err.Error())
		return
	}

	solution, err := sudoku.Solve(board)
	if err != nil {
		fmt.Println("Error: This sudoku has multiple solutions.")
		return
	}

	fmt.Printf("This sudoku has been solved in %v.\n", solution.Duration)
	sudoku.Display(solution.Board)
}
