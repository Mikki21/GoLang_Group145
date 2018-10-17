package main

import "testing"

func TestOrderLetters(t *testing.T) {
	testCase := []struct {
		name   string
		input  letter
		expRes letter
	}{
		{"1.", letter{"Triangle", 1.51, 1.5}, letter{"Triangle", 1.5, 1.51}},
		{"2.", letter{"Triangle", 15, 0}, letter{"Triangle", 0, 15}},
		{"3.", letter{"Triangle", 0, 12}, letter{"Triangle", 0, 12}},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := orderLetters(testCase.input)
			if actRes.a != testCase.expRes.a || actRes.b != testCase.expRes.b {
				t.Errorf("\nExp:%+v\nAct:%+v", testCase.expRes, actRes)
			}
		})

	}
}

func TestCompareLetters(t *testing.T) {
	testCase := []struct {
		name   string
		input1 letter
		input2 letter
		expRes int
	}{
		{"1.", letter{"Triangle_1", 2, 2}, letter{"Triangle_2", 1, 1}, 1},
		{"2.", letter{"Triangle_1", 1, 1}, letter{"Triangle_2", 2, 2}, 2},
		{"3.", letter{"Triangle_!", 1, 2}, letter{"Triangle_2", 2, 1}, 3},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := compareLetters(testCase.input1, testCase.input2)
			if actRes != testCase.expRes {
				t.Errorf("\nExp:%v -> Act:%v", testCase.expRes, actRes)
			}
		})
	}
}

func TestReadConsole(t *testing.T) {
	testCase := []struct {
		name   string
		input1 string
		expRes string
	}{
		{"1.", "Privet", "Privet"},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := readConsole(testCase.input1)
			if actRes != testCase.expRes {
				t.Errorf("\nExp:%v -> Act:%v", testCase.expRes, actRes)
			}
		})
	}
}
func TestFalseChecker(t *testing.T) {
	testCase := []struct {
		name   string
		input1 string
		expRes bool
	}{
		{"1.", "Privet", false},
		{"2.", "y", true},
		{"3.", "yes", true},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := falseChecker(testCase.input1)
			if actRes != testCase.expRes {
				t.Errorf("\nExp:%v -> Act:%v", testCase.expRes, actRes)
			}
		})
	}
}
func TestCreateLetter(t *testing.T) {
	testCase := []struct {
		name         string
		func1        func(string) string
		nameEnvelope string
		height       string
		width        string
		func2        func(letter) letter
		expRes       letter
	}{
		{"Test_1.", readConsole, "name", "10", "5", orderLetters, letter{"name", 5, 10}},
		{"Test_2.", readConsole, "Triangle", "15", "1", orderLetters, letter{"Triangle", 1, 15}},
	}
	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			actRes := createLetter(testCase.func1, testCase.nameEnvelope, testCase.height, testCase.width, testCase.func2)
			if actRes != testCase.expRes {
				t.Errorf("\nExp:%+v\nAct:%+v", testCase.expRes, actRes)
			}
		})
	}
}
