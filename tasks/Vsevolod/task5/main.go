package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkData(data []string) bool {
	return len(data) == 2
}

func getNumber(data []string) (number int, res bool) {
	if checkData(data) {
		number, err := strconv.ParseUint(data[1], 10, 64)
		return int(number), err == nil
	}
	return
}

func convertNumber(number int) (convertedNumber string) {
	translateNumber := map[int]string{
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
		convertedNumber = "zero"
	} else if number <= 20 {
		convertedNumber = translateNumber[number]
	} else {
		convertedNumber = translateNumber[number/10*10] + " " + translateNumber[number%10]
	}
	return strings.Trim(convertedNumber, " ")
}

func errorMessage() {
	fmt.Print("Incorrect input. You need input one number from 0 to 99.\n Try again.")
}

func main() {
	number, ok := getNumber(os.Args)
	if ok && number < 100 {
		fmt.Printf("%v - %v", number, convertNumber(number))
	} else {
		errorMessage()
	}
}
