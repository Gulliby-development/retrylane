package main

import "testing"

var rowsNames = []row{
	{[]string{"check"}, "A: 200\n", 1, "", "retrylane: line 1: bad id A\n"},
	{[]string{"check"}, "1a: 200\n", 1, "", "retrylane: line 1: bad id 1a\n"},
	{[]string{"check"}, "a_b: 200\n", 1, "", "retrylane: line 1: bad id a_b\n"},
	{[]string{"check"}, "abcdefghijklmnopq: 200\n", 1, "", "retrylane: line 1: bad id abcdefghijklmnopq\n"},
	{[]string{"check"}, "a-1: 200\nabcdefghijklmnop: 500\n", 0, "ok 2 requests\n", ""},
	{[]string{"check"}, "ab: 200\nab: 500\n", 1, "", "retrylane: line 2: duplicate request ab\n"},
	{[]string{"check"}, ": 200\n", 1, "", "retrylane: line 1: expected ID: OUTCOME...\n"},
	{[]string{"check"}, "a : 200\n", 1, "", "retrylane: line 1: expected ID: OUTCOME...\n"},
}

func TestRowsNames(t *testing.T) { checkRows(t, rowsNames) }
