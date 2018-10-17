package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

const errorMessage = "Incorrect Input. You should input one integer more than 1.\n Please, try again."

func getInt(data string) (int, error) {
	number, err := strconv.ParseUint(data, 10, 64)
	return int(number), err
}

func sequence(n int) (sqrSequence []int) {
	for i := 1; int(math.Pow(float64(i), 2)) < n; i++ {
		sqrSequence = append(sqrSequence, i)
	}
	return
}

func prntSeq(elements []int) (str string) {
	for _, elem := range elements {
		str += strconv.Itoa(elem) + ", "
	}
	return strings.Trim(str, ", ")
}

func main() {
	data := os.Args[1:]
	if len(data) != 1 {
		println(errorMessage)
		return
	}

	n, err := getInt(data[0])
	if err != nil {
		println(errorMessage)
		return
	}
	if n < 2 {
		println(errorMessage)
		return
	}
	println(prntSeq(sequence(n)))

}
