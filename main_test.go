package main

import (
	"testing"
)

type test struct {
	Input  []string
	Output string
	ErrMsg string
}

var tests = []test{
	{
		Input:  []string{},
		Output: "",
		ErrMsg: "Wrong number of arguments - 0. 2 expected",
	},
	{
		Input:  []string{"10"},
		Output: "",
		ErrMsg: "Wrong number of arguments - 1. 2 expected",
	},
	{
		Input:  []string{"a", "12"},
		Output: "",
		ErrMsg: "Wrong value for km - a. Integer expected",
	},
	{
		Input:  []string{"5", "a:b:1"},
		Output: "",
		ErrMsg: "Wrong time format. H:M or M expected",
	},
	{
		Input:  []string{"5", "a:2"},
		Output: "",
		ErrMsg: "Wrong value for hours - a. Integer expected",
	},
	{
		Input:  []string{"5", "1:aa"},
		Output: "",
		ErrMsg: "Wrong value for minutes - aa. Integer expected",
	},
	{
		Input:  []string{"5", "a"},
		Output: "",
		ErrMsg: "Wrong value for minutes - a. Integer expected",
	},
	{
		Input:  []string{"10", "42"},
		Output: "Pace to finish 10 km in 42 - 4m 12s per km",
		ErrMsg: "",
	},
	{
		Input:  []string{"5", "17"},
		Output: "Pace to finish 5 km in 17 - 3m 24s per km",
		ErrMsg: "",
	},
	{
		Input:  []string{"42", "3:30"},
		Output: "Pace to finish 42 km in 3:30 - 5m 00s per km",
		ErrMsg: "",
	},
}

func TestWhatPace(t *testing.T) {
	for _, test := range tests {
		res, err := whatPace(test.Input)
		if res != test.Output {
			t.Fatalf(
				"Output check failed {%v} != {%v} for input - %s",
				res,
				test.Output,
				test.Input,
			)
		}
		if err != nil && err.Error() != test.ErrMsg {
			t.Fatalf(
				"Error check failed {%v} != {%v} for input - %s",
				err,
				test.ErrMsg,
				test.Input,
			)
		}
	}
}
