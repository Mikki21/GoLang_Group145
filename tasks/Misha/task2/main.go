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
		//fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v' ", l2.name, l1.name)
	}
	if l2.a > l1.a && l2.b > l1.b {
		a = 2
		return
		//fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v' ", l1.name, l2.name)
	}
	a = 3
	return
	//fmt.Println("Конверты невозможно вложить один в другой !")
}
func createLetter(readCnl func(string) string, imya string, visota string, shirina string, orderLet func(letter) letter) (let letter) {
	var temp string
	fmt.Println("Введите название конверта: ")
	let.name = readCnl(imya)

	for {
		fmt.Print("Значение высоты равно: ")
		temp = readCnl(visota)
		let.a, _ = strconv.ParseFloat(temp, 64)
		if let.a <= 0 {
			fmt.Printf("Значение недопустимо '%v', введите повторно !\n", let.a)
		} else {
			break
		}
	}
	for {
		fmt.Print("Значение ширины равно:")
		temp = readCnl(shirina)
		let.b, _ = strconv.ParseFloat(temp, 64)
		if let.b <= 0 {
			fmt.Printf("Значение недопустимо '%v', введите повторно !\n", let.b)
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
		fmt.Println("\tПервый конерт!")
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
