package sudoku

type digitSet struct {
	s [9]bool
}

func (ds *digitSet) add(d uint8) bool {
	v := ds.s[d-1]

	if !v {
		ds.s[d-1] = true
		return true
	}

	return false
}

// Check if a board is valid according to sudoku rules.
func IsValid(board Board) bool {
	for i := 0; i < 9; i++ {

		// Check each rows
		digits := digitSet{}
		for row := 0; row < 9; row++ {
			n := board[row][i]

			if n != 0 && !digits.add(n) {
				return false
			}
		}

		// Check each columns
		digits = digitSet{}
		for col := 0; col < 9; col++ {
			n := board[i][col]

			if n != 0 && !digits.add(n) {
				return false
			}
		}

		// Check each blocks
		digits = digitSet{}
		bx, by := (i%3)*3, (i/3)*3

		for cell := 0; cell < 9; cell++ {
			cx := bx + cell/3
			cy := by + cell%3
			n := board[cy][cx]

			if n != 0 && !digits.add(n) {
				return false
			}
		}

	}

	return true
}
