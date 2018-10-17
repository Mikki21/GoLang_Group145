package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPath() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text(), in.Err()
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	in := bufio.NewScanner(file)
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return lines, in.Err()

}

func fortunateTicket(text []string) (string, int) {
	indicator := strings.Trim(strings.Join(text, ""), " 	")
	switch indicator {
	case "Moskow":
		return "Moskow", moscow()
	case "Piter":
		return "Piter", piter()
	default:
		return "Incorrect Indicator. It must be Moscow or Piter. Only one word", -1
	}
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

func printFortunateTicket(s []string) {
	indicator, amt := fortunateTicket(s)
	if amt != -1 {
		fmt.Printf("%v %v", indicator, amt)
	} else {
		fmt.Printf("%v", indicator)
	}

}

func main() {

	fmt.Print("Enter path: ")
	path, errPath := readPath()
	if errPath == nil {
		data, err := readFile(path)
		if err == nil {
			printFortunateTicket(data)
		} else {
			fmt.Printf("%v", err)
		}
	} else {
		fmt.Printf("%v", errPath)
	}
}
