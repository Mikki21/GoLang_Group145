package main

import (
	"testing"
)

func TestParseParameters(t *testing.T) {
	testCase := []struct {
		name        string
		sliceString []string
		expResH     int
		expResW     int
	}{
		{"1.", []string{"", "-1", "-1"}, 0, 0},
		{"2.", []string{"", "2", "2"}, 2, 2},
		{"2.", []string{"", "2", "2"}, 2, 2},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actH, actW, actErr := parseParametrs(testCase.sliceString)
			if actH != testCase.expResH || actW != testCase.expResW {
				t.Errorf("Exp:%v , %v -> Act:%v , %v\n ActError: %v",
					testCase.expResH, testCase.expResW, actH, actW, actErr)
			}
		})
	}
}

func TestCreateDesk(t *testing.T) {
	testCase := []struct {
		name   string
		h      int
		w      int
		expRes string
	}{
		{"2x2: ", 2, 2, "* * \n * *\n"},
		{"3x3: ", 3, 3, "* * * \n * * *\n* * * \n"},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes, err := createDesk(testCase.h, testCase.w)
			if actRes != testCase.expRes {
				t.Errorf("Exp:\n%s\nAct:\n%s\n ActError: %v", testCase.expRes, actRes, err)
			}
		})
	}
}
