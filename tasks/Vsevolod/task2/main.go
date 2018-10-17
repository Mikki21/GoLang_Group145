package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const message = "Incorrect input. It must be one non-Negative number."
const msgAnsw = "Would you like to continue? If Yes, press y/yes."

func scan() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	err := in.Err()
	if err != nil {
		return "", err
	}
	return in.Text(), err
}

func getFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func getAnswer(str string) bool {
	str = strings.ToLower(str)
	return str == "y" || str == "yes"
}

/*
func game() {
	a, _ := getSide('a', scan)
	b, _ := getSide('b', scan)
	c, _ := getSide('c', scan)
	d, _ := getSide('d', scan)
	if a < b {
		a, b = b, a
	}
	if c < d {
		c, d = d, c
	}

	if c < a && d < b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")

	}
}
*/

func getSide(scan func() (string, error)) (float64, error) {
	aStr, err := scan()
	if err != nil {
		return -1, errors.New(message)
	}
	side, err := getFloat(aStr)
	if err != nil {
		return -1, errors.New(message)
	}
	if side < 0 {
		return side, errors.New(message)
	}
	return side, err
}

func main() {
	answer := true
	for answer {
		var side [4]float64
		for i := 0; i < 4; i++ {
			errGetSide := errors.New("init")
			for errGetSide != nil {
				fmt.Println("Enter side " + string(i+97) + ": ")
				side[i], errGetSide = getSide(scan)
				if errGetSide != nil {
					println(message)
				}
			}
		}
		a, b, c, d := side[0], side[1], side[2], side[3]

		if a < b {
			a, b = b, a
		}
		if c < d {
			c, d = d, c
		}

		if c <= a && d <= b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")

		}
		println(msgAnsw)
		answStr, _ := scan()
		answer = getAnswer(answStr)
	}

}

/*answer := true
for answer {
	game()
	answer = getAnswer()
}
*/
