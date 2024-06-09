package main

import (
	"fmt"
	"os"

	pb "sudoku/printboard"
	sv "sudoku/solver"
	vp "sudoku/validateparams"
)

func main() {
	p := os.Args[1:] // variable that holds the arguments and passes it to ValidateParams function
	validata, err := vp.ValidateParams(p)

	if err != nil { // if there is any error from ValidateParams function print it
		fmt.Println(err)
	} else { // otherwise proceed to solve the sudoku and print it
		if sv.SolveSudoku(&validata) {
			pb.PrintBoard(validata)
		} else { // if it's not solvable
			fmt.Println("Error")
		}
	}
}
