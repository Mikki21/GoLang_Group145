package main

import (
	"fmt"
	"os"
	"strconv"
)

func rule() {
	fmt.Print("When you are running the program,\n",
		"you should enter height and width\n",
		"of the check board via space,\n",
		"after calling program in command line.\n",
		"Example: go run main.go 4 5\n",
		"Try again.")
}

func incorrectInput() {
	fmt.Print("Incorrect input\n")
	rule()
}
func chessBoard(arguments []string) {
	height, errh := strconv.ParseInt(arguments[0], 10, 64)
	width, errw := strconv.ParseInt(arguments[1], 10, 64)
	if errh == nil && errw == nil {
		if height > 0 && width > 0 {
			for i := 0; i < int(height); i++ {
				if i%2 == 0 {
					for j := 0; j < int(width); j++ {
						fmt.Print("* ")
					}
				} else {
					for j := 0; j < int(width); j++ {
						fmt.Print(" *")
					}
				}
				fmt.Print("\n")
			}
		} else {
			incorrectInput()
		}
	} else {
		incorrectInput()
	}
}

func main() {
	arguments := os.Args
	switch len(arguments) {
	case 1:
		rule()
	case 3:
		chessBoard(arguments[1:])
	default:
		incorrectInput()
	}

}
