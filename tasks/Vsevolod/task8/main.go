package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const magicNumber = 93
const errorMessage = "Incprrect parameters. You should input 2 different\nintegers from 0 to 18446744073709551615.\n Try again"

func getInt(data string) (uint64, error) {
	number, err := strconv.ParseUint(data, 10, 64)
	return number, err
}

func fibonacci(n int) (fib uint64) {
	var fibBeforeBefore uint64
	var fibBefore uint64 = 1
	if n > 1 {
		for i := 2; i <= n; i++ {
			fib = fibBefore + fibBeforeBefore
			fibBefore, fibBeforeBefore = fib, fibBefore
		}
		return fib
	}
	return uint64(n)
}

func sequence(a, b uint64) (elements []uint64) {
	for i := 0; i <= magicNumber; i++ {
		if a <= fibonacci(i) && fibonacci(i) < b {
			elements = append(elements, fibonacci(i))
		}
	}
	return
}

func prntSeq(elements []uint64) (str string) {
	for _, elem := range elements {
		str += strconv.FormatUint(elem, 10) + ", "
	}
	return strings.Trim(str, ", ")
}

func main() {
	data := os.Args[1:]
	if len(data) != 2 {
		println(errorMessage)
		return
	}
	a, err := getInt(data[0])
	if err != nil {
		println(errorMessage)
		return
	}
	b, err := getInt(data[1])
	if err != nil {
		println(errorMessage)
		return
	}
	if a == b {
		println(errorMessage)
		return
	}
	if a > b {
		a, b = b, a
	}
	fmt.Printf("%v", prntSeq(sequence(a, b)))
}
