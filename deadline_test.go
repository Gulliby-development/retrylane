package main

import "testing"

var rowsDeadline = []row{
	{[]string{"plan", "--deadline", "1000"}, "a: 500 500 200\n", 1, "a: gave up after 2 attempts (deadline), waited 500 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--deadline", "1500"}, "a: 500 500 200\n", 0, "a: ok after 3 attempts, waited 1500 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--deadline", "1"}, "a: 503 200\n", 1, "a: gave up after 1 attempt (deadline), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--deadline", "100000"}, "a: 429@40000 200\n", 1, "a: gave up after 1 attempt (retry-after), waited 0 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"waits", "--deadline", "2000"}, "a: 500 500 500 200\nb: 500 200\n", 0, "a #1 500 ms\na #2 1000 ms\nb #1 500 ms\n3 waits, 2000 ms total\n", ""},
	{[]string{"plan", "--deadline", "700", "--tries", "2"}, "a: 500 500\n", 1, "a: gave up after 2 attempts (tries), waited 500 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"breaker", "--deadline", "600", "--open", "20"}, "a: 500 500 200\n", 0, "final closed at 700 ms, 0 opens\n", ""},
	{[]string{"breaker", "--deadline", "500", "--open", "2"}, "a: 503 200\nb: 503 200\n", 0, "final closed at 1400 ms, 0 opens\n", ""},
	{[]string{"breaker", "--deadline", "100000", "--open", "20"}, "a: 503 503 503 200\n", 0, "final closed at 3900 ms, 0 opens\n", ""},
}

func TestRowsDeadline(t *testing.T) { checkRows(t, rowsDeadline) }
