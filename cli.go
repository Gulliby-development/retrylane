package main

// readsStdin reports whether args is a complete, valid invocation of a command
// that consumes the log. It deliberately does not depend on handler readiness.
func readsStdin(args []string) bool {
	if _, _, ok := specialReply(args); ok {
		return false
	}
	_, _, err := validateCommandLine(args)
	return err == nil
}

func run(args []string, text string) (int, string, string) {
	if status, output, ok := specialReply(args); ok {
		if len(args) == 0 {
			return status, "", output
		}
		return status, output, ""
	}

	cmd, opts, cliFailure := validateCommandLine(args)
	if cliFailure != nil {
		return renderFailure(cliFailure)
	}
	reqs, inputFailure := readLog(text)
	if inputFailure != nil {
		return renderFailure(inputFailure)
	}
	return cmd.run(opts, reqs)
}

func specialReply(args []string) (int, string, bool) {
	if len(args) == 0 {
		return 2, usage, true
	}
	switch args[0] {
	case "help", "-h", "--help":
		return 0, usage, true
	case "--version":
		return 0, version + "\n", true
	default:
		return 0, "", false
	}
}

func validateCommandLine(args []string) (*command, options, *failure) {
	if len(args) == 0 {
		return nil, nil, fail(2, usage)
	}
	cmd := findCommand(args[0])
	if cmd == nil {
		return nil, nil, fail(2, "unknown command "+args[0])
	}

	opts := make(options, len(optionSpecs))
	for _, spec := range optionSpecs {
		opts[spec.name] = spec.fallback
	}

	seen := make(map[string]bool)
	for i := 1; i < len(args); i++ {
		word := args[i]
		if word == "" || word == "-" || word[0] != '-' {
			return nil, nil, fail(2, "unexpected argument "+word)
		}
		spec := findSpec(word)
		if spec == nil {
			return nil, nil, fail(2, "unknown option "+word)
		}
		if !acceptsOption(cmd, word) {
			return nil, nil, fail(2, "option "+word+" is not valid for "+cmd.name)
		}
		if seen[word] {
			return nil, nil, fail(2, "given twice: "+word)
		}
		seen[word] = true
		if i+1 == len(args) {
			return nil, nil, fail(2, "missing value for "+word)
		}
		value, ok := parseCanonicalDecimal(args[i+1], spec.high)
		if !ok || value < spec.low {
			return nil, nil, fail(2, "bad value for "+word)
		}
		opts[word] = value
		i++
	}
	return cmd, opts, nil
}

func acceptsOption(cmd *command, name string) bool {
	for _, allowed := range cmd.options {
		if allowed == name {
			return true
		}
	}
	return false
}

func renderFailure(f *failure) (int, string, string) {
	message := f.msg
	if message == usage {
		return f.status, "", message
	}
	return f.status, "", "retrylane: " + message + "\n"
}
