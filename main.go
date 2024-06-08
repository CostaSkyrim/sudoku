package main

import (
	"fmt"
	"os"
	pb "sudoku/printboard"
	vp "sudoku/validateparams"
)

func main() {
	p := os.Args[1:]
	validata, err := vp.ValidateParams(p)

	if err != nil {
		fmt.Println(err)
	} else {
		pb.PrintBoard(validata)
	}

}
