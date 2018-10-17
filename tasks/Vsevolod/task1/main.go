package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rule string = "You should input height and width (only positive integers), separated by space.\nExample: go run main.go 4 5"

func getData() []string {
	return os.Args[1:]
}

func getInt(str string) (int, error) {
	intValue, err := strconv.ParseInt(str, 10, 64)
	return int(intValue), err
}

func getStrChBoard(height, width int) string {
	strChBoard := ""
	line := append([]string{strings.Repeat("* ", width) + "\n"}, strings.Repeat(" *", width)+"\n")
	for i := 0; i < height; i++ {
		strChBoard += line[i%2]
	}
	return strChBoard
}

func main() {
	data := getData()
	if len(data) != 2 {
		println(rule)
		return
	}
	height, err := getInt(data[0])
	if err != nil {
		println(rule)
		return
	}
	width, err := getInt(data[1])
	if err != nil {
		println(rule)
		return
	}
	if height < 1 || width < 1 {
		println(rule)
		return
	}
	fmt.Print(getStrChBoard(height, width))

}
