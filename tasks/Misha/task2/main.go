package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type letter struct {
	name string
	a    float64
	b    float64
}

func orderLetters(let letter) letter {
	if let.a > let.b {
		let.a, let.b = let.b, let.a
	}
	return let
}

func compareLetters(l1 letter, l2 letter) (a int) {
	if l2.a < l1.a && l2.b < l1.b {
		a = 1
		return
	}
	if l2.a > l1.a && l2.b > l1.b {
		a = 2
		return

	}
	a = 3
	return
}
func createLetter(readCnl func(string) string, nameEnvelope string, height string, width string, orderLet func(letter) letter) (let letter) {
	var temp string
	fmt.Println("Введите название конверта: ")
	let.name = readCnl(nameEnvelope)

	for {
		fmt.Print("Значение высоты равно: ")
		temp = readCnl(height)
		let.a, _ = strconv.ParseFloat(temp, 64)
		if let.a <= 0 {
			fmt.Printf("Значение недопустимо , введите повторно !\n")
		} else {
			break
		}
	}
	for {
		fmt.Print("Значение ширины равно:")
		temp = readCnl(width)
		let.b, _ = strconv.ParseFloat(temp, 64)
		if let.b <= 0 {
			fmt.Printf("Значение недопустимо , введите повторно !\n")
		} else {
			break
		}
	}

	let = orderLet(let)
	return
}

func readConsole(check string) (res string) {
	if check == "" {
		fmt.Fscanln(os.Stdin, &check)
		res = check
		return
	}
	res = check
	return
}

func falseChecker(checkStr string) bool {
	checkStr = strings.ToLower(checkStr)
	if checkStr == "y" || checkStr == "yes" {
		return true
	}
	return false
}

func main() {
	check := true
	for check {
		fmt.Println("\tПервый конверт!")
		exmp := createLetter(readConsole, "", "", "", orderLetters)
		fmt.Println("\tВторой конверт!")
		exmp1 := createLetter(readConsole, "", "", "", orderLetters)
		checkEnv := compareLetters(exmp, exmp1)
		switch checkEnv {
		case 1:
			fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v'.\n", exmp1.name, exmp.name)
		case 2:
			fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v'.\n", exmp.name, exmp1.name)
		case 3:
			fmt.Println("Конверты невозможно вложить один в другой !")
		}
		fmt.Print("Хотите сыграть еще раз ?(y/yes для продолжения): ")
		check = falseChecker(readConsole(""))
		fmt.Println()
	}
}
