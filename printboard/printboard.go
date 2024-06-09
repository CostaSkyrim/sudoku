package pb

import (
	"fmt"
	"strconv"
)

func PrintBoard(b [9][9]int) {
	var p string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j == 0 {
				p = p + strconv.Itoa(b[i][j])
			} else {
				p = p + " " + strconv.Itoa(b[i][j])
			}
		}
		fmt.Println(p)
		p = ""
	}
}
