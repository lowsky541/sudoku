package sudoku

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type (
	Board    [9][9]uint8
	Solution struct {
		Board    Board
		Duration time.Duration
	}
)

var isTesting = false

func init() {
	if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		isTesting = true
	}
}

func isNumberPlaceable(board Board, row int, col int, num uint8) bool {
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
		if board[row][i] == num {
			return false
		}

		// In iteration over each numbers in the column the cell's a member of (axis-y/vertical)
		// If the number we search for (num) is found then the sudoku constraint has failed
		if board[i][col] == num {
			return false
		}

		// Calculate the position of the next neighbour cell in the block
		cx := bx + i/3
		cy := by + i%3

		// In iteration over each numbers in the block the cell's a member of (3x3 section)
		// If the number we search for (num) is found then the sudoku constraint has failed
		if board[cy][cx] == num {
			return false
		}
	}

	return true
}

func getPossibilities(board Board, row int, col int) []uint8 {
	// Get a list of all the possible numbers (returns `out`) that could be placed
	// on this cell that DOES NOT breaks any sudoku rules.

	// The possible numbers that could be placed on (row, col)
	var out []uint8

	// For each numbers from 1 to 9, check if the number could be placed here.
	for n := uint8(1); n <= 9; n++ {
		if isNumberPlaceable(board, row, col, n) {
			// The number could be placed here; consider it a possibility
			// (append it to `out`) and continue the loop.
			out = append(out, n)
		}
	}

	return out
}

func findNextZero(board Board) (int, int) {
	// Find the next empty place on the board.
	// Iterates over all the places (empty/non-empty)
	// of the board--for each rows then each columns.

	// ri = row index
	for ri := 0; ri < 9; ri++ {
		// ci = column index
		for ci := 0; ci < 9; ci++ {
			// The element at place (ri, ci)
			if board[ri][ci] == 0 {
				return ri, ci
			}
		}
	}

	// The board has no more empties--signal it to the caller
	return -1, -1
}

func findNextZeroReversed(board Board) (int, int) {
	for ri := 8; ri >= 0; ri-- {
		for ci := 8; ci >= 0; ci-- {
			if board[ri][ci] == 0 {
				return ri, ci
			}
		}
	}
	return -1, -1
}

func backtrack(board *Board, reversed bool) bool {
	row, col := 0, 0
	if reversed {
		row, col = findNextZeroReversed(*board)
	} else {
		row, col = findNextZero(*board)
	}

	if row == -1 && col == -1 {
		return true
	}

	possibilities := getPossibilities(*board, row, col)

	if reversed {
		slices.Reverse(possibilities)
	}

	for _, possibility := range possibilities {
		board[row][col] = possibility

		// -- config -- Visualize the backtracking step-by-step (disabled in testing)
		if !isTesting && BacktrackingVisualize {
			DisplayHighlight(*board, row, col)
			time.Sleep(BacktrackingStepTime)
			fmt.Print("\033[H\033[2J")
		}

		if backtrack(board, reversed) {
			return true
		}

		// Undo placement
		board[row][col] = 0
	}

	return false
}

func Solve(board Board) (Solution, error) {
	forward, backward := board, deepCopy(board)

	// Start solving the sudoku
	start := time.Now()
	isSolvedForward, isSolvedBackward := backtrack(&forward, false),
		backtrack(&backward, true)
	duration := time.Since(start)

	// Check if the forward and backward versions are solved and identical
	if isSolvedForward && isSolvedBackward && IsMatch(forward, backward) {
		return Solution{
			Board:    backward,
			Duration: duration,
		}, nil
	} else {
		return Solution{}, ErrMultipleSolutions
	}
}

func ParseAndSolve(input []string) (Solution, error) {
	board, err := Parse(input)
	if err != nil {
		return Solution{}, err
	}

	solution, err := Solve(board)
	if err != nil {
		return Solution{}, err
	}

	return solution, nil
}
