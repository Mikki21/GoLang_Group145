package main

import (
	"fmt"
)

func main() {
	number := 12
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println("один")
	case 2:
		fmt.Println("два")
	case 3:
		fmt.Println("три")
	case 4:
		fmt.Println("четыре")
	case 5:
		fmt.Println("пять")
	case 6:
		fmt.Println("шесть")
	case 7:
		fmt.Println("семь")
	case 8:
		fmt.Println("восемь")
	case 9:
		fmt.Println("девять")
	case 10:
		fmt.Println("десять")
	case 11:
		fmt.Println("одиннадцать")
	case 12:
		fmt.Println("двенадцать")
	case 13:
		fmt.Println("тринадцать")

	default:
		fmt.Println("default")

	}
}
