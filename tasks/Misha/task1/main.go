package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	oddString  string = "* "
	evenString string = " *"
	instruct   string = "Program prints ChessDesk according to the entered parametrs '(go run .\\main.go <height> <width>)' "
)

func parseParametrs(param []string) (int, int, error) {
	h, errH := strconv.ParseUint(param[1], 10, 64)
	if errH != nil {
		return 0, 0, errH
	}
	w, errW := strconv.ParseUint(param[2], 10, 64)
	if errW != nil {
		return 0, 0, errW
	}
	return int(h), int(w), nil
}

func createDesk(h, w int) (res string, err error) {
	if h < 1 || w < 1 {
		err = fmt.Errorf("invalid values")
		return
	}
	for x := 1; x < h+1; x++ {
		for y := 1; y < w+1; y++ {
			if x%2 == 0 {
				fmt.Print(evenString)
				res += evenString
			} else {
				fmt.Print(oddString)
				res += oddString
			}
		}
		fmt.Println()
		res += "\n"
	}
	return
}

func main() {
	console := os.Args
	switch len(console) {
	case 1:
		fmt.Println(instruct)
	case 3:
		height, width, err := parseParametrs(console)
		if err != nil {
			fmt.Println("Invalid paramaetrs!")
		} else {
			tempStr, err := createDesk(height, width)
			if err != nil {
				fmt.Println("Invalid paramaetrs!")
			}
			fmt.Print(tempStr)
		}
	default:
		fmt.Println("Only two parametrs are available to enter !")
	}
}
