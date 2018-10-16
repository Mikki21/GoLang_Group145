package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func getNum(inStr []string) (int, bool) {
	if len(inStr) == 2 {
		num, err := strconv.ParseInt(inStr[1], 10, 64)
		return int(num), err == nil
	}
	return 0, false
}

func squareNums(n int) (res []int) {
	for i := 1; int(math.Pow(float64(i), 2)) <= n; i++ {
		res = append(res, int(math.Pow(float64(i), 2)))
	}
	return
}

func printRes(res []int) {
	for _, el := range res {
		fmt.Printf("%v ", el)
	}
}
func main() {
	n, err := getNum(os.Args)

	if n > 1 && err {
		printRes(squareNums(n))
	} else {
		fmt.Println("Incorrect input")
	}
}
