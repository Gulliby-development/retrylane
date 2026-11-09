package main

import "strings"

// parseCanonicalDecimal accepts an unsigned ASCII decimal without leading
// zeroes and bounds its value before arithmetic can overflow.
func parseCanonicalDecimal(word string, max int) (int, bool) {
	if word == "" || (len(word) > 1 && word[0] == '0') || max < 0 {
		return 0, false
	}
	value := 0
	for i := 0; i < len(word); i++ {
		if word[i] < '0' || word[i] > '9' {
			return 0, false
		}
		digit := int(word[i] - '0')
		if value > max/10 || (value == max/10 && digit > max%10) {
			return 0, false
		}
		value = value*10 + digit
	}
	return value, true
}

func readLog(text string) ([]request, *failure) {
	if text == "" {
		return nil, nil
	}

	lines := strings.Split(text, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) > maxLines {
		return nil, fail(1, "input is too long (max 20 lines)")
	}

	var requests []request
	seen := make(map[string]bool)
	for i, line := range lines {
		if strings.HasSuffix(line, "\r") {
			line = line[:len(line)-1]
		}
		words := splitLogWords(line)
		if len(words) == 0 || strings.HasPrefix(words[0], "#") {
			continue
		}
		lineNumber := i + 1
		if len(words) < 2 || words[0] == ":" || !strings.HasSuffix(words[0], ":") {
			return nil, fail(1, lineError(lineNumber, "expected ID: OUTCOME..."))
		}
		id := words[0][:len(words[0])-1]
		if !validID(id) {
			return nil, fail(1, lineError(lineNumber, "bad id "+id))
		}

		outcomes := make([]outcome, 0, len(words)-1)
		for _, word := range words[1:] {
			parsed, ok := parseOutcome(word)
			if !ok {
				return nil, fail(1, lineError(lineNumber, "bad outcome "+word))
			}
			outcomes = append(outcomes, parsed)
		}
		if len(outcomes) > maxOutcomes {
			return nil, fail(1, lineError(lineNumber, "too many outcomes"))
		}
		if seen[id] {
			return nil, fail(1, lineError(lineNumber, "duplicate request "+id))
		}
		seen[id] = true
		requests = append(requests, request{id: id, line: lineNumber, outcomes: outcomes})
	}
	return requests, nil
}

func splitLogWords(line string) []string {
	var words []string
	for i := 0; i < len(line); {
		for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
			i++
		}
		start := i
		for i < len(line) && line[i] != ' ' && line[i] != '\t' {
			i++
		}
		if start < i {
			words = append(words, line[start:i])
		}
	}
	return words
}

func lineError(line int, reason string) string {
	return "line " + decimalString(line) + ": " + reason
}

func decimalString(value int) string {
	if value == 0 {
		return "0"
	}
	var digits [20]byte
	i := len(digits)
	for value > 0 {
		i--
		digits[i] = byte('0' + value%10)
		value /= 10
	}
	return string(digits[i:])
}

func validID(id string) bool {
	if len(id) < 1 || len(id) > 16 || id[0] < 'a' || id[0] > 'z' {
		return false
	}
	for i := 1; i < len(id); i++ {
		b := id[i]
		if (b < 'a' || b > 'z') && (b < '0' || b > '9') && b != '-' {
			return false
		}
	}
	return true
}

func parseOutcome(word string) (outcome, bool) {
	if word == "timeout" {
		return outcome{code: 0}, true
	}
	if at := strings.IndexByte(word, '@'); at >= 0 {
		codeText, afterText := word[:at], word[at+1:]
		if codeText != "429" && codeText != "503" {
			return outcome{}, false
		}
		after, ok := parseCanonicalDecimal(afterText, maxAfter)
		if !ok || after < 1 {
			return outcome{}, false
		}
		return outcome{code: int(codeText[0]-'0')*100 + int(codeText[1]-'0')*10 + int(codeText[2]-'0'), after: after}, true
	}
	if len(word) != 3 {
		return outcome{}, false
	}
	code := 0
	for i := 0; i < len(word); i++ {
		if word[i] < '0' || word[i] > '9' {
			return outcome{}, false
		}
		code = code*10 + int(word[i]-'0')
	}
	if code < 100 || code > 599 {
		return outcome{}, false
	}
	return outcome{code: code}, true
}
