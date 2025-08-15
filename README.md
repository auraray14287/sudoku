Sudoku Solver (Go)

A command-line Sudoku solver written in Go.
It takes a 9x9 Sudoku puzzle as 9 separate string arguments and prints the solved puzzle in a specific output format.
Invalid inputs or unsolvable puzzles result in Error output.
How It Works

    Accepts exactly 9 arguments, each representing a row of the Sudoku board.
    Each row must be exactly 9 characters long.
    Digits 1-9 are pre-filled cells, . represents an empty cell.
    Validates that the starting grid has no duplicate digits in any row, column, or 3×3 subgrid.
    Uses backtracking to solve the puzzle.
    Prints the solved grid if successful, otherwise prints Error.

Usage
Run from the module root

go run . "<row1>" "<row2>" "<row3>" "<row4>" "<row5>" "<row6>" "<row7>" "<row8>" "<row9>"
