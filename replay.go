package main

// replayAttempt records one consumed outcome and the wait that follows it.
// A nonzero wait is present only when replay proceeds to another attempt.
type replayAttempt struct {
	outcome outcome
	attempt int
	wait    int
}

// replayResult is the consumed prefix and terminal disposition of a request.
type replayResult struct {
	attempts []replayAttempt
	waited   int
	kind     string
	reason   string
	status   int
}

// classifyOutcome classifies a validated result independently of replay.
func classifyOutcome(result outcome) string {
	if result.code >= 200 && result.code <= 299 {
		return "success"
	}
	switch result.code {
	case 0, 429, 500, 502, 503, 504:
		return "retryable"
	default:
		return "fatal"
	}
}

// replayRequest applies the retry policy to a request without changing it or
// the supplied policy options.
func replayRequest(req request, opts options) replayResult {
	result := replayResult{}
	base := opts["--base"]
	factor := opts["--factor"]
	capWait := opts["--cap"]
	tries := opts["--tries"]
	deadline := opts["--deadline"]
	delay := base
	if delay > capWait {
		delay = capWait
	}

	for attempt, index := 1, 0; ; attempt, index = attempt+1, index+1 {
		current := replayAttempt{outcome: req.outcomes[index], attempt: attempt}
		result.attempts = append(result.attempts, current)
		class := classifyOutcome(current.outcome)
		if class == "success" {
			result.kind = "ok"
			return result
		}
		if class == "fatal" {
			result.kind = "failed"
			result.status = current.outcome.code
			return result
		}
		if attempt == tries {
			result.kind = "gave up"
			result.reason = "tries"
			return result
		}
		if current.outcome.after > capWait {
			result.kind = "gave up"
			result.reason = "retry-after"
			return result
		}

		candidate := delay
		if current.outcome.after > candidate {
			candidate = current.outcome.after
		}
		if deadline != 0 && result.waited+candidate > deadline {
			result.kind = "gave up"
			result.reason = "deadline"
			return result
		}
		if index+1 >= len(req.outcomes) {
			result.kind = "unfinished"
			return result
		}

		result.attempts[len(result.attempts)-1].wait = candidate
		result.waited += candidate
		if delay < capWait {
			if delay > capWait/factor {
				delay = capWait
			} else {
				delay *= factor
			}
			if delay > capWait {
				delay = capWait
			}
		}
	}
}
