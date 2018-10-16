package main

import (
	"fmt"
	"os"
	"strconv"
)

func osChecks(data []string) int {
	return len(data)
}

func getParametrs(data []string) (int, int, error) {
	if osChecks(data) == 3 {
		p1, err := strconv.ParseInt(data[1], 10, 64)
		p2, err := strconv.ParseInt(data[2], 10, 64)
		return int(p1), int(p2), err
	} else {
		if osChecks(data) == 1 {
			err := "exception"
		}
	}

}
func main() {
	param1, param2, err := getParametrs(os.Args)
	if err != nil {
		fmt.Println("program ......")
	} else {
		//fibonach(param1, param2)
	}

}
