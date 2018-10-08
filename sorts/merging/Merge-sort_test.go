package Merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		inputSlice []int
		expSlice   []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 2, 0, 1, 0}, []int{0, 0, 0, 1, 2}},
		{[]int{1}, []int{1}},
		{[]int{7, 3, 5, 7, 8, 4, 3, 65, -43, 6, -5}, []int{-43, -5, 3, 3, 4, 5, 6, 7, 7, 8, 65}},
	}

	for _, testCase := range testCases {
		actSlice := MergeSort(testCase.inputSlice)

		if !reflect.DeepEqual(testCase.expSlice, actSlice) {
			t.Errorf("want %v but got %v", testCase.expSlice, actSlice)
		}
	}
}
