package sudoku

import "errors"

// I want to use constants so badly.....

var (
	ErrMultipleSolutions = errors.New("multiple solutions")
	ErrMismatchRows      = errors.New("board has not enough/too much rows")
	ErrMismatchCells     = errors.New("board has not enough/too much cells")
	ErrInvalidCell       = errors.New("board has one or more invalid cells")
	ErrInvalidBoard      = errors.New("board is invalid")
)
