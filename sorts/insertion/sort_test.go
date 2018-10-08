package insertion

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expRes   []int
		compFunc func(int, int) bool
	}{
		{"Zero in Acc", []int{0}, []int{0}, Acc},
		{"Negative in Acc", []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1}, Acc},
		{"Positive in Acc", []int{1, 6, 8, 9, 5, 6, 0, 3, 2, 4, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9}, Acc},
		{"Positive and Negative in Acc", []int{5, 0, -8, 6, -4, -6, 3, 8, 4}, []int{-8, -6, -4, 0, 3, 4, 5, 6, 8}, Acc},
		{"Zero in Dec", []int{0}, []int{0}, Dec},
		{"Negative in Dec", []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1}, Dec},
		{"Positive in Dec", []int{1, 6, 8, 9, 5, 6, 0, 3, 2, 4, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9}, Dec},
		{"Positive and Negative in Dec", []int{5, 0, -8, 6, -4, -6, 3, 8, 4}, []int{-8, -6, -4, 0, 3, 4, 5, 6, 8}, Dec},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := Sort(testCase.arr, Acc)
			if !reflect.DeepEqual(actRes, testCase.expRes) {
				t.Errorf("want %d but got %d", testCase.expRes, actRes)
			}
		})

	}
}
