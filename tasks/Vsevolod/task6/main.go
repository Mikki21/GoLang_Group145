package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const message = "Incorrect data. You need to input path.\nFile must consist of word Moskow or Piter"

func readPath() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text(), in.Err()
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	in := bufio.NewScanner(file)
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return strings.Trim(strings.Join(lines, ""), " 	"), in.Err()

}

func moscow() (amt int) {

	for i := 0; i <= 999999; i++ {
		left := i/100000 + i/10000%10 + i/1000%10
		right := i/100%10 + i/10%10 + i%10
		if left == right {
			amt++
		}
	}
	return
}

func piter() (amt int) {
	for i := 0; i <= 999999; i++ {
		odd := i/100000 + i/1000%10 + i/10%10
		even := i/10000%10 + i/100%10 + i%10
		if odd == even {
			amt++
		}
	}
	return
}

func main() {
	fmt.Print("Enter path: ")
	path, err := readPath()
	if err != nil {
		fmt.Printf("%v\n", err)
		println(message)
		return
	}
	text, err := readFile(path)
	if err != nil {
		fmt.Printf("%v", err)
		println(message)
		return
	}
	algorythm := map[string]func() int{
		"Moskow": moscow,
		"Piter":  piter,
	}
	answer, ok := algorythm[text]
	if !ok {
		println(message)
		return
	}
	fmt.Printf("%v", answer())

}
