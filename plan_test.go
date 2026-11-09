package main

import "testing"

var rowsPlan = []row{
	{[]string{"plan"}, "a: 200\n", 0, "a: ok after 1 attempt, waited 0 ms\n1 request: 1 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 404\n", 1, "a: failed at attempt 1 with 404, waited 0 ms\n1 request: 0 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 503 200\nb: 200\nc: 400\n", 1, "a: ok after 2 attempts, waited 500 ms\nb: ok after 1 attempt, waited 0 ms\nc: failed at attempt 1 with 400, waited 0 ms\n3 requests: 2 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 503\n", 1, "a: unfinished after 1 attempt, waited 0 ms\n1 request: 0 ok, 0 failed, 0 gave up, 1 unfinished\n", ""},
	{[]string{"plan"}, "a: 501\n", 1, "a: failed at attempt 1 with 501, waited 0 ms\n1 request: 0 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 429 429 429 429 429\n", 1, "a: gave up after 4 attempts (tries), waited 3500 ms\n1 request: 0 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 300\n", 1, "a: failed at attempt 1 with 300, waited 0 ms\n1 request: 0 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "", 0, "0 requests: 0 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 200 500\nb: 500 200 500\n", 0, "a: ok after 1 attempt, waited 0 ms\nb: ok after 2 attempts, waited 500 ms\n2 requests: 2 ok, 0 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "# a comment\nx1: timeout timeout timeout 200\nx2: 100\n", 1, "x1: ok after 4 attempts, waited 3500 ms\nx2: failed at attempt 1 with 100, waited 0 ms\n2 requests: 1 ok, 1 failed, 0 gave up, 0 unfinished\n", ""},
	{[]string{"plan", "--tries", "3", "--base", "10"}, "a: 502 504 200\nb: 502 504 504 200\n", 1, "a: ok after 3 attempts, waited 30 ms\nb: gave up after 3 attempts (tries), waited 30 ms\n2 requests: 1 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
	{[]string{"plan"}, "a: 503 503 503 503\nb: 204\n", 1, "a: gave up after 4 attempts (tries), waited 3500 ms\nb: ok after 1 attempt, waited 0 ms\n2 requests: 1 ok, 0 failed, 1 gave up, 0 unfinished\n", ""},
}

func TestRowsPlan(t *testing.T) { checkRows(t, rowsPlan) }
