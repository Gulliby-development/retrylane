package main

import "strconv"

// plural formats a count with the noun's plural form unless the count is one.
func plural(count int, noun string) string {
	if count != 1 {
		noun += "s"
	}
	return strconv.Itoa(count) + " " + noun
}

// share formats part as a percentage of whole, rounded half up to one decimal.
// The supported values are nonnegative and bounded by the validated log size.
func share(part, whole int) string {
	if whole == 0 {
		return "0.0%"
	}
	tenths := (part*1000 + whole/2) / whole
	return strconv.Itoa(tenths/10) + "." + strconv.Itoa(tenths%10) + "%"
}
