package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readFile(folder string) string {
	input, err := ioutil.ReadFile(folder)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "")
	q := strings.Join(lines, "")
	return q
}

func fileCount(s string, subs string) (a int) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a += strings.Count(scanner.Text(), subs)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return a
}

func changeFile(s string, subs string, change string) {
	input, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "")
	for i := range lines {
		if strings.Contains(lines[i], subs) {
			lines[i] = change
		}
	}
	output := strings.Join(lines, "")
	ioutil.WriteFile(s, []byte(output), 0644)
}

func main() {

	args := os.Args
	switch len(args) {
	case 3:
		fmt.Println("Number of filled strings: ", fileCount(args[1], args[2]))
	case 4:
		changeFile(args[1], args[2], args[3])
		fmt.Printf("Result string in the file:\n\t%v", readFile(args[1]))
	default:
		fmt.Println("Program works with file and has two modes:\n\t 1: counts string in the file 'go run .\\main.go <folder> <string_to_find>\n\t 2: changes strings in file to the entered one 'go run .\\main.go <folder> <string_to_change> <changing_value>")

	}
}
