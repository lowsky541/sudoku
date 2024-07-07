package main

import "fmt"

func DisplayBoard(board [9][9]int) {
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

func DisplayBoardHighlight(board [9][9]int, row int, col int) {
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
