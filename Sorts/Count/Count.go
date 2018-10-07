package Count

import (
	"math"
)

//MaxMin ...
func MaxMin(arr []int) (max, min int) {
	max = arr[0]
	min = arr[0]
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

//CountSort ...
func CountSort(arr []int) (res []int) {
	if len(arr) > 1 {

		max, _ := MaxMin(arr)
		help := make([]int, max+1)
		for _, v := range arr {
			help[v]++
		}
		for i, v := range help {
			for j := 0; j < v; j++ {
				res = append(res, i)
			}
		}
		return res
	} else {
		return arr
	}
}

//CountSortSigned ...
func CountSortSigned(arr []int) (res []int) {
	var arrPositive []int
	var arrNegative []int
	for _, v := range arr {
		if v > 0 {
			arrPositive = append(arrPositive, v)
		} else {
			arrNegative = append(arrNegative, int(math.Abs(float64(v))))
		}
	}

	arrPositive = CountSort(arrPositive)
	arrNegative = CountSort(arrNegative)

	for i := 0; i < len(arrNegative)/2; i++ {
		arrNegative[i], arrNegative[len(arrNegative)-i-1] = arrNegative[len(arrNegative)-i-1], arrNegative[i]
	}
	for i := range arrNegative {
		arrNegative[i] *= -1
	}
	return append(arrNegative, arrPositive...)

}
