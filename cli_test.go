package main

import "testing"

var rowsCli = []row{
	{[]string{"frob"}, "", 2, "", "retrylane: unknown command frob\n"},
	{[]string{""}, "", 2, "", "retrylane: unknown command \n"},
	{[]string{"Check"}, "", 2, "", "retrylane: unknown command Check\n"},
	{[]string{"check"}, "a: 200\nb: 503 200\n", 0, "ok 2 requests\n", ""},
	{[]string{"check", "extra"}, "", 2, "", "retrylane: unexpected argument extra\n"},
	{[]string{"plan"}, "a: 503 200\nb: 404\n", 1, "a: ok after 2 attempts, waited 500 ms\nb: failed at attempt 1 with 404, waited 0 ms\n2 requests: 1 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"waits"}, "a: 503 500 200\n", 0, "a #1 500 ms\na #2 1000 ms\n2 waits, 1500 ms total\n", ""},
	{[]string{"breaker"}, "a: 503 503 503 200\n", 1, "1800 open a\n11800 half-open a\n11900 closed a\nfinal closed at 11900 ms, 1 open\n", ""},
	{[]string{"tally"}, "a: 503 200\nb: timeout\n", 0, "200 1\n503 1\ntimeout 1\n3 outcomes in 2 requests\nsuccess 1, retryable 2, fatal 0\nretryable share 66.7%\n", ""},
}

func TestRowsCli(t *testing.T) { checkRows(t, rowsCli) }
