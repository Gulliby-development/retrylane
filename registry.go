package main

const (
	version     = "retrylane 0.1.0"
	maxLines    = 20
	maxOutcomes = 12
	maxAfter    = 600000
)

const usage = "usage: retrylane check\n" +
	"       retrylane plan [POLICY]\n" +
	"       retrylane waits [POLICY]\n" +
	"       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]\n" +
	"       retrylane tally\n"

// failure is an error that carries the exit status of the program: 2 for the
// command line, 1 for the input. fail builds one; run turns it into the reply
// (status, "", "retrylane: " + msg + "\n").
type failure struct {
	status int
	msg    string
}

func (f *failure) Error() string { return f.msg }

func fail(status int, msg string) *failure { return &failure{status: status, msg: msg} }

// optionSpec describes one option: its word, the lowest and the highest value
// and the default (0 for --deadline means no deadline).
type optionSpec struct {
	name     string
	low      int
	high     int
	fallback int
}

var optionSpecs = []optionSpec{
	{"--base", 1, 60000, 500},
	{"--factor", 1, 10, 2},
	{"--cap", 1, 600000, 30000},
	{"--tries", 1, 10, 4},
	{"--deadline", 1, 3600000, 0},
	{"--open", 1, 20, 3},
	{"--cooldown", 1, 600000, 10000},
	{"--call", 0, 60000, 100},
}

// options maps an option word to its value; run fills in every default.
type options map[string]int

// outcome is one result of a call: the HTTP status, or 0 for a timeout, and the
// Retry-After wait in milliseconds (0 when the outcome names none).
type outcome struct {
	code  int
	after int
}

// request is one meaningful line of the input: the id, the line number (counted
// from 1, skipped lines included) and the outcomes of its calls in order.
type request struct {
	id       string
	line     int
	outcomes []outcome
}

// command is one entry of the command table; run returns the exit status, the
// whole standard output and the standard error of the command.
type command struct {
	name    string
	options []string
	run     func(opts options, reqs []request) (int, string, string)
}

var policyOptions = []string{"--base", "--factor", "--cap", "--tries", "--deadline"}

var commands = []command{
	{"check", nil, cmdCheck},
	{"plan", policyOptions, cmdPlan},
	{"waits", policyOptions, cmdWaits},
	{"breaker", append(append([]string{}, policyOptions...), "--open", "--cooldown", "--call"), cmdBreaker},
	{"tally", nil, cmdTally},
}

func findCommand(name string) *command {
	for i := range commands {
		if commands[i].name == name {
			return &commands[i]
		}
	}
	return nil
}

func findSpec(name string) *optionSpec {
	for i := range optionSpecs {
		if optionSpecs[i].name == name {
			return &optionSpecs[i]
		}
	}
	return nil
}
