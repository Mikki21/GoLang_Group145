package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getData() []string {
	return os.Args[1:]
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines string
	in := bufio.NewScanner(file)
	for in.Scan() {
		lines += in.Text()
	}
	return lines, in.Err()
}

func reWrite(path, data string) error {
	file, err := os.Open(path)
	defer file.Close()
	file.WriteString(data)
	return err
}

func parserCount(data []string) {
	subString := data[1]
	mainString, err := readFile(data[0])
	if err == nil {
		fmt.Printf("%v", strings.Count(mainString, subString))
	} else {
		fmt.Printf("%v", err)
	}
}
func parserReplace(data []string) {

	oldSubString := data[1]
	newSubString := data[2]
	newMainString := ""
	oldMainString, err := readFile(data[0])
	if err == nil {
		newMainString = strings.Replace(oldMainString, oldSubString, newSubString, -1)
		err := reWrite(data[0], newMainString)
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else {
		fmt.Printf("%v", err)

	}
}

func errorMessage() {
	fmt.Print("Incorrect input. You should input <path> <string for count> or\n<path> <string for searching> <string for replacement>")
}

func parser(data []string) {
	switch len(data) {
	case 2:
		parserCount(data)
	case 3:
		parserReplace(data)
	default:
		errorMessage()
	}
}

func main() {
	data := getData()
	parser(data)
}
