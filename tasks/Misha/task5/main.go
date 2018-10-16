package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputCheck(data []string) int {
	return len(data)
}

func getNumber(data []string) (number int, err error) {
	if inputCheck(data) == 2 {
		number, err := strconv.ParseUint(data[1], 10, 64)
		if number == 0 {
			number = 101
			err = nil
		}
		return int(number), err
	}
	err = fmt.Errorf("exception")

	return
}

func convertingNum(number int) (convertedNum string) {
	numTranslator := map[int]string{
		0:  "",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
	if number == 0 {
		convertedNum = "zero"
	} else if number <= 20 {
		convertedNum = numTranslator[number]
	} else {
		convertedNum = numTranslator[number/10*10] + " " + numTranslator[number%10]
		fmt.Println("Converted > 20 ", convertedNum)
	}
	return strings.Trim(convertedNum, " ")
}
func main() {
	number, err := getNumber(os.Args)
	if err != nil {
		fmt.Println("Program converts number to word interpretation. \n\tPossible parametr: ' go run \\main.go <Num>', where <Num>=[0:99]")

	} else {
		if number < 100 {
			fmt.Printf("%v - %v", number, convertingNum(number))
		} else {
			fmt.Println("Invalid <Num> value!\n\tTry again.\n\tPossible parametr: ' go run \\main.go <Num>', where <Num>=[0:99]")
		}
	}
}
