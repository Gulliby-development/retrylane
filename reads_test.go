package main

import "testing"

var rowsReads = []row{
	{[]string{}, "", 0, "", ""},
	{[]string{"check"}, "", 1, "", ""},
	{[]string{"tally"}, "", 1, "", ""},
	{[]string{"plan", "--base", "5"}, "", 1, "", ""},
	{[]string{"plan", "--base"}, "", 0, "", ""},
	{[]string{"help"}, "", 0, "", ""},
	{[]string{"breaker", "--call", "0"}, "", 1, "", ""},
	{[]string{"check", "--base", "5"}, "", 0, "", ""},
	{[]string{"waits", "x"}, "", 0, "", ""},
	{[]string{"--version"}, "", 0, "", ""},
}

func TestRowsReads(t *testing.T) { checkReadsRows(t, rowsReads) }
