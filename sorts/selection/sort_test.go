package selection

import (
	"reflect"
	"testing"
)

func TestComparison(t *testing.T) {
	testCase := []struct {
		name   string
		a      int
		b      int
		comp   func(int, int) bool
		expRes bool
	}{
		{"4>5= 'false'", 4, 5, bigger, false},
		{"8>-5= 'true'", 8, 5, bigger, true},
		{"0<1= 'true'", 0, 1, less, true},
		{"4<6='true'", 4, 6, less, true},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := testCase.comp(testCase.a, testCase.b)
			if testCase.expRes != actRes {
				t.Errorf("Want '%t' but got '%t'", testCase.expRes, actRes)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	testCase := []struct {
		name   string
		a      []int
		expRes int
		comp   func([]int) int
	}{
		{"Asc:Empty slice:", []int{}, 0, asc},
		{"Asc:One element slice", []int{5}, 0, asc},
		{"Asc:Full slice", []int{-4, 3, 5, 2, 6, 1}, 0, asc},
		{"Desc:Empty slice:", []int{}, 0, desc},
		{"Desc:One element slice", []int{5}, 0, desc},
		{"Desc:Full slice", []int{-4, 3, 5, 2, 6, 1}, 4, desc},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := testCase.comp(testCase.a)
			if testCase.expRes != actRes {
				t.Errorf("Want '%v' but got '%v'", testCase.expRes, actRes)
			}
		})
	}

}

func TestSortOnPlace(t *testing.T) {
	testCase := []struct {
		name    string
		a       []int
		expRes  []int
		optfunc func(int, int) bool
	}{
		{"Asc:Empty slice:", []int{}, []int{}, less},
		{"Asc:Simple slice:", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, less},
		{"Asc:Slice with negative:", []int{-4, 3, 5, 2, 6, 1}, []int{-4, 1, 2, 3, 5, 6}, less},
		{"Desc:Empty slice:", []int{}, []int{}, bigger},
		{"Desc:Simple slice:", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, bigger},
		{"Desc:Slice with negatives:", []int{-4, 3, 5, 2, 6, 1}, []int{6, 5, 3, 2, 1, -4}, bigger},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := SortOnPlace(testCase.a, testCase.optfunc)
			if !reflect.DeepEqual(testCase.expRes, actRes) {
				t.Errorf("Exp: %+v -> Act:%+v ", testCase.expRes, actRes)
			}
		})
	}
}

func TestSort(t *testing.T) {
	testCase := []struct {
		name    string
		a       []int
		expRes  []int
		optfunc func([]int) int
	}{
		{"Asc:Empty slice:", []int{}, []int{}, asc},
		{"Asc:One element slice:", []int{1}, []int{1}, asc},
		{"Asc:Slice:", []int{0, 5, 3, 7, 9, 2, 6, 8}, []int{0, 2, 3, 5, 6, 7, 8, 9}, asc},
		{"Asc:Slice with negatives:", []int{-4, 3, 5, 2, 6, 1}, []int{-4, 1, 2, 3, 5, 6}, asc},
		{"Desc:Empty slice:", []int{}, []int{}, desc},
		{"Desc:One element slice:", []int{-4}, []int{-4}, desc},
		{"Desc:Simple slice:", []int{0, 5, 3, 7, 9, 2, 6, 8}, []int{9, 8, 7, 6, 5, 3, 2, 0}, desc},
		{"Desc:Slice with negatives:", []int{-4, 3, 5, 2, 6, 1}, []int{6, 5, 3, 2, 1, -4}, desc},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := Sort(testCase.a, testCase.optfunc)
			if !reflect.DeepEqual(testCase.expRes, actRes) {
				t.Errorf("Exp: %+v -> Act:%+v ", testCase.expRes, actRes)
			}
		})
	}
}
