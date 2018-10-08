package main

import (
	"fmt"
)

func opDec(a, b []int) []int {
	a = a[j]
	b = a[j-1]
	fromSmaller := func(a, b int) bool {
		return a <= b
	}

	fromBigger := func(a, b int) bool {
		return a >= b
	}
	return opDec(a, b, fromBigger)
}

// InsertionSort
// At each iteration, InsertionSort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there. It repeats until no input elements remain.
// See: https://ru.wikipedia.org/wiki/Сортировка_вставками
func InsertionSort(a []int, operation func(int, int) bool) []int {
	var n = len(a)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}
	return a
}

func main() {
	a := []int{10, 77, 4, 68, 1, 707, 8, 12, 17, 8, 21, 15, 9, 55, 18, 58, 101, 48}
	fmt.Printf("elements before sorting = %v\n", a)
	a = InsertionSort(a)
	fmt.Printf("sorted elements = %v\n", a)
}

/*elements sorted in descending order


//elements sorted in ascending order


func main() {
	a := []int{10, 77, 4, 68, 1, 707, 8, 12, 17, 8, 21, 15, 9, 55, 18, 58, 101, 48}
	fmt.Printf("Elements before sorting = %v\n", a)
	a = InsertionSort(a, )
	fmt.Printf("Sorted elements = %v\n", a)
	//b := InsertionSort(elem[:], actionDec(desc))
	//fmt.Printf("Elements sorted in descending order: %v", b)*/
