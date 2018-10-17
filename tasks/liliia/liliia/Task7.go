package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	x := 0
	for {
		y := x * x
		if y < n {
			fmt.Print(y, " ")
			x++
		}
	}
}
