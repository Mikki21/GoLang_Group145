package counting

import (
	"math"
)

// maxMin returns max and min items from the input slice.
func maxMin(arr []int) (max, min int) {
	if len(arr) == 0 {
		return
	}
	max, min = arr[0], arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

// sortForPositives returns sorted slice of positive integers.
func sortForPositives(arr []int) (res []int) {
	if len(arr) > 1 {
		max, _ := maxMin(arr)
		tmpArr := make([]int, max+1)
		for _, item := range arr {
			tmpArr[item]++
		}
		for i, item := range tmpArr {
			for j := 0; j < item; j++ {
				res = append(res, i)
			}
		}
		return res
	}
	return arr
}

// Sort returns sorted slice of integers
func Sort(arr []int) (res []int) {
	var arrPositive, arrNegative []int

	for _, v := range arr {
		if v > 0 {
			arrPositive = append(arrPositive, v)
		} else {
			arrNegative = append(arrNegative, int(math.Abs(float64(v))))
		}
	}

	arrPositive = sortForPositives(arrPositive)
	arrNegative = sortForPositives(arrNegative)

	for i := 0; i < len(arrNegative)/2; i++ {
		arrNegative[i], arrNegative[len(arrNegative)-i-1] = arrNegative[len(arrNegative)-i-1], arrNegative[i]
	}
	for i := range arrNegative {
		arrNegative[i] *= -1
	}
	return append(arrNegative, arrPositive...)

}
