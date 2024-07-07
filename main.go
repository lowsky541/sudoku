package main

import (
	"fmt"
	"os"
	"time"
)

func IsNumberPlaceable(puzzle [9][9]int, row int, col int, num int) bool {
	// Check if the number `num` is present in the sub-sudoku, the row or the column
	// of the cell at `row` and `col`.
	// We need to give this function the sudoku as it currently look.

	// Calculate the position of the block
	//  1. Calculate the position in blocks
	//  2. Calculate/convert the position in cells
	bx := (col / 3) * 3
	by := (row / 3) * 3

	// For each places (empty/non-empty) in the row, the column and the block
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
	// Get a list of all the possible numbers (returns `out`) that could be placed
	// on this cell that DOES NOT breaks any sudoku rules.

	// The possible numbers that could be placed on (row, col)
	var out []int

	// For each numbers from 1 to 9, check if the number could be placed here.
	for n := 1; n <= 9; n++ {
		if IsNumberPlaceable(puzzle, row, col, n) {
			// The number could be placed here; consider it a possibility
			// (append it to `out`) and continue the loop.
			out = append(out, n)
		}
	}

	return out
}

func FindNextZero(puzzle [9][9]int) (int, int) {
	// Find the next empty place on the board.
	// Iterates over all the places (empty/non-empty)
	// of the board--for each rows then each columns.

	// ri = row index
	for ri := 0; ri < 9; ri++ {
		// ci = column index
		for ci := 0; ci < 9; ci++ {
			// The element at place (ri, ci)
			if puzzle[ri][ci] == 0 {
				return ri, ci
			}
		}
	}

	// The board has no more empties--signal it to the caller
	return -1, -1
}

func FindNextZeroReversed(puzzle [9][9]int) (int, int) {
	for ri := 8; ri >= 0; ri-- {
		for ci := 8; ci >= 0; ci-- {
			if puzzle[ri][ci] == 0 {
				return ri, ci
			}
		}
	}

	return -1, -1
}

func SolveSudoku(puzzle *[9][9]int, reversed bool) bool {
	row, col := 0, 0
	if reversed {
		row, col = FindNextZeroReversed(*puzzle)
	} else {
		row, col = FindNextZero(*puzzle)
	}

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

		if SolveSudoku(puzzle, reversed) {
			return true
		}

		// Undo placement
		puzzle[row][col] = 0
	}

	return false
}

func Duplicate(array [9][9]int) [9][9]int {
	var out [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			out[i][j] = array[i][j]
		}
	}
	return out
}

func Mirrored(a [9][9]int, b [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// Parse the command line arguments into a 2D array of integers
	args := os.Args[1:]
	puzzle, err := ParseBoard9x9(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	puzzlev1 := Duplicate(puzzle)
	puzzlev2 := Duplicate(puzzle)

	start := time.Now()

	solvev1, solvev2 := SolveSudoku(&puzzlev1, false), SolveSudoku(&puzzlev2, true)

	if solvev1 && solvev2 && Mirrored(puzzlev1, puzzlev2) {
		duration := time.Since(start)
		fmt.Printf("This sudoku has been solved in %v.\n", duration)
		DisplayBoard(puzzlev1)
	} else {
		fmt.Printf("This sudoku has no solutions.\n")
		fmt.Printf("Check your inputs, damn it!\n")
		return
	}
}
