package main

import "testing"

var rowsOptions = []row{
	{[]string{"check", "--base", "5"}, "", 2, "", "retrylane: option --base is not valid for check\n"},
	{[]string{"tally", "--cap", "5"}, "", 2, "", "retrylane: option --cap is not valid for tally\n"},
	{[]string{"plan", "--open", "3"}, "", 2, "", "retrylane: option --open is not valid for plan\n"},
	{[]string{"waits", "--call", "1"}, "", 2, "", "retrylane: option --call is not valid for waits\n"},
	{[]string{"plan", "--base", "5", "--base", "6"}, "", 2, "", "retrylane: given twice: --base\n"},
	{[]string{"plan", "--tries"}, "", 2, "", "retrylane: missing value for --tries\n"},
	{[]string{"plan", "stray"}, "", 2, "", "retrylane: unexpected argument stray\n"},
	{[]string{"plan", "--nope", "1"}, "", 2, "", "retrylane: unknown option --nope\n"},
	{[]string{"plan", "--base", "--base"}, "", 2, "", "retrylane: bad value for --base\n"},
	{[]string{"breaker", "--open", "2", "--cooldown", "500"}, "a: 503 503 200\n", 1, "700 open a\n1700 half-open a\n1800 closed a\nfinal closed at 1800 ms, 1 open\n", ""},
	{[]string{"breaker", "--call", "0"}, "a: 500 500 500 200\n", 1, "1500 open a\n11500 half-open a\n11500 closed a\nfinal closed at 11500 ms, 1 open\n", ""},
}

func TestRowsOptions(t *testing.T) { checkRows(t, rowsOptions) }
