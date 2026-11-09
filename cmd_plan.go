package main

import (
	"strconv"
	"strings"
)

// cmdPlan replays each request and reports its terminal outcome and totals.
func cmdPlan(opts options, reqs []request) (int, string, string) {
	var output strings.Builder
	okCount, failedCount, gaveUpCount, unfinishedCount := 0, 0, 0, 0

	for _, req := range reqs {
		result := replayRequest(req, opts)
		last := result.attempts[len(result.attempts)-1]
		attempts := plural(last.attempt, "attempt")
		output.WriteString(req.id)
		output.WriteString(": ")
		switch result.kind {
		case "ok":
			okCount++
			output.WriteString("ok after ")
			output.WriteString(attempts)
		case "failed":
			failedCount++
			output.WriteString("failed at attempt ")
			output.WriteString(strconv.Itoa(last.attempt))
			output.WriteString(" with ")
			output.WriteString(strconv.Itoa(result.status))
		case "gave up":
			gaveUpCount++
			output.WriteString("gave up after ")
			output.WriteString(attempts)
			output.WriteString(" (")
			output.WriteString(result.reason)
			output.WriteByte(')')
		case "unfinished":
			unfinishedCount++
			output.WriteString("unfinished after ")
			output.WriteString(attempts)
		}
		output.WriteString(", waited ")
		output.WriteString(strconv.Itoa(result.waited))
		output.WriteString(" ms\n")
	}

	output.WriteString(plural(len(reqs), "request"))
	output.WriteString(": ")
	output.WriteString(strconv.Itoa(okCount))
	output.WriteString(" ok, ")
	output.WriteString(strconv.Itoa(failedCount))
	output.WriteString(" failed, ")
	output.WriteString(strconv.Itoa(gaveUpCount))
	output.WriteString(" gave up, ")
	output.WriteString(strconv.Itoa(unfinishedCount))
	output.WriteString(" unfinished\n")

	if okCount == len(reqs) {
		return 0, output.String(), ""
	}
	return 1, output.String(), ""
}
