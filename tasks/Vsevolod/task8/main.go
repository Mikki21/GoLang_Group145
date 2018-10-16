package main

import (
	"fmt"
	"os"
	"strconv"
)

func getData(osArgs []string) (uint64, uint64, bool) {
	if len(osArgs) == 3 {
		lower, errLower := strconv.ParseUint(osArgs[1], 10, 64)
		upper, errUpper := strconv.ParseUint(osArgs[2], 10, 64)
		return lower, upper, errLower == nil && errUpper == nil
	}
	return 0, 0, false
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

func swap(a, b uint64) (uint64, uint64) {
	if b < a {
		return b, a
	}
	return a, b
}

func sequence(a, b uint64) (elements []uint64) {

	for i := 0; i <= 93; i++ {
		if a <= fibonacci(i) && fibonacci(i) < b {
			elements = append(elements, fibonacci(i))
		}
	}
	return
}

func errorMessage() {
	fmt.Print("Incprrect parameters.\nYou should input 2 different\nintegers from 0 to 18446744073709551615.\n Try again")
}

func printSequence(elements []uint64) {
	switch len(elements) {
	case 0:
		fmt.Print("No elements in this range")
	case 1:
		fmt.Printf("%v", elements[0])
	default:
		for i := 0; i < len(elements)-1; i++ {
			fmt.Printf("%v, ", elements[i])
		}
		fmt.Printf("%v", elements[len(elements)-1])
	}
}

func main() {
	a, b, isUint := getData(os.Args)
	if isUint {
		a, b = swap(a, b)
		printSequence(sequence(a, b))
	} else {
		errorMessage()
	}
}
