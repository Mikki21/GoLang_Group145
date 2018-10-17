package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readConsole(check string) (res string) {
	if check == "" {
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		res = reader.Text()
		return
	}
	res = check
	return
}

func newTriangle(sent triangle) triangle {
	errTriangle := triangle{}
	s := strings.Split(readConsole(""), ",")
	if len(s) == 4 {
		for i := 1; i < len(s); i++ {
			s[i] = strings.Replace(s[i], " ", "", -1)
			s[i] = strings.Replace(s[i], "	", "", -1)
		}

		a, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return errTriangle
		}

		b, err := strconv.ParseFloat(s[2], 64)
		if err != nil {
			return errTriangle
		}

		c, err := strconv.ParseFloat(s[3], 64)
		if err != nil {
			return errTriangle
		}
		if a+b > c && a+c > b && b+c > a {
			sent.name = s[0]
			p := (a + b + c) / 2
			sent.pl = math.Sqrt(p * (p - a) * (p - b) * (p - c))
			return sent
		}
		// fmt.Println(s)
	}
	return errTriangle
}

//struct for triangles
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

//fmt.Print("Do you want to add another triangle (enter y/yes to confirm)? ")
func falseChecker(checkStr string) bool {
	checkStr = strings.ToLower(checkStr)
	if checkStr == "y" || checkStr == "yes" {
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
	var x triangle
	var sliceOfTriangle []triangle
	check := true
	for check {
		fmt.Print("Enter new triangle <name>,<side_1>,<side_2>,<side_3>\n")
		x = newTriangle(x)
		if x.pl == 0 {
			fmt.Println("Impossible to create triangle with entered parameters...")
		} else {
			sliceOfTriangle = append(sliceOfTriangle, x)
		}
		fmt.Println("Do you want to add another triangle (enter y/yes to confirm)? ")
		check = falseChecker(readConsole(""))
	}
	squaresList := SortOnPlace(addPlosha(sliceOfTriangle), bigger)
	fmt.Println("=====================Triangles list:=====================")
	for i := 0; i < len(sliceOfTriangle); i++ {
		fmt.Printf("\t%v.[Triangle %v]: %.2f cm\n", i+1, sliceOfTriangle[i].name, squaresList[i])
	}
}
