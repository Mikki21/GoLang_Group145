package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const amtParam = 4
const message = "Incorrect input. You should input <name>,<a>,<b>,<b>."

func scan() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	err := in.Err()
	if err != nil {
		return "", err
	}
	return in.Text(), err
}

func getAnswer(str string) bool {
	str = strings.ToLower(str)
	return str == "y" || str == "yes"
}

func getDataTriangle(scan func() (string, error)) (outpstrs []string, err error) {
	inpstr, err := scan()
	if err != nil {
		return nil, err
	}
	elements := strings.Split(inpstr, ",")
	if len(elements) != amtParam {
		return nil, errors.New("Incorrect amount of parameters")
	}
	for _, elem := range elements {
		outpstrs = append(outpstrs, elem)
	}
	return
}

func getFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func checkTriangle(a, b, c float64) bool {
	isPositiveSides := a > 0 && b > 0 && c > 0
	isExist := a+b > c && a+c > b && b+c > a
	return isPositiveSides && isExist
}

type triangle struct {
	name string
	a    float64
	b    float64
	c    float64
	area float64
}

func (t *triangle) setArea() {
	if t.area == 0 {
		p := (t.a + t.b + t.c) / 2
		t.area = math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	}
}

func sort(t []triangle) {
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i].area < t[j].area {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
}

func addTriangle() (newTriangle triangle, ok bool) {
	fmt.Println("Enter Triangle: ")
	data, err := getDataTriangle(scan)
	if err != nil {
		fmt.Println(message)
		return
	}
	a, err := getFloat(data[1])
	if err != nil {
		fmt.Println(message)
		return
	}
	b, err := getFloat(data[2])
	if err != nil {
		fmt.Println(message)
		return
	}
	c, err := getFloat(data[3])
	if err != nil {
		fmt.Println(message)
		return
	}
	if !checkTriangle(a, b, c) {
		fmt.Println("Triangle with your sides don't exist")
		return
	}
	name := data[0]
	newTriangle = triangle{name: name, a: a, b: b, c: c}
	newTriangle.setArea()
	return newTriangle, true
}

func printStrTr(t []triangle) (str string) {
	str = "=================Triangle list:=================\n"
	for i, item := range t {
		str += strconv.Itoa(i+1) + ". [" + item.name + "]: " + fmt.Sprintf("%.2f", item.area) + " cm\n"
	}
	return
}

func main() {
	triangles := []triangle{}

	answer := true
	for answer {
		var NewTriangle triangle
		var ok bool
		for !ok {
			NewTriangle, ok = addTriangle()
		}
		triangles = append(triangles, NewTriangle)
		fmt.Println("Would you like to add one more triangle? If yes, press y/yes: ")
		respons, err := scan()
		if err != nil {
			fmt.Println(message)
			return
		}
		answer = getAnswer(respons)
	}
	sort(triangles)
	fmt.Println(printStrTr(triangles))
}
