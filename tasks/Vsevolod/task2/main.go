package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Entering mistake", err)
	}
	return in.Text()
}

func getValue(letter rune) (side string) {
	fmt.Printf("Please, enter side %v: ", string(letter))
	return scan()
}

func getAnswer() bool {
	fmt.Print("Dear player, would you like to continue the game?\nIf you would like, press y/yes.\n")
	res := strings.ToLower(scan())
	return res == "y" || res == "yes"
}

func getFloat(item string) (float64, error) {
	return strconv.ParseFloat(item, 64)
}

func getSide(letter rune) (side float64, err error) {
	for {
		side, err = getFloat(getValue(letter))
		if err != nil {
			fmt.Print("Incorrect Input. It must be one number, like \"3.14\". Please, try again.\n")
			continue
		}
		if side <= 0 {
			fmt.Print("Incorrect Input. Side must be positive. Please, try again.\n")
			continue
		}
		break
	}
	return
}

func game() {
	a, _ := getSide('a')
	b, _ := getSide('b')
	c, _ := getSide('c')
	d, _ := getSide('d')
	if a < b {
		a, b = b, a
	}
	if c < d {
		c, d = d, c
	}

	if c < a && d < b {
		fmt.Print("Yes. You can put this one into the other.\n")
	} else {
		fmt.Print("No. You cannot put this one into the other.\n")

	}
}

func main() {
	answer := true
	for answer {
		game()
		answer = getAnswer()
	}
}
