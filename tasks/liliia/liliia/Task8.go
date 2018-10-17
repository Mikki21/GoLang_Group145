package main

import (
	"fmt"
	"math"
)

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
func main() {
	var a, b, i float64
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	i = math.Min(a, b)
	for ; i < math.Max(a, b); i++ {
		fmt.Println(i)
	}
}
