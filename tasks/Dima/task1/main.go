package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseInput(dimentions []string) (height int64, width int64) {
	var err error
	height, err = strconv.ParseInt(dimentions[0], 10, 64)
	if height < 0 || err != nil {
		fmt.Printf("Incorrect input")
		return
	}
	width, err = strconv.ParseInt(dimentions[1], 10, 64)
	if width < 0 || err != nil {
		fmt.Printf("Incorrect input")
		return
	}
	return
}

func chessField(height, width int64) (output string) {

	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			if i%2 == 0 {
				output += "* "
			} else {
				output += " *"
			}
		}
		output += "\n"
	}
	return
}

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		fmt.Print("Please, enter height and width after caling the program \n(go run .\\main.go <height> <width>)")
	case 2 + 1:
		fmt.Printf("%s", chessField(parseInput(args[1:])))

	default:
		fmt.Print("Incorrect number of values")
	}
}
