package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func deleteSpaces(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "	", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	return s
}

func makeSlice() (res []float64, err error) {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	s := reader.Text()
	s = deleteSpaces(s)

	limit, err := strconv.ParseFloat(s, 64)
	if err != nil || limit < 0 {
		err = fmt.Errorf("Exception")
		return
	}

	for i := 0.0; i < limit; i++ {
		if i*i < limit {
			res = append(res, i)
		}
	}

	return
}

func printLine(print []float64) {
	for i := 0; i < len(print); i++ {
		if i == len(print)-1 {
			fmt.Printf("%v.", print[i])
		} else {
			fmt.Printf("%v, ", print[i])
		}
	}
}

func falseChecker() bool {
	fmt.Print("\nDo you want to run once more ?(y/yes to continue): ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	s := strings.Replace(reader.Text(), " ", "", -1)
	s = strings.Replace(s, "	", "", -1)
	s = strings.ToLower(s)
	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func main() {
	fmt.Println("Program prints list of numbers squares of which are less then entered 'n'(where 'n'>0).")
	check := true
	for check {
		fmt.Println("\tEnter 'n' value:")
		try, err := makeSlice()
		if err != nil {
			fmt.Print("Please, enter valid number 'n'.\nProgram prints list of numbers squares of which are less then entered 'n'() where n>0).")
		} else {
			printLine(try)
		}
		check = falseChecker()
	}
}
