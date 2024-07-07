# Solving sudokus with backtracking

## Using the program

The program let's you input a sudoku via the command line by entering each row of the sudoku as separate arguments. Dots are used as placeholders for empties.

### Building

The program has been written in Go-be sure to have it installed on your system.

Input the follwing commands in your terminal to build the program in a newly created *./build* directory.

```sh
make .
```

## What's backtracking ?

![alt text](https://camo.githubusercontent.com/cfa2fdd940fc2570245e9ed9a6323567978253384697b1293e40f12d6985cfc6/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f382f38632f5375646f6b755f736f6c7665645f62795f626163747261636b696e672e676966 "Visualization of backtracking")

"Backtracking is a general algorithm for finding solutions to some computational problems, notably constraint satisfaction problems, that incrementally builds candidates to the solutions, and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be completed to a valid solution."

In the case of a sudoku solver, the solver tries to find the first empty cell and fills it with a valid digit that doesn't already occurs in the corresponding row, column or 3x3 section. In the following step we move to the next empty cell and insert another valid digit and so on. If we get stuck and tried all possible values for the current cell then we move back to the previous one (which is called as backtracking). Now we try there our luck with the next valid digit in this cell and move on. The board is finally solved if the program was able to fill all cells with a valid digit.

## Documentation for developers

### File structure

**main.go**: Entry point and logics of the program.  
**parser.go**: Command-line reading functions.  
**display.go**: Board display related functions.  
**config.go**: Configuration of the program--be sure to rebuild after changements.