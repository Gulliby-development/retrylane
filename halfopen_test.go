package main

import "testing"

var rowsHalfopen = []row{
	{[]string{"breaker", "--open", "2", "--cooldown", "500"}, "a: 503 503 503 200\n", 1, "700 open a\n1700 half-open a\n1800 open a\n3800 half-open a\n3900 closed a\nfinal closed at 3900 ms, 2 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "100"}, "a: 503 503 503 503 200\n", 1, "700 open a\n1700 half-open a\n1800 open a\n3800 half-open a\n3900 open a\nfinal open at 3900 ms, 3 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "100", "--base", "1000"}, "a: 503 503 200\n", 1, "1200 open a\n3200 half-open a\n3300 closed a\nfinal closed at 3300 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "10000", "--call", "0"}, "a: 503 503 404\n", 1, "500 open a\n10500 half-open a\n10500 closed a\nfinal closed at 10500 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "1", "--cooldown", "1"}, "a: 500 500 500 200\n", 1, "100 open a\n600 half-open a\n700 open a\n1700 half-open a\n1800 open a\n3800 half-open a\n3900 closed a\nfinal closed at 3900 ms, 3 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--tries", "10"}, "a: 500 500 500 500 500 500 200\n", 1, "700 open a\n10700 half-open a\n10800 open a\n20800 half-open a\n20900 open a\n30900 half-open a\n31000 open a\n41000 half-open a\n41100 open a\n57100 half-open a\n57200 closed a\nfinal closed at 57200 ms, 5 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--tries", "2", "--cooldown", "50"}, "a: 500 500\nb: 500 500\nc: 200\n", 1, "700 open a\n750 half-open b\n850 open b\n1350 half-open b\n1450 open b\n1500 half-open c\n1600 closed c\nfinal closed at 1600 ms, 3 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "700", "--call", "50"}, "a: 429@3000 503 200\n", 1, "3100 open a\n4100 half-open a\n4150 closed a\nfinal closed at 4150 ms, 1 open\n", ""},
	{[]string{"breaker", "--open", "3", "--cooldown", "1"}, "a: 500 500 200\nb: 500 500 200\n", 0, "final closed at 3600 ms, 0 opens\n", ""},
	{[]string{"breaker", "--open", "2", "--cooldown", "9000"}, "a: 500 500 500\n", 1, "700 open a\n9700 half-open a\n9800 open a\nfinal open at 9800 ms, 2 opens\n", ""},
}

func TestRowsHalfopen(t *testing.T) { checkRows(t, rowsHalfopen) }
