package merging

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name       string
		inputSlice []int
		expSlice   []int
		compFunc   func(int, int) bool
	}{
		{"Already sorted use Asc ", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, Asc},
		{"Just simple case use Asc ", []int{0, 2, 0, 1, 0}, []int{0, 0, 0, 1, 2}, Asc},
		{"Case with one element use Asc", []int{1}, []int{1}, Asc},
		{"Case with neg and poz numbers use Asc", []int{7, 3, 5, 7, 8, 4, 3, 65, -43, 6, -5}, []int{-43, -5, 3, 3, 4, 5, 6, 7, 7, 8, 65}, Asc},
		{"Empty slice use Asc", []int{}, []int{}, Asc},

		{"The worth case use Desc ", []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}, Desc},
		{"Just simple case use Desc ", []int{0, 2, 0, 1, 0}, []int{2, 1, 0, 0, 0}, Desc},
		{"Case with one element use Desc", []int{1}, []int{1}, Desc},
		{"Case with neg and poz numbers use Desc", []int{7, 3, 5, 7, 8, 4, 3, 65, -43, 6, -5}, []int{65, 8, 7, 7, 6, 5, 4, 3, 3, -5, -43}, Desc},
		{"Empty slice use Desc", []int{}, []int{}, Desc},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actSlice := Sort(testCase.inputSlice, testCase.compFunc)
			if !reflect.DeepEqual(testCase.expSlice, actSlice) {
				t.Errorf("want %v but got %v", testCase.expSlice, actSlice)
			}
		})
	}
}
