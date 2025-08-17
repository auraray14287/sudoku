package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	var board [9][9]int
	// Kefa code block
	// This is to create a 2D array with 9 colmn and 9 rows
	// Loop over each command-line argument starting from index 1 after the program name
	for i, row := range os.Args[1:] {
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}
		// checks through row and converts (.) to (0). In Go all int starts at 0
		for j, c := range row {
			if c == '.' {
				board[i][j] = 0
				// if betwen (1-9) it's then being converted to rune then int through the ASCII method in GO.
			} else if c >= '1' && c <= '9' {
				board[i][j] = int(c - '0')
			} else {
				fmt.Println("Error") // if the character is not a number between 1 and 9
				return
			}
		}
	}

	// Raymond's code block
	// Check if 'board' is valid
	if !isBoardValid(board) {
		fmt.Println("Error")
		return
	}

	var originalBoard [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			originalBoard[i][j] = board[i][j]
		}
	}

	// Albanus code block
	// Backtracking solver
	// Gets hold of the board and checks the validity using sudoku rules.
	var solve func() bool
	solve = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == 0 {
					for num := 1; num <= 9; num++ {
						valid := true
						// Check row and column - checks each row against all cols.
						for k := 0; k < 9; k++ {
							if board[i][k] == num || board[k][j] == num {
								valid = false
								break
							}
						}
						// Check 3x3 box
						if valid {
							startRow := (i / 3) * 3
							startCol := (j / 3) * 3
							for r := 0; r < 3; r++ {
								for c := 0; c < 3; c++ {
									if board[startRow+r][startCol+c] == num {
										valid = false
										break
									}
								}
								if !valid {
									break
								}
							}
						}
						if valid {
							board[i][j] = num
							if solve() {
								return true
							}
							board[i][j] = 0
						}
					}
					return false
				}
			}
		}
		return true
	}

	if !solve() {
		fmt.Println("Error")
		return
	}

	// Raymond's code block
	// Check for multiple solutions
	if hasMultipleSolutions(originalBoard) {
		fmt.Println("Error")
		return
	}

	// Raymond's code block
	// Print the solution
	printBoard(board)
}

// Raymond's code block
// Check if 'board' is valid
func isValid(board [9][9]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	// Check 3x3 box
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func isBoardValid(board [9][9]int) bool {
	// We start with rows
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}

	}

	// Columns
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != 0 {
				if seen[board[i][j]] {
					return false // duplicate found
				}
				seen[board[i][j]] = true
			}
		}
	}

	// Check 3x3 boxes
	for box := 0; box < 9; box++ {
		seen := make(map[int]bool)
		startRow, startCol := (box/3)*3, (box%3)*3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[startRow+i][startCol+j] != 0 {
					if seen[board[startRow+i][startCol+j]] {
						return false
					}
					seen[board[startRow+i][startCol+j]] = true
				}
			}
		}
	}
	return true
}

func hasMultipleSolutions(board [9][9]int) bool {
	var count int
	var tempBoard [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			tempBoard[i][j] = board[i][j]
		}
	}

	var solveForMultiple func() bool
	solveForMultiple = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if tempBoard[i][j] == 0 {
					for num := 1; num <= 9; num++ {
						if isValid(tempBoard, i, j, num) {
							tempBoard[i][j] = num
							if solveForMultiple() {
								if count > 1 {
									return true
								}
							}
							tempBoard[i][j] = 0
						}
					}
					return false
				}
			}
		}
		count++
		return true
	}

	solveForMultiple()
	return count > 1
}

func printBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
