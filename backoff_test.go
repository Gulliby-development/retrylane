package main

import "testing"

var rowsBackoff = []row{
	{[]string{"waits"}, "a: 503 503 503 200\n", 0, "a #1 500 ms\na #2 1000 ms\na #3 2000 ms\n3 waits, 3500 ms total\n", ""},
	{[]string{"waits", "--base", "100", "--factor", "3"}, "a: 500 500 500 500\n", 0, "a #1 100 ms\na #2 300 ms\na #3 900 ms\n3 waits, 1300 ms total\n", ""},
	{[]string{"waits", "--factor", "1"}, "a: 503 503 503 200\n", 0, "a #1 500 ms\na #2 500 ms\na #3 500 ms\n3 waits, 1500 ms total\n", ""},
	{[]string{"waits", "--cap", "700"}, "a: 503 503 503 200\n", 0, "a #1 500 ms\na #2 700 ms\na #3 700 ms\n3 waits, 1900 ms total\n", ""},
	{[]string{"waits", "--tries", "10", "--base", "1000", "--factor", "10", "--cap", "600000"}, "a: 500 500 500 500 500 500 500 500 500 200\n", 0, "a #1 1000 ms\na #2 10000 ms\na #3 100000 ms\na #4 600000 ms\na #5 600000 ms\na #6 600000 ms\na #7 600000 ms\na #8 600000 ms\na #9 600000 ms\n9 waits, 3711000 ms total\n", ""},
	{[]string{"plan", "--tries", "1"}, "a: 503\n", 1, "a: gave up after 1 attempt (tries), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--tries", "2"}, "a: 500 200\n", 0, "a: ok after 2 attempts, waited 500 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"waits", "--base", "60000", "--factor", "10", "--cap", "600000", "--tries", "5"}, "a: 502 502 502 502 200\n", 0, "a #1 60000 ms\na #2 600000 ms\na #3 600000 ms\na #4 600000 ms\n4 waits, 1860000 ms total\n", ""},
	{[]string{"waits"}, "a: timeout timeout 200\n", 0, "a #1 500 ms\na #2 1000 ms\n2 waits, 1500 ms total\n", ""},
	{[]string{"waits", "--base", "800", "--cap", "300"}, "a: 504 504 200\n", 0, "a #1 300 ms\na #2 300 ms\n2 waits, 600 ms total\n", ""},
	{[]string{"waits", "--tries", "3"}, "a: 503 503 503 503\n", 0, "a #1 500 ms\na #2 1000 ms\n2 waits, 1500 ms total\n", ""},
	{[]string{"plan", "--base", "250", "--factor", "4", "--tries", "6"}, "a: 500 500 500 500 500 200\n", 0, "a: ok after 6 attempts, waited 51250 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
}

func TestRowsBackoff(t *testing.T) { checkRows(t, rowsBackoff) }
