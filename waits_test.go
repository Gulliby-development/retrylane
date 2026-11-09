package main

import "testing"

var rowsWaits = []row{
	{[]string{"waits"}, "a: 200\n", 0, "0 waits, 0 ms total\n", ""},
	{[]string{"waits"}, "a: 503 200\n", 0, "a #1 500 ms\n1 wait, 500 ms total\n", ""},
	{[]string{"waits"}, "a: 503 200\nb: 500 500 200\n", 0, "a #1 500 ms\nb #1 500 ms\nb #2 1000 ms\n3 waits, 2000 ms total\n", ""},
	{[]string{"waits"}, "", 0, "0 waits, 0 ms total\n", ""},
	{[]string{"waits"}, "a: 404\n", 0, "0 waits, 0 ms total\n", ""},
	{[]string{"waits", "--factor", "3", "--base", "20"}, "a: 500 500 500 200\nb: 429@100 200\n", 0, "a #1 20 ms\na #2 60 ms\na #3 180 ms\nb #1 100 ms\n4 waits, 360 ms total\n", ""},
	{[]string{"waits"}, "a: 503\n", 0, "0 waits, 0 ms total\n", ""},
	{[]string{"waits", "--tries", "2"}, "a: 503 503 200\n", 0, "a #1 500 ms\n1 wait, 500 ms total\n", ""},
	{[]string{"waits", "--base", "1"}, "one-1: 500 500 500 200\n", 0, "one-1 #1 1 ms\none-1 #2 2 ms\none-1 #3 4 ms\n3 waits, 7 ms total\n", ""},
	{[]string{"waits"}, "a: 429@1500 429 429@5000 200\n", 0, "a #1 1500 ms\na #2 1000 ms\na #3 5000 ms\n3 waits, 7500 ms total\n", ""},
}

func TestRowsWaits(t *testing.T) { checkRows(t, rowsWaits) }
