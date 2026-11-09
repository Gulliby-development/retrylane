package main

import "testing"

var rowsLimits = []row{
	{[]string{"check"}, "a: 200 200 200 200 200 200 200 200 200 200 200 200 200\n", 1, "", "retrylane: line 1: too many outcomes\n"},
	{[]string{"check"}, "a: 200 200 200 200 200 200 200 200 200 200 200 200\n", 0, "ok 1 request\n", ""},
	{[]string{"check"}, "a: 200 200 200 200 200 200 200 200 200 200 200 200 20\n", 1, "", "retrylane: line 1: bad outcome 20\n"},
	{[]string{"check"}, "a: 200 200 200 200 200 200 200 200 200 200 200 200 200\na: 500\n", 1, "", "retrylane: line 1: too many outcomes\n"},
	{[]string{"tally"}, "r1: 500 200\nr2: 500 200\nr3: 500 200\nr4: 500 200\nr5: 500 200\nr6: 500 200\nr7: 500 200\nr8: 500 200\nr9: 500 200\nr10: 500 200\nr11: 500 200\nr12: 500 200\nr13: 500 200\nr14: 500 200\nr15: 500 200\nr16: 500 200\nr17: 500 200\nr18: 500 200\nr19: 500 200\nr20: 500 200\n", 0, "200 20\n500 20\n40 outcomes in 20 requests\nsuccess 20, retryable 20, fatal 0\nretryable share 50.0%\n", ""},
}

func TestRowsLimits(t *testing.T) { checkRows(t, rowsLimits) }
