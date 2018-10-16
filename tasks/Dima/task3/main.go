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
		fmt.Fprintln(os.Stderr, "Entering error: ", err)
	}
	return in.Text()
}

type triangle struct {
	name   string
	a      float64
	b      float64
	c      float64
	square float64
}

func (t *triangle) halfPerimeter() float64 {
	return (t.a + t.b + t.c) / 2
}

func (t *triangle) calcSquare() float64 {
	return math.Sqrt(t.halfPerimeter() * (t.halfPerimeter() - t.a) * (t.halfPerimeter() - t.b) * (t.halfPerimeter() - t.c))
}

func isTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Incorrect input")
		return false
	} else if a+b <= c || b+c <= a || c+a <= b {
		fmt.Println("It`s not a triangle, try again")
		return false
	}
	return true
}

func getData() (values []string) {
	fmt.Print("Please, enter parameters of your Triangle:\n")
	for _, value := range strings.Split(scan(), ",") {
		values = append(values, strings.Trim(value, " 	"))
	}
	return
}

func (t *triangle) setTriangle() {

	for {
		data := getData()
		if len(data) != 4 {
			fmt.Println("Incorrect number of values")
			continue
		}

		var emptyItems int
		for _, item := range data {
			if item == "" {
				emptyItems++
			}
		}

		if emptyItems > 0 {
			fmt.Println("Empty value, try again")
			continue
		}

		var err error
		t.name = data[0]
		t.a, err = strconv.ParseFloat(data[1], 64)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		t.b, err = strconv.ParseFloat(data[2], 64)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		t.c, err = strconv.ParseFloat(data[3], 64)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		if isTriangle(t.a, t.b, t.c) {
			t.square = t.calcSquare()
			break
		}
	}
}

func addTriangle() (newTriangle triangle) {
	newTriangle.setTriangle()
	return
}

func sort(t []triangle) {
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i].square < t[j].square {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
}

func getAnswer() bool {
	fmt.Print("Would you like to add one more triangle? If yes, press y/yes.\n")
	res := strings.ToLower(scan())
	return res == "y" || res == "yes"
}

func main() {
	triangles := []triangle{}
	cont := true
	for cont {
		triangles = append(triangles, addTriangle())
		cont = getAnswer()
	}
	sort(triangles)
	fmt.Print("================Triangle list:================\n")
	for _, item := range triangles {
		fmt.Printf("[%v]: %.2f cm\n", item.name, item.square)
	}

}
