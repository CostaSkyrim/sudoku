package bv

// function that checks if the board has duplicates
func IsValid(board *[9][9]int) bool {
	// check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{} // a counter variable to support the hasDuplicates function and hold a array stack for the rows
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++ // iterate and check for each row
		}
		if hasDuplicates(counter) { // if it has duplicates exit
			return false
		}
	}

	// check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{} // same but for the columns
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	// check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{} // same but for the 3x3 square
			for row := i; row < i+3; row++ { // iterate for each of the 9 3x3 squares
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

// function that helps to check if there are duplicates
func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 { // if no duplicates continue
			continue
		}
		if count > 1 { // if the same number appears more than once exit
			return true
		}
	}
	return false
}
