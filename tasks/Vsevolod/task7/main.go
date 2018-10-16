package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func errorMessage() {
	fmt.Print("Incorrect Input. You should input one integer more than 1.\n Please, try again.")
}

func getNumber(osArgs []string) (int, bool) {
	if len(osArgs) == 2 {
		number, err := strconv.ParseInt(osArgs[1], 10, 64)
		return int(number), err == nil
	}
	return 0, false
}

func sequence(n int) (sqrSequence []int) {
	for i := 1; int(math.Pow(float64(i), 2)) < n; i++ {
		sqrSequence = append(sqrSequence, i)
	}
	return
}

func printSequence(sqrSequence []int) {
	for i := 0; i < len(sqrSequence)-1; i++ {
		fmt.Printf("%v, ", sqrSequence[i])
	}
	fmt.Printf("%v", sqrSequence[len(sqrSequence)-1])
}

func main() {
	n, isInt := getNumber(os.Args)
	if n > 1 && isInt {
		printSequence(sequence(n))
	} else {
		errorMessage()
	}

}
