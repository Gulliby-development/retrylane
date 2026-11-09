package main

import "testing"

// row is one literal test row: the arguments, the standard input, and the
// expected exit status, standard output and standard error.
type row struct {
	args   []string
	in     string
	status int
	out    string
	err    string
}

func checkRows(t *testing.T, rows []row) {
	t.Helper()
	for i, r := range rows {
		status, out, errText := run(r.args, r.in)
		if status != r.status || out != r.out || errText != r.err {
			t.Errorf("row %d %q: got (%d, %q, %q), want (%d, %q, %q)", i+1, r.args, status, out, errText, r.status, r.out, r.err)
		}
	}
}

// checkReadsRows reads the status of a row as a flag: 1 when readsStdin must be
// true for the arguments, 0 when it must be false; the other fields are unused.
func checkReadsRows(t *testing.T, rows []row) {
	t.Helper()
	for i, r := range rows {
		if got := readsStdin(r.args); got != (r.status == 1) {
			t.Errorf("row %d %q: readsStdin is %v, want %v", i+1, r.args, got, r.status == 1)
		}
	}
}
