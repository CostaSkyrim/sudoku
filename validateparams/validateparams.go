package vp

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateParams(s []string) ([9][9]int, error) {
	arr := [9][9]int{}

	if len(s) != 9 {
		return arr, errors.New("Error")
	}

	for i := 0; i < len(s); i++ {
		r := strings.ReplaceAll(s[i], ".", "0")
		s[i] = r
	}

	for i := 0; i < len(s); i++ {
		if len(s[i]) != 9 {
			return arr, errors.New("Error")
		}
		for _, r := range s[i] {
			if !isNumeric(r) {
				return arr, errors.New("Error")
			}
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num, err := strconv.Atoi(string(s[row][col]))
			if err != nil {
				return arr, errors.New("Error: failed to convert character to integer")
			}
			arr[row][col] = num

		}
	}

	return arr, nil
}

func isNumeric(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}
