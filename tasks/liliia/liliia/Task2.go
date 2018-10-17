package main

import (
	"fmt"
)

func main() {
	var a, b, c, d float64
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Print("Enter c: ")
	fmt.Scan(&c)
	fmt.Print("Enter d: ")
	fmt.Scan(&d)

	switch {
	case a < c && b < d:
		fmt.Println("Great! You can put one envelope into another!")
	case a < d && b < c:
		fmt.Println("Great! You can put one envelope into another!")
	case c < a && d < b:
		fmt.Println("Great! You can put one envelope into another!")
	case c < b && d < a:
		fmt.Println("Great! You can put one envelope into another!")
	default:
		fmt.Println("You can not put one envelope into another!")
	}

}
