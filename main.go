package main

import (
	"fmt"
	"os"
	"time"
)

func IsNumberPlaceable(puzzle [9][9]int, row int, col int, num int) bool {

	// Check if the number defined by "num" is placeable on this `row` and `column`

	// Calculate the position of the block
	//  1. Calculate the position in blocks
	//  2. Calculate/convert the position in cells
	bx := (col / 3) * 3
	by := (row / 3) * 3

	// For each target cells in the row, the column and the block
	for i := 0; i < 9; i++ {

		// In iteration over each numbers in the row the cell's a member of (axis-x/horizontal)
		// If the number we search for (num) is found then the sudoku constraint has failed
		if puzzle[row][i] == num {
			return false
		}

		// In iteration over each numbers in the column the cell's a member of (axis-y/vertical)
		// If the number we search for (num) is found then the sudoku constraint has failed
		if puzzle[i][col] == num {
			return false
		}

		// Calculate the position of the next neighbour cell in the block
		cx := bx + i/3
		cy := by + i%3

		// In iteration over each numbers in the block the cell's a member of (3x3 section)
		// If the number we search for (num) is found then the sudoku constraint has failed
		if puzzle[cy][cx] == num {
			return false
		}
	}

	return true
}

func GetPossibilities(puzzle [9][9]int, row int, col int) []int {
	// Get a list of all the numbers possible

	var out []int

	for n := 1; n <= 9; n++ {
		if IsNumberPlaceable(puzzle, row, col, n) {
			out = append(out, n)
		}
	}

	return out
}

func FindNextZero(puzzle [9][9]int) (int, int) {
	// Find the next empty on the board

	for ri := 0; ri < 9; ri++ {
		for ci := 0; ci < 9; ci++ {
			if puzzle[ri][ci] == 0 {
				return ri, ci
			}
		}
	}

	// The board has no more empties
	return -1, -1
}

func SolveSudoku(puzzle *[9][9]int) bool {

	row, col := FindNextZero(*puzzle)
	if row == -1 && col == -1 {
		return true
	}

	possibilities := GetPossibilities(*puzzle, row, col)
	for _, possibility := range possibilities {
		puzzle[row][col] = possibility

		// -- config -- Visualize the backtracking step-by-step
		if BacktrackingVisualize {
			DisplayBoardHighlight(*puzzle, row, col)
			time.Sleep(BacktrackingStepTime)
			fmt.Print("\033[H\033[2J")
		}

		if SolveSudoku(puzzle) {
			return true
		}

		// Undo placement
		puzzle[row][col] = 0
	}

	return false
}

func main() {

	// Parse the command line arguments into a 2D array of integers
	args := os.Args[1:]
	puzzle, err := ParseBoard9x9(args)

	if err != nil {
		fmt.Println("Error")
		return
	}

	// -- config -- Show sudoku solving execution time
	var start time.Time
	if ShowSolveTime {
		start = time.Now()
	}

	if SolveSudoku(&puzzle) {

		// -- config -- Show sudoku solving execution time
		if ShowSolveTime {
			duration := time.Since(start)
			fmt.Printf("This sudoku has been solved in %v.\n", duration)
		}

		DisplayBoard(puzzle)
	} else {
		//fmt.Printf("This sudoku has no solutions\n")
		fmt.Println("Error")
		return
	}
}
