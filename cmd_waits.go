package main

import (
	"strconv"
	"strings"
)

// cmdWaits lists only waits committed before a following attempt.
func cmdWaits(opts options, reqs []request) (int, string, string) {
	var out strings.Builder
	waitCount := 0
	totalWaited := 0
	for _, req := range reqs {
		result := replayRequest(req, opts)
		for _, attempt := range result.attempts {
			if attempt.wait == 0 {
				continue
			}
			out.WriteString(req.id)
			out.WriteString(" #")
			out.WriteString(strconv.Itoa(attempt.attempt))
			out.WriteByte(' ')
			out.WriteString(strconv.Itoa(attempt.wait))
			out.WriteString(" ms\n")
			waitCount++
			totalWaited += attempt.wait
		}
	}
	out.WriteString(plural(waitCount, "wait"))
	out.WriteString(", ")
	out.WriteString(strconv.Itoa(totalWaited))
	out.WriteString(" ms total\n")
	return 0, out.String(), ""
}
