package main

import (
	"errors"
)

func ParseBoard9x9(args []string) ([9][9]int, error) {

	var out [9][9]int

	if len(args) != 9 || len(args) == 0 {
		return out, errors.New("ParseBoard: too much or not enough arguments to create the board")
	}

	for ai, arg := range args {

		if len(arg) != 9 {
			out = [9][9]int{}
			return out, errors.New("ParseBoard: row has less or more characters than needed")
		}

		for ci, char := range arg {
			var digit int = int(char - '0')
			if char == '.' {
				out[ai][ci] = 0
			} else if digit >= 0 && digit <= 9 {
				out[ai][ci] = digit
			} else {
				out = [9][9]int{}
				return out, errors.New("ParseBoard: invalid character encountered while parsing the board")
			}
		}
	}

	return out, nil
}
