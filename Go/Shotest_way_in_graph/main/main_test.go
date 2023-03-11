package main

import (
	"testing"
)

type testPair struct {
	num    string
	result int
}

var tests = []testPair{
	{"1", 2},
	{"2", 1},
	{"3", -1},
	{"4", 3},
	{"5", 1},
	{"6", -1},
	{"7", 0},
}

func TestRealMain(t *testing.T) {
	for _, test := range tests {
		result := RealMain("tests/" + test.num + ".txt")
		if result != test.result {
			t.Error(
				"For test", test.num,
				"expected", test.result,
				"got", result,
			)
		}
	}
}
