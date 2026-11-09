package main

import "testing"

var rowsEdges = []row{
	{[]string{"plan", "--tries", "1", "--deadline", "1"}, "a: 503@1\n", 1, "a: gave up after 1 attempt (tries), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--base", "1", "--factor", "10", "--cap", "5", "--tries", "10"}, "a: 500 500 500 500 500 500 500 500 500 200\n", 0, "a: ok after 10 attempts, waited 41 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"waits", "--base", "60000", "--cap", "600000", "--factor", "10", "--tries", "10"}, "a: 500 500 500 500 500 500 500 500 500 200\n", 0, "a #1 60000 ms\na #2 600000 ms\na #3 600000 ms\na #4 600000 ms\na #5 600000 ms\na #6 600000 ms\na #7 600000 ms\na #8 600000 ms\na #9 600000 ms\n9 waits, 4860000 ms total\n", ""},
	{[]string{"plan", "--deadline", "3600000", "--cap", "600000", "--tries", "3"}, "a: 503@600000 503@600000 200\n", 0, "a: ok after 3 attempts, waited 1200000 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"tally"}, "a: 503@600000 429@1 timeout\n", 0, "429 1\n503 1\ntimeout 1\n3 outcomes in 1 request\nsuccess 0, retryable 3, fatal 0\nretryable share 100.0%\n", ""},
	{[]string{"breaker", "--open", "20", "--cooldown", "600000", "--call", "60000"}, "a: 500 500 200\n", 0, "final closed at 181500 ms, 0 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--call", "60000", "--cooldown", "600000", "--base", "60000", "--factor", "10", "--cap", "600000"}, "a: 500 500 500 200\n", 1, "180000 open a\n780000 half-open a\n840000 open a\n1440000 half-open a\n1500000 closed a\nfinal closed at 1500000 ms, 2 opens\n", ""},
	{[]string{"tally"}, "a: 200\n# gap\n\nb: 200 200\n", 0, "200 3\n3 outcomes in 2 requests\nsuccess 3, retryable 0, fatal 0\nretryable share 0.0%\n", ""},
}

func TestRowsEdges(t *testing.T) { checkRows(t, rowsEdges) }
