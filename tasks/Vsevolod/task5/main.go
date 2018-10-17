package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const message = "Incorrect input. You need to input one number from 0 to 99. Try again."

func getInt(data string) (int, error) {
	number, err := strconv.ParseUint(data, 10, 64)
	return int(number), err
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

func main() {
	data := os.Args[1:]
	if len(data) != 1 {
		println(message)
		return
	}
	number, err := getInt(data[0])
	if err != nil {
		println(message)
		return
	}
	fmt.Printf("%v - %v", number, convertNumber(number))

}
