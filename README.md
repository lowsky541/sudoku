# sudoku

This repository contains the code for a `sudoku` solver written in Go using a backtracking approach for solving.

## Using the program

The program let you input a sudoku via the command line by entering each row of the sudoku as separate arguments. Dot characters are used as placeholders for fillable cells.

```console
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
This sudoku has been solved in 138.789µs.
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

## What's backtracking and how does it work ?

![Visualization of backtracking](/assets/backtracking.gif)

"Backtracking is a class of algorithms for finding solutions to some computational problems, notably constraint satisfaction or enumeration problems, that incrementally builds candidates to the solutions, and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be completed to a valid solution." -- from the ["Backtracking" article on Wikipedia](https://en.wikipedia.org/wiki/Backtracking).

In the case of a sudoku solver, the solver tries to find the first empty cell and fills it with a valid digit that doesn't already occurs in the corresponding row, column or 3x3 section. In the following step we move to the next empty cell and insert another valid digit and so on. If we get stuck and tried all possible values for the current cell then we move back to the previous one (which is called backtracking). Now we try there our luck with the next valid digit in this cell and move on. The board is finally solved if the program was able to fill all cells with a valid digit.
