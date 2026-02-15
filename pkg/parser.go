package sudoku

// Parse tries to converts a list of string into a table of 9x9 integers.
//
// Iterate over each string of `args` then map each characters into digits
// before placing them in a 9x9 array.
func Parse(args []string) (Board, error) {
	var out Board

	if len(args) != 9 {
		return out, ErrMismatchRows
	}

	for ai, arg := range args {
		if len(arg) != 9 {
			out = Board{}
			return out, ErrMismatchCells
		}

		for ci, char := range arg {
			digit := int(char - '0')

			if char == '.' {
				out[ai][ci] = 0
			} else if digit >= 0 && digit <= 9 {
				out[ai][ci] = uint8(digit)
			} else {
				out = Board{}
				return out, ErrInvalidCell
			}
		}
	}

	if !IsValid(out) {
		out = Board{}
		return out, ErrInvalidBoard
	}

	return out, nil
}
