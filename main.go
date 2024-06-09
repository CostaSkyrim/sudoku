package main

import (
	"fmt"
	"os"

	pb "sudoku/printboard"
	sv "sudoku/solver"
	vp "sudoku/validateparams"
)

func main() {
	p := os.Args[1:]
	validata, err := vp.ValidateParams(p)

	if err != nil {
		fmt.Println(err)
	} else {
		if sv.SolveSudoku(&validata) {
			pb.PrintBoard(validata)
		} else {
			fmt.Println("Error")
		}
	}
}
