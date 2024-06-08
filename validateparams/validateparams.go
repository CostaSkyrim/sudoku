package vp

import "fmt"

func ValidateParams(s []string) {
	var boardr []rune
	for i := 0; i < len(s); i++ {
		for _, n := range s[i] {
			if n == '.' {
				n = '0'
			}
		}
	}
	fmt.Println(string(boardr))
}
