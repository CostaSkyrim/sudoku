package main

import (
	"os"
	vp "sudoku/validateparams"
)

func main() {
	p := os.Args[1:]
	vp.ValidateParams(p)
}
