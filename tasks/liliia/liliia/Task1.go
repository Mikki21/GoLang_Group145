package main

import (
	"fmt"
)

func main() {
	const s1 string = "* * * * "
	const s2 string = " * * * *"
	for i := 0; i < 2; i++ {
		fmt.Println(s1 + "\n" + s2)
	}
}
