package sudoku

import "fmt"

// Display a board.
func Display(board Board) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Printf("%c", board[y][x]+'0')
			if x != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Display a board with the cell, located at row and col, with an highlight red color.
func DisplayHighlight(board Board, row int, col int) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if y == row && x == col {
				fmt.Printf("\u001b[41m%c\033[0m", board[y][x]+'0')
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c ", board[y][x]+'0')
			}
		}
		fmt.Println()
	}
}
