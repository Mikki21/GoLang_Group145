package Merge

import "fmt"

//This is mergesort algortim which uses to sort slices
//Used src: https://en.wikipedia.org/wiki/Merge_sort

func MergeSort(slice []int) []int {

	if len(slice) <= 1 {
		return slice
	}

	left := make([]int, 0)
	right := make([]int, 0)
	mid := len(slice) / 2

	for i, x := range slice {
		switch {
		case i < mid:
			left = append(left, x)
		case i >= mid:
			right = append(right, x)
		}
	}

	left = MergeSort(left)
	right = MergeSort(right)

	fromLower := func(leftEl, rightEl int) bool {
		return leftEl <= rightEl
	}

	fromBigger := func(leftEl, rightEl int) bool {
		return leftEl >= rightEl
	}

	_ = fromLower
	_ = fromBigger

	return merge(left, right, fromLower)
}

func merge(left, right []int, comp func(int, int) bool) []int {

	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if comp(left[0], right[0]) {
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:len(left)]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	return results
}

func main() {

	sl := []int{1, 6, -2, 7, 2, 99, 23, 311, 7, -8}
	fmt.Println("Unsorted slice: ", sl)
	sl = MergeSort(sl)
	fmt.Println("Sorted slice: ", sl)
}
