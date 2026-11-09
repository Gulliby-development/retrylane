package main

// cmdCheck reports how many requests were validated in the input log.
func cmdCheck(_ options, reqs []request) (int, string, string) {
	return 0, "ok " + plural(len(reqs), "request") + "\n", ""
}
