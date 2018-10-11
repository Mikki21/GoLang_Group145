package coctail

import (
	"reflect"
	"testing"
)

const amount = 5

func TestSort(t *testing.T) {

	var arrInOut [amount][3][]int
	arrInOut[0][0], arrInOut[0][1], arrInOut[0][2] = []int{}, []int{}, []int{}
	arrInOut[1][0], arrInOut[1][1], arrInOut[1][2] = []int{0}, []int{0}, []int{0}
	arrInOut[2][0], arrInOut[2][1], arrInOut[2][2] = []int{-2, -1, -5, -4, -3}, []int{-5, -4, -3, -2, -1}, []int{-1, -2, -3, -4, -5}
	arrInOut[3][0], arrInOut[3][1], arrInOut[3][2] = []int{1, 6, 8, 9, 5, 6, 0, 3, 2, 4, 7, 8, 9},
		[]int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9}, []int{9, 9, 8, 8, 7, 6, 6, 5, 4, 3, 2, 1, 0}
	arrInOut[4][0], arrInOut[4][1], arrInOut[4][2] = []int{5, 0, -8, 6, -4, -6, 3, 8, 4},
		[]int{-8, -6, -4, 0, 3, 4, 5, 6, 8}, []int{8, 6, 5, 4, 3, 0, -4, -6, -8}

	testCases := []struct {
		name        string
		arrInput    []int
		arrExpRes   []int
		compareFunc func(int, int) bool
	}{
		{"Void", arrInOut[0][0], arrInOut[0][1], Asc},
		{"Zero", arrInOut[1][0], arrInOut[1][1], Asc},
		{"Negative", arrInOut[2][0], arrInOut[2][1], Asc},
		{"Positive", arrInOut[3][0], arrInOut[3][1], Asc},
		{"Positive and Negative", arrInOut[4][0], arrInOut[4][1], Asc},
		{"Void", arrInOut[0][0], arrInOut[0][2], Desc},
		{"Zero", arrInOut[1][0], arrInOut[1][2], Desc},
		{"Negative", arrInOut[2][0], arrInOut[2][2], Desc},
		{"Positive", arrInOut[3][0], arrInOut[3][2], Desc},
		{"Positive and Negative", arrInOut[4][0], arrInOut[4][2], Desc},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := Sort(testCase.arrInput, testCase.compareFunc)
			if !reflect.DeepEqual(actRes, testCase.arrExpRes) {
				t.Errorf("want %d but got %d", testCase.arrExpRes, actRes)
			}
		})

	}
}
