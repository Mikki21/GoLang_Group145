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

func (letter *letter) letSort() {
	if letter.a > letter.b {
		letter.a, letter.b = letter.b, letter.a
	}
}

func takeTwoLetters() {
	l1 := letter{}
	l2 := letter{}
	fmt.Println("-----Значение для первого конверта-----")
	l1.createLetter()
	fmt.Println("-----Значение для второго конверта-----")
	l2.createLetter()
	l1.letSort()
	l2.letSort()
	if l2.a < l1.a && l2.b < l1.b {
		fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v' ", l2.name, l1.name)
	} else {
		if l2.a > l1.a && l2.b > l1.b {
			fmt.Printf("Конверт '%+v' можно вложить в конверт '%+v' ", l1.name, l2.name)
		} else {
			fmt.Println("Конверты невозможно вложить один в другой !")
		}
	}

}

func (letter *letter) createLetter() {

	var temp string

	fmt.Print("Введите название конверта: ")
	fmt.Fscanln(os.Stdin, &temp)
	letter.name = temp
	for {
		fmt.Print("Значение высоты равно: ")
		fmt.Fscanln(os.Stdin, &temp)
		letter.a, _ = strconv.ParseFloat(temp, 64)
		fmt.Println(letter.a)
		if letter.a == 0 || letter.a <= 0 {
			fmt.Printf("Значение недопустимо '%v', введите повторно !", letter.a)
		} else {
			break
		}
	}
	for {
		fmt.Print("Значение ширины равно:")
		fmt.Fscanln(os.Stdin, &temp)
		letter.b, _ = strconv.ParseFloat(temp, 64)
		if letter.b == 0 || letter.b <= 0 {
			fmt.Printf("Значение недопустимо '%v', введите повторно !", letter.b)
		} else {
			break
		}
	}
}

func falseChecker() bool {
	fmt.Print("Хотите сыграть еще раз ?(y/yes для продолжения): ")
	var container string
	fmt.Fscanln(os.Stdin, &container)
	container = strings.ToLower(container)
	if container == "y" || container == "yes" {
		return true
	}
	return false
}

func main() {
	check := true
	for check {
		takeTwoLetters()
		fmt.Println("__________________________________________________")
		check = falseChecker()
		fmt.Println("__________________________________________________")
	}
}
