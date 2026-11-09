package main

import "testing"

var rowsValues = []row{
	{[]string{"plan", "--base", "0"}, "", 2, "", "retrylane: bad value for --base\n"},
	{[]string{"plan", "--base", "60001"}, "", 2, "", "retrylane: bad value for --base\n"},
	{[]string{"plan", "--factor", "11"}, "", 2, "", "retrylane: bad value for --factor\n"},
	{[]string{"plan", "--tries", "05"}, "", 2, "", "retrylane: bad value for --tries\n"},
	{[]string{"plan", "--deadline", "0"}, "", 2, "", "retrylane: bad value for --deadline\n"},
	{[]string{"breaker", "--call", "60001"}, "", 2, "", "retrylane: bad value for --call\n"},
	{[]string{"plan", "--cap", "-5"}, "", 2, "", "retrylane: bad value for --cap\n"},
	{[]string{"plan", "--base", "60000", "--cap", "600000"}, "a: 503 200\n", 0, "a: ok after 2 attempts, waited 60000 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--tries", "10", "--factor", "10"}, "a: 503 503 503 503 503 503 503 503 503 503\n", 1, "a: gave up after 10 attempts (tries), waited 215500 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
}

func TestRowsValues(t *testing.T) { checkRows(t, rowsValues) }
