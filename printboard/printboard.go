package pb

import (
	"fmt"
	"strconv"
)

// function that prints the board in the output
func PrintBoard(b [9][9]int) {
	var p string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j == 0 { // if it's the first column don't put space on edge
				p = p + strconv.Itoa(b[i][j])
			} else {
				p = p + " " + strconv.Itoa(b[i][j])
			}
		}
		fmt.Println(p)
		p = "" // resets it after each iteration so you don't output brackets
	}
}
