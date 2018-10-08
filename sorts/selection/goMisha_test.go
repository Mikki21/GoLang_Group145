package goMisha

import (
	"reflect"
	"testing"
)

func TestMore(t *testing.T) {
	testCase := []struct {
		name   string
		a      int
		b      int
		expRes bool
	}{
		{"4<5= 'true'", 4, 5, true},
		{"8<-5= 'false'", 8, 5, false},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := more(testCase.a, testCase.b)
			if testCase.expRes != actRes {
				t.Errorf("Want '%t' but got '%t'", testCase.expRes, actRes)
			}
		})
	}
}

func TestLess(t *testing.T) {
	testCase := []struct {
		name   string
		a      int
		b      int
		expRes bool
	}{
		{"1>0=true", 1, 0, true},
		{"0>1=false", 0, 1, false},
		{"4>6=false", 4, 6, false},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := less(testCase.a, testCase.b)
			if testCase.expRes != actRes {
				t.Errorf("Want '%t' but got '%t'", testCase.expRes, actRes)
			}
		})
	}
}

func TestSelectSortOperate(t *testing.T) {
	testCase := []struct {
		a      []int
		expRes []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{-4, 3, 5, 2, 6, 1}, []int{-4, 1, 2, 3, 5, 6}},
	}
	for _, testCase := range testCase {
		actRes := SelectSortOperate(testCase.a, more)
		if !reflect.DeepEqual(testCase.expRes, actRes) {
			t.Errorf("Exp: %+v -> Act:%+v ", testCase.expRes, actRes)
		}
	}
}

func TestSelectSortAppModifyMin(t *testing.T) {
	testCase := []struct {
		a      []int
		expRes []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{0, 5, 3, 7, 9, 2, 6, 8}, []int{0, 2, 3, 5, 6, 7, 8, 9}},
		{[]int{-4, 3, 5, 2, 6, 1}, []int{-4, 1, 2, 3, 5, 6}},
	}
	for _, testCase := range testCase {
		actRes := SelectSortAppModify(testCase.a, Min)

		if !reflect.DeepEqual(testCase.expRes, actRes) {
			t.Errorf("Exp: %+v -> Act:%+v ", testCase.expRes, actRes)
		}
	}
}

func TestSelectSortAppModifyMax(t *testing.T) {
	testCase := []struct {
		a      []int
		expRes []int
	}{
		{[]int{}, []int{}},
		{[]int{-4}, []int{-4}},
		{[]int{0, 5, 3, 7, 9, 2, 6, 8}, []int{9, 8, 7, 6, 5, 3, 2, 0}},
		{[]int{-4, 3, 5, 2, 6, 1}, []int{6, 5, 3, 2, 1, -4}},
	}
	for _, testCase := range testCase {
		actRes := SelectSortAppModify(testCase.a, Max)

		if !reflect.DeepEqual(testCase.expRes, actRes) {
			t.Errorf("Exp: %+v -> Act:%+v ", testCase.expRes, actRes)
		}
	}
}
