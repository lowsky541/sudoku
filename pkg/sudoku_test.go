package sudoku_test

import (
	sudoku "sudoku/pkg"
	"testing"
)

type test struct {
	input       []string
	expectedErr error
	expectedSol sudoku.Board
}

var tests = []test{

	// Good

	{
		input: []string{
			".96.4...1",
			"1...6...4",
			"5.481.39.",
			"..795..43",
			".3..8....",
			"4.5.23.18",
			".1.63..59",
			".59.7.83.",
			"..359...7",
		},
		expectedSol: sudoku.Board{
			{3, 9, 6, 2, 4, 5, 7, 8, 1},
			{1, 7, 8, 3, 6, 9, 5, 2, 4},
			{5, 2, 4, 8, 1, 7, 3, 9, 6},
			{2, 8, 7, 9, 5, 1, 6, 4, 3},
			{9, 3, 1, 4, 8, 6, 2, 7, 5},
			{4, 6, 5, 7, 2, 3, 9, 1, 8},
			{7, 1, 2, 6, 3, 8, 4, 5, 9},
			{6, 5, 9, 1, 7, 4, 8, 3, 2},
			{8, 4, 3, 5, 9, 2, 1, 6, 7},
		},
	},
	{
		input: []string{
			"1.58.2...",
			".9..764.5",
			"2..4..819",
			".19..73.6",
			"762.83.9.",
			"....61.5.",
			"..76...3.",
			"43..2.5.1",
			"6..3.89..",
		},
		expectedSol: sudoku.Board{
			{1, 4, 5, 8, 9, 2, 6, 7, 3},
			{8, 9, 3, 1, 7, 6, 4, 2, 5},
			{2, 7, 6, 4, 3, 5, 8, 1, 9},
			{5, 1, 9, 2, 4, 7, 3, 8, 6},
			{7, 6, 2, 5, 8, 3, 1, 9, 4},
			{3, 8, 4, 9, 6, 1, 7, 5, 2},
			{9, 5, 7, 6, 1, 4, 2, 3, 8},
			{4, 3, 8, 7, 2, 9, 5, 6, 1},
			{6, 2, 1, 3, 5, 8, 9, 4, 7},
		},
	},
	{
		input: []string{
			"..5.3..81",
			"9.285..6.",
			"6....4.5.",
			"..74.283.",
			"34976...5",
			"..83..49.",
			"15..87..2",
			".9....6..",
			".26.495.3",
		},
		expectedSol: sudoku.Board{
			{4, 7, 5, 9, 3, 6, 2, 8, 1},
			{9, 3, 2, 8, 5, 1, 7, 6, 4},
			{6, 8, 1, 2, 7, 4, 3, 5, 9},
			{5, 1, 7, 4, 9, 2, 8, 3, 6},
			{3, 4, 9, 7, 6, 8, 1, 2, 5},
			{2, 6, 8, 3, 1, 5, 4, 9, 7},
			{1, 5, 3, 6, 8, 7, 9, 4, 2},
			{7, 9, 4, 5, 2, 3, 6, 1, 8},
			{8, 2, 6, 1, 4, 9, 5, 7, 3},
		},
	},

	// Wrong

	{
		// Only 8 cells for first row
		input: []string{
			".932..8.",
			"27.3.85..",
			".8.73.254",
			"9758...31",
			"....74.6.",
			"6.45.38.7",
			"7....2.48",
			"32.4...7.",
			"..8.579..",
		},
		expectedErr: sudoku.ErrMismatchCells,
	},
	{
		// Dups 2 at B2:B5
		input: []string{
			".867.2..4",
			".2.5..8..",
			"154.9.237",
			".7.9.5..1",
			".29..4.18",
			"51.6...42",
			"2.5.7..83",
			"...153...",
			"39...8.75",
		},
		expectedErr: sudoku.ErrInvalidBoard,
	},
	{
		// Dups 9 at B9:I9:I6
		input: []string{
			".7....28.",
			".2...6.57",
			"8654729..",
			"..925..64",
			".4..19.7.",
			"7.8..4..9",
			"3..7..698",
			"..79.1...",
			"59..28.39",
		},
		expectedErr: sudoku.ErrInvalidBoard,
	},
	{
		// Dups 2 at E7:H7:H9
		input: []string{
			"..213.748",
			"8.4.....2",
			".178.26..",
			".68.9.27.",
			".932....4",
			"5..46.3..",
			"..9.24.23",
			"..63..19.",
			"385..1.2.",
		},
		expectedErr: sudoku.ErrInvalidBoard,
	},
	{
		// Dups 2 at B6:B8:C9
		input: []string{
			"9.46.3..1",
			"37.1..2.6",
			"..6..93.4",
			"..13..9.5",
			"56..91...",
			"82...461.",
			"..79...4.",
			"425.167..",
			"1.2..75.8",
		},
		expectedErr: sudoku.ErrInvalidBoard,
	},
	{
		// Dups 2 at B6:C9
		input: []string{
			"9.46.3..1",
			"37.1..2.6",
			"..6..93.4",
			"..13..9.5",
			"56..91...",
			"8....461.",
			"..79...4.",
			"425.167..",
			"1.2..75.8",
		},
		expectedErr: sudoku.ErrInvalidBoard,
	},
	{
		// Not ever near from a valid sudoku
		input: []string{
			"not",
			"a",
			"sudoku",
		},
		expectedErr: sudoku.ErrMismatchRows,
	},
	{
		// Only 8 rows
		input: []string{
			"53..8294.",
			"8..34...5",
			"3542761..",
			"..6.3...4",
			"9....162.",
			".9...3.78",
			"7438.9...",
			"..5..43.1",
		},
		expectedErr: sudoku.ErrMismatchRows,
	},
	{
		// Multiple solutions
		input: []string{
			"295743861",
			"4318659..",
			"876192543",
			"387459216",
			"612387495",
			"549216738",
			"763524189",
			"928671354",
			"1549386..",
		},
		expectedErr: sudoku.ErrMultipleSolutions,
	},
}

func TestSudoku(t *testing.T) {
	for ti, test := range tests {
		sol, err := sudoku.ParseAndSolve(test.input)
		board := sol.Board

		if test.expectedErr != nil && err != test.expectedErr {
			t.Errorf("test %d: got %v, expected %v", ti+1, err, test.expectedErr)
		}

		if test.expectedErr == nil && !sudoku.IsMatch(board, test.expectedSol) {
			t.Errorf("test %d: got %+v, expected %+v", ti+1, board, test.expectedSol)
		}
	}
}
