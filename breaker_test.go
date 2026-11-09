package main

import "testing"

var rowsBreaker = []row{
	{[]string{"breaker"}, "a: 200\n", 0, "final closed at 100 ms, 0 opens\n", ""},
	{[]string{"breaker"}, "", 0, "final closed at 0 ms, 0 opens\n", ""},
	{[]string{"breaker"}, "a: 503 503 503 200\n", 1, "1800 open a\n11800 half-open a\n11900 closed a\nfinal closed at 11900 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "2"}, "a: 503 503 200\n", 1, "700 open a\n10700 half-open a\n10800 closed a\nfinal closed at 10800 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "1"}, "a: 500 200\n", 1, "100 open a\n10100 half-open a\n10200 closed a\nfinal closed at 10200 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "5"}, "a: 500 500 500 500\nb: 500 200\n", 1, "4000 open b\n14000 half-open b\n14100 closed b\nfinal closed at 14100 ms, 1 open\n", ""},
	{[]string{"breaker"}, "a: 400\nb: 404\n", 0, "final closed at 200 ms, 0 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--tries", "1"}, "a: 500\nb: 500\nc: 200\n", 1, "200 open b\n10200 half-open c\n10300 closed c\nfinal closed at 10300 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "2", "--call", "0", "--cooldown", "1"}, "a: 500 500 200\n", 1, "500 open a\n1500 half-open a\n1500 closed a\nfinal closed at 1500 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "3"}, "a: 500 500 400 500 500 200\n", 0, "final closed at 1800 ms, 0 opens\n", ""},
	{[]string{"breaker"}, "a: timeout timeout timeout\n", 1, "1800 open a\nfinal open at 1800 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "100000"}, "a: 503 503 200\nb: 200\n", 1, "700 open a\n100700 half-open a\n100800 closed a\nfinal closed at 100900 ms, 1 open\n", ""},
}

func TestRowsBreaker(t *testing.T) { checkRows(t, rowsBreaker) }
