package main

import (
	"fmt"
	"os"
	"strconv"
)

func chessField(args []string) {

	height, incorHeight := strconv.ParseInt(args[0], 10, 64)
	width, incorWidth := strconv.ParseInt(args[1], 10, 64)

	if incorHeight == nil && incorWidth == nil {
		if width > 0 && height > 0 {
			for i := 0; i < int(height); i++ {
				for j := 0; j < int(width); j++ {
					if i%2 == 0 {
						fmt.Print("* ")
					} else {
						fmt.Print(" *")
					}
				}
				fmt.Print("\n")
			}
		} else {
			fmt.Print("Incorrect values")
		}
	} else {
		fmt.Print("Incorrect values")
	}
}

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		fmt.Print("Please, enter height and width after caling the program \n(go run .\\main.go <height> <width>)")
	case 2 + 1:
		chessField(args[1:])

	default:
		fmt.Print("Incorrect number of values")
	}
}
