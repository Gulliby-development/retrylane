package main

import "testing"

var rowsAfterwait = []row{
	{[]string{"waits"}, "a: 429@2000 200\n", 0, "a #1 2000 ms\n1 wait, 2000 ms total\n", ""},
	{[]string{"waits"}, "a: 429@100 200\n", 0, "a #1 500 ms\n1 wait, 500 ms total\n", ""},
	{[]string{"waits"}, "a: 503@30000 200\n", 0, "a #1 30000 ms\n1 wait, 30000 ms total\n", ""},
	{[]string{"plan"}, "a: 503@30001 200\n", 1, "a: gave up after 1 attempt (retry-after), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--cap", "1000"}, "a: 429@1001 200\n", 1, "a: gave up after 1 attempt (retry-after), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"waits"}, "a: 429@800 503 200\n", 0, "a #1 800 ms\na #2 1000 ms\n2 waits, 1800 ms total\n", ""},
	{[]string{"plan"}, "a: 503@900\n", 1, "a: unfinished after 1 attempt, waited 0 ms\n1 request: 0 ok, 0 failed, 0 gave up, 1 unfinished\n", ""},
	{[]string{"waits", "--cap", "600000"}, "a: 429@600000 200\n", 0, "a #1 600000 ms\n1 wait, 600000 ms total\n", ""},
	{[]string{"plan", "--tries", "2"}, "a: 503@99999 503@99999\n", 1, "a: gave up after 1 attempt (retry-after), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
}

func TestRowsAfterwait(t *testing.T) { checkRows(t, rowsAfterwait) }
