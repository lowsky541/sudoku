package sudoku

// Deeply copy a board returning its copy.
func deepCopy(board Board) (out Board) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			out[i][j] = board[i][j]
		}
	}
	return out
}

// Checks if two boards are equal.
func IsMatch(a Board, b Board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
