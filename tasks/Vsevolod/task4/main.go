package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(data)
	return err
}

func parserCount(data []string) (string, error) {
	path := data[0]
	subString := data[1]
	mainString, err := readFile(path)
	if err != nil {
		return "", err
	}
	return "Quantity of substring " + subString + " in " + path + " is " + strconv.Itoa(strings.Count(mainString, subString)), err
}

func parserReplace(data []string) (string, error) {
	path := data[0]
	oldSubString := data[1]
	newSubString := data[2]
	oldMainString, err := readFile(path)
	if err != nil {
		return "", err
	}
	newMainString := strings.Replace(oldMainString, oldSubString, newSubString, -1)
	err = reWrite(path, newMainString)
	return "Replacement " + oldSubString + " to " + newSubString + " in " + path + " is successfully done.", err
}

func parser(data []string) (string, error) {
	switch len(data) {
	case 2:
		return parserCount(data)
	case 3:
		return parserReplace(data)
	default:
		return "", errors.New("Incorrect input. You should input: <path> <string for count> or\n" +
			"<path> <string for searching> <string for replacement>")
	}
}

func main() {
	data := getData()
	answer, err := parser(data)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(answer)
	}
}
