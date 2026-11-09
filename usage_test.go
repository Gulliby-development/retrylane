package main

import "testing"

var rowsUsage = []row{
	{[]string{}, "", 2, "", "usage: retrylane check\n       retrylane plan [POLICY]\n       retrylane waits [POLICY]\n       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]\n       retrylane tally\n"},
	{[]string{"help"}, "", 0, "usage: retrylane check\n       retrylane plan [POLICY]\n       retrylane waits [POLICY]\n       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]\n       retrylane tally\n", ""},
	{[]string{"-h", "extra"}, "", 0, "usage: retrylane check\n       retrylane plan [POLICY]\n       retrylane waits [POLICY]\n       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]\n       retrylane tally\n", ""},
	{[]string{"--help", "--bogus"}, "", 0, "usage: retrylane check\n       retrylane plan [POLICY]\n       retrylane waits [POLICY]\n       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]\n       retrylane tally\n", ""},
	{[]string{"--version"}, "", 0, "retrylane 0.1.0\n", ""},
	{[]string{"--version", "x"}, "", 0, "retrylane 0.1.0\n", ""},
}

func TestRowsUsage(t *testing.T) { checkRows(t, rowsUsage) }
