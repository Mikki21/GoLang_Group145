package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Entering mistake", err)
	}
	return in.Text()
}

func getAnswer() bool {
	fmt.Print("Sir(Miss), would you like to add one more triangle? If yes, press y/yes.\n")
	res := strings.ToLower(scan())
	return res == "y" || res == "yes"
}

func getData() (values []string) {
	fmt.Print("Please, enter parameters of your Triangle:\n")
	for _, value := range strings.Split(scan(), ",") {
		values = append(values, strings.Trim(value, " 	"))
	}
	return
}

func warningMessage() {
	fmt.Print("Incorrect input. You have to input 4 parametrs,\n separated only by ,: <name>,<a>,<b>,<c>.\n",
		"It must be name of Triangle and sides\nTry again.\n")

}

func (t *triangle) setTriangle() {
	var triangleName string
	var triangleSides [3]float64
	var err error
	for {
		data := getData()
		if len(data) != 4 {
			warningMessage()
			continue
		}

		var countVoidItems int
		for _, item := range data {
			if item == "" {
				countVoidItems++
			}
		}
		if countVoidItems > 0 {
			warningMessage()
			continue
		}

		triangleName = data[0]

		for i := 0; i < 3; i++ {
			var errTemp error
			triangleSides[i], errTemp = strconv.ParseFloat(data[i+1], 64)
			if errTemp != nil {
				err = errTemp
			}
		}

		if err == nil && checkTriangle(triangleSides) {
			break
		}
		if !checkTriangle(triangleSides) {
			fmt.Print("You input wrong sides of Triangle.\n")
		}
		warningMessage()
	}
	t.name = triangleName
	t.a = triangleSides[0]
	t.b = triangleSides[1]
	t.c = triangleSides[2]
}

func checkTriangle(sides [3]float64) bool {
	isPositiveSides := sides[0] > 0 && sides[1] > 0 && sides[2] > 0
	isExist := sides[0]+sides[1] > sides[2] && sides[1]+sides[2] > sides[0] && sides[0]+sides[2] > sides[1]
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

func addTriangle() (newTriangle triangle) {
	newTriangle.setTriangle()
	newTriangle.setArea()
	return
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

func main() {
	triangles := []triangle{}
	answer := true
	for answer {
		triangles = append(triangles, addTriangle())
		answer = getAnswer()
	}
	sort(triangles)
	fmt.Print("=================Triangle list:=================\n")
	for _, item := range triangles {
		fmt.Printf("[%v]: %.2f cm\n", item.name, item.area)
	}

}
