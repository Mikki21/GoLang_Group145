package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func newTriangle() (res triangle) {
	fmt.Print("Enter new triangle <name>,<side_1>,<side_2>,<side_3>\n")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	errTriangle := triangle{pl: 0.0}
	s := strings.Split(reader.Text(), ",")
	if len(s) == 4 {
		for i := 1; i < len(s); i++ {
			s[i] = strings.Replace(s[i], " ", "", -1)
			s[i] = strings.Replace(s[i], "	", "", -1)
		}
		a, err := strconv.ParseFloat(s[1], 64)
		b, err := strconv.ParseFloat(s[2], 64)
		c, err := strconv.ParseFloat(s[3], 64)
		_ = err
		fmt.Println(a, b, c)
		if a+b > c && a+c > b && b+c > a {
			res.name = s[0]
			p := a + b + c
			res.pl = math.Sqrt(p * (p - a) * (p - b) * (p - c))
			return res
		}
		// fmt.Println(s)
	}
	return errTriangle
}

type triangle struct {
	name string
	pl   float64
}

//bigger funcs to compare integers
func bigger(a, b float64) (k bool) {
	if a > b {
		k = true

	}
	return k
}

//SortOnPlace sort array of int
func SortOnPlace(a []float64, operation func(a, b float64) bool) []float64 {
	k := len(a)
	for i := 0; i < k; i++ {
		indMin := i
		for j := i; j < k; j++ {
			if operation(a[j], a[indMin]) {
				indMin = j
			}
		}
		a[i], a[indMin] = a[indMin], a[i]
	}
	return a
}

func falseChecker() bool {
	fmt.Print("Do you want to add another triangle (enter y/yes to confirm)? ")
	var container string
	fmt.Fscanln(os.Stdin, &container)
	container = strings.ToLower(container)
	if container == "y" || container == "yes" {
		return true
	}
	return false
}

func addPlosha(x []triangle) (x1 []float64) {
	for i := 0; i < len(x); i++ {
		x1 = append(x1, x[i].pl)
	}
	return x1
}

func main() {
	x := []triangle{}
	check := true
	for check {
		if s := newTriangle(); s.pl == 0 {
			fmt.Println("Impossible to create triangle with entered parameters...")
		} else {
			x = append(x, s)
		}
		check = falseChecker()
	}
	x1 := SortOnPlace(addPlosha(x), bigger)
	fmt.Println("=====================Triangles list:=====================")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v.[Triangle %v]: %.2f cm\n", i+1, x[i].name, x1[i])
	}
}
