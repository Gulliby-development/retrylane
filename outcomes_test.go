package main

import "testing"

var rowsOutcomes = []row{
	{[]string{"check"}, "a: 99\n", 1, "", "retrylane: line 1: bad outcome 99\n"},
	{[]string{"check"}, "a: 600\n", 1, "", "retrylane: line 1: bad outcome 600\n"},
	{[]string{"check"}, "a: 20\n", 1, "", "retrylane: line 1: bad outcome 20\n"},
	{[]string{"check"}, "a: 2000\n", 1, "", "retrylane: line 1: bad outcome 2000\n"},
	{[]string{"check"}, "a: 200@5\n", 1, "", "retrylane: line 1: bad outcome 200@5\n"},
	{[]string{"check"}, "a: 429@0\n", 1, "", "retrylane: line 1: bad outcome 429@0\n"},
	{[]string{"check"}, "a: 429@600001\n", 1, "", "retrylane: line 1: bad outcome 429@600001\n"},
	{[]string{"check"}, "a: 429@05\n", 1, "", "retrylane: line 1: bad outcome 429@05\n"},
	{[]string{"check"}, "a: TIMEOUT\n", 1, "", "retrylane: line 1: bad outcome TIMEOUT\n"},
	{[]string{"check"}, "a: 503@3@4\n", 1, "", "retrylane: line 1: bad outcome 503@3@4\n"},
	{[]string{"check"}, "a: 429@\n", 1, "", "retrylane: line 1: bad outcome 429@\n"},
	{[]string{"check"}, "a: 429@600000 503@1 timeout 100 599\n", 0, "ok 1 request\n", ""},
}

func TestRowsOutcomes(t *testing.T) { checkRows(t, rowsOutcomes) }
