package main

import (
	"strconv"
	"strings"
)

// cmdBreaker replays consumed retry attempts through one virtual-clock breaker.
func cmdBreaker(opts options, reqs []request) (int, string, string) {
	var out strings.Builder
	state := "closed"
	now := 0
	openedAt := 0
	fails := 0
	opens := 0
	threshold := opts["--open"]
	cooldown := opts["--cooldown"]
	call := opts["--call"]

	for _, req := range reqs {
		result := replayRequest(req, opts)
		for _, attempt := range result.attempts {
			if state == "open" {
				if ready := openedAt + cooldown; ready > now {
					now = ready
				}
				state = "half-open"
				writeBreakerTransition(&out, now, state, req.id)
			}

			now += call
			if classifyOutcome(attempt.outcome) == "retryable" {
				if state == "half-open" {
					state = "open"
					openedAt = now
					opens++
					writeBreakerTransition(&out, now, state, req.id)
				} else {
					fails++
					if fails >= threshold {
						state = "open"
						openedAt = now
						fails = 0
						opens++
						writeBreakerTransition(&out, now, state, req.id)
					}
				}
			} else {
				fails = 0
				if state == "half-open" {
					state = "closed"
					writeBreakerTransition(&out, now, state, req.id)
				}
			}
			now += attempt.wait
		}
	}

	out.WriteString("final ")
	out.WriteString(state)
	out.WriteString(" at ")
	out.WriteString(strconv.Itoa(now))
	out.WriteString(" ms, ")
	out.WriteString(plural(opens, "open"))
	out.WriteByte('\n')
	status := 0
	if opens > 0 {
		status = 1
	}
	return status, out.String(), ""
}

func writeBreakerTransition(out *strings.Builder, now int, state, id string) {
	out.WriteString(strconv.Itoa(now))
	out.WriteByte(' ')
	out.WriteString(state)
	out.WriteByte(' ')
	out.WriteString(id)
	out.WriteByte('\n')
}
