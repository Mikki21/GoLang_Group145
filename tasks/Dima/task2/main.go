package main

import (
	"fmt"
	"strconv"
)

// getFloat function gets input vlaue till you will print correct float, which will return
func getFloat() (inVar float64, err error) {
	isOk := false
	for !isOk {
		var input string
		fmt.Scan(&input)

		inVar, err = strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Printf("%v \nPlease try again: ", err)
		} else if inVar <= 0 {
			fmt.Print("Please enter positive number: ")
		} else {
			isOk = true
		}

	}
	return
}

// isPossible function checkes if It`s possible to put a envelope in the other
// using diagonals
func isPossible(a, b, c, d float64) (poss bool) {
	if a*a+b*b > c*c+d*d {
		poss = true
	} else if a*a+b*b < c*c+d*d {
		poss = true
	} else {
		poss = false
	}
	return
}

func main() {
	again := true

	for again {
		fmt.Print("\nEnter width of the first envelope: ")
		firstWidth, _ := getFloat()

		fmt.Print("Enter height of the first envelope: ")
		firstHeight, _ := getFloat()

		fmt.Print("Enter width of the second envelope: ")
		secondWidth, _ := getFloat()

		fmt.Print("Enter height of the second envelope: ")
		secondHeight, _ := getFloat()

		if isPossible(firstWidth, firstHeight, secondWidth, secondHeight) {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}

		var cont string
		fmt.Print("Would you like to continue? Press y/yes to continue or ")
		fmt.Scan(&cont)

		switch cont {
		case "y":
			fmt.Print("Enter new values: ")
		case "Y":
			fmt.Print("Enter new values: ")
		case "yes":
			fmt.Print("Enter new values: ")
		case "Yes":
			fmt.Print("Enter new values: ")
		default:
			again = false
		}

	}
}
