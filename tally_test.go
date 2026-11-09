package main

import "testing"

var rowsTally = []row{
	{[]string{"tally"}, "", 0, "0 outcomes in 0 requests\n", ""},
	{[]string{"tally"}, "a: 200\n", 0, "200 1\n1 outcome in 1 request\nsuccess 1, retryable 0, fatal 0\nretryable share 0.0%\n", ""},
	{[]string{"tally"}, "a: 503 200\nb: timeout\n", 0, "200 1\n503 1\ntimeout 1\n3 outcomes in 2 requests\nsuccess 1, retryable 2, fatal 0\nretryable share 66.7%\n", ""},
	{[]string{"tally"}, "a: 200 200 200 500 500 timeout\n", 0, "200 3\n500 2\ntimeout 1\n6 outcomes in 1 request\nsuccess 3, retryable 3, fatal 0\nretryable share 50.0%\n", ""},
	{[]string{"tally"}, "a: 404 404 401\nb: 429@5 200\n", 0, "200 1\n401 1\n404 2\n429 1\n5 outcomes in 2 requests\nsuccess 1, retryable 1, fatal 3\nretryable share 20.0%\n", ""},
	{[]string{"tally"}, "a: 500\nb: 500\nc: 500\n", 0, "500 3\n3 outcomes in 3 requests\nsuccess 0, retryable 3, fatal 0\nretryable share 100.0%\n", ""},
	{[]string{"tally"}, "a: 200 200 500\n", 0, "200 2\n500 1\n3 outcomes in 1 request\nsuccess 2, retryable 1, fatal 0\nretryable share 33.3%\n", ""},
	{[]string{"tally"}, "a: 200 200 200 200 200 200 200 500\n", 0, "200 7\n500 1\n8 outcomes in 1 request\nsuccess 7, retryable 1, fatal 0\nretryable share 12.5%\n", ""},
	{[]string{"tally"}, "a: 100 301 599\n", 0, "100 1\n301 1\n599 1\n3 outcomes in 1 request\nsuccess 0, retryable 0, fatal 3\nretryable share 0.0%\n", ""},
	{[]string{"tally"}, "a: 200 500 500 500 500 500 500 500 500 500 500 500\n", 0, "200 1\n500 11\n12 outcomes in 1 request\nsuccess 1, retryable 11, fatal 0\nretryable share 91.7%\n", ""},
}

func TestRowsTally(t *testing.T) { checkRows(t, rowsTally) }
