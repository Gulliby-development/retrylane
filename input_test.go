package main

import "testing"

var rowsInput = []row{
	{[]string{"check"}, "", 0, "ok 0 requests\n", ""},
	{[]string{"check"}, "\n", 0, "ok 0 requests\n", ""},
	{[]string{"check"}, "# only a comment\n   \n\n", 0, "ok 0 requests\n", ""},
	{[]string{"check"}, "a: 200", 0, "ok 1 request\n", ""},
	{[]string{"check"}, "a: 200\r\nb: 500 200\r\n", 0, "ok 2 requests\n", ""},
	{[]string{"check"}, "a:\t200 \t 500 \n", 0, "ok 1 request\n", ""},
	{[]string{"check"}, "l1: 200\nl2: 200\nl3: 200\nl4: 200\nl5: 200\nl6: 200\nl7: 200\nl8: 200\nl9: 200\nl10: 200\nl11: 200\nl12: 200\nl13: 200\nl14: 200\nl15: 200\nl16: 200\nl17: 200\nl18: 200\nl19: 200\nl20: 200\nl21: 200\n", 1, "", "retrylane: input is too long (max 20 lines)\n"},
	{[]string{"check"}, "# c\n\nbad\n", 1, "", "retrylane: line 3: expected ID: OUTCOME...\n"},
	{[]string{"check"}, "a:200\n", 1, "", "retrylane: line 1: expected ID: OUTCOME...\n"},
}

func TestRowsInput(t *testing.T) { checkRows(t, rowsInput) }
