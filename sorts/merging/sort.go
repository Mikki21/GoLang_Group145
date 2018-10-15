package merging

//This is mergesort algortim which uses to sort slices
//Used src: https://en.wikipedia.org/wiki/Merge_sort

func Sort(slice []int, comparing func(int, int) bool) []int {

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

	left = Sort(left, comparing)
	right = Sort(right, comparing)

	return merge(left, right, comparing)
}

//merge function merge slices and return resulting slices
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

func Asc(leftEl, rightEl int) bool {
	return leftEl <= rightEl
}

func Desc(leftEl, rightEl int) bool {
	return leftEl >= rightEl
}
