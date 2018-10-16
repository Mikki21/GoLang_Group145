package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	odd  string = "* "
	even string = " *"
)

func ChessDesk(args []string) {
	h, errH := strconv.ParseUint(args[0], 10, 64)
	w, errW := strconv.ParseUint(args[1], 10, 64)
	fmt.Println(w)
	if errH == nil && errW == nil {
		for iH := 1; iH < int(h+1); iH++ {
			for iW := 1; iW < int(w+1); iW++ {
				if iH%2 == 0 {
					fmt.Print(even)
				} else {
					fmt.Print(odd)
				}

			}
			fmt.Println("")
		}
	}
	if h == 0 || w == 0 {
		fmt.Println("Enter positive Values")
	}
}

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		fmt.Println("Program prints ChessDesk according to the entered parametrs '(go run .\\main.go <height> <width>)' ")
	case 3:
		ChessDesk(args[1:])
	default:
		fmt.Println("Only two parametrs are available to enter!")
	}
}
