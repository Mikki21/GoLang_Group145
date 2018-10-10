package counting

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		expRes []int
	}{
		{"Zero", []int{0}, []int{0}},
		{"Negative", []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1}},
		{"Positive", []int{1, 6, 8, 9, 5, 6, 0, 3, 2, 4, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9}},
		{"Positive and Negative", []int{5, 0, -8, 6, -4, -6, 3, 8, 4}, []int{-8, -6, -4, 0, 3, 4, 5, 6, 8}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := Sort(testCase.arr)
			if !reflect.DeepEqual(actRes, testCase.expRes) {
				t.Errorf("want %d but got %d", testCase.expRes, actRes)
			}
		})

	}
}
