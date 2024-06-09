package sv

import bv "sudoku/isvalid"

// function that solves the sudoku
func SolveSudoku(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 { // if the value is 0 in this position proceed
				for candidate := 9; candidate >= 1; candidate-- { // test if it's solvable with a specified candidate value
					board[i][j] = candidate
					 // assign the candidate value and run the sudoku assuming it is correct
					if bv.IsValid(board) {
						if SolveSudoku(board) { // backtrack on the function and execute it again
							return true
						}
						board[i][j] = 0 // reset the value if it's not solvable
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

// function that checks if the value is 0
func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
