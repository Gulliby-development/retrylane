package main

import (
	"strconv"
	"strings"
)

// cmdTally summarizes every outcome in the validated log, including outcomes
// beyond the point where a retry replay would stop.
func cmdTally(_ options, reqs []request) (int, string, string) {
	var counts [600]int
	timeouts := 0
	total := 0
	successes := 0
	retryable := 0
	fatal := 0

	for _, req := range reqs {
		for _, result := range req.outcomes {
			total++
			if result.code == 0 {
				timeouts++
			} else {
				counts[result.code]++
			}
			switch classifyOutcome(result) {
			case "success":
				successes++
			case "retryable":
				retryable++
			default:
				fatal++
			}
		}
	}

	var out strings.Builder
	for code := 100; code <= 599; code++ {
		if counts[code] == 0 {
			continue
		}
		out.WriteString(strconv.Itoa(code))
		out.WriteByte(' ')
		out.WriteString(strconv.Itoa(counts[code]))
		out.WriteByte('\n')
	}
	if timeouts != 0 {
		out.WriteString("timeout ")
		out.WriteString(strconv.Itoa(timeouts))
		out.WriteByte('\n')
	}
	out.WriteString(plural(total, "outcome"))
	out.WriteString(" in ")
	out.WriteString(plural(len(reqs), "request"))
	out.WriteByte('\n')
	if total != 0 {
		out.WriteString("success ")
		out.WriteString(strconv.Itoa(successes))
		out.WriteString(", retryable ")
		out.WriteString(strconv.Itoa(retryable))
		out.WriteString(", fatal ")
		out.WriteString(strconv.Itoa(fatal))
		out.WriteByte('\n')
		out.WriteString("retryable share ")
		out.WriteString(share(retryable, total))
		out.WriteByte('\n')
	}
	return 0, out.String(), ""
}
