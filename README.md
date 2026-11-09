<p align="center"><img src="docs/logo.svg" width="96" alt="" /></p>

# retrylane

[![CI](https://github.com/Gulliby-development/retrylane/actions/workflows/ci.yml/badge.svg)](https://github.com/Gulliby-development/retrylane/actions/workflows/ci.yml) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

retrylane is a deterministic, standard-library-only Go command-line simulator for people building clients for language-model APIs. Use it to replay a call log against retry settings and inspect waits, outcomes, and circuit-breaker behavior before changing a client. It makes no network calls, never sleeps, reads no clock, and uses no randomness: the same log and options always produce the same result. It runs with Go 1.27 on Windows, macOS, and Linux.

## Usage

Run commands from the repository directory with `go run . COMMAND`. Each command reads its log from standard input. In the POSIX-shell examples below, `retrylane` stands for this shell function:

```sh
retrylane() { go run . "$@"; }
```

Synopsis:

```text
usage: retrylane check
       retrylane plan [POLICY]
       retrylane waits [POLICY]
       retrylane breaker [POLICY] [--open N] [--cooldown MS] [--call MS]
       retrylane tally
```

`POLICY` consists of any of these options, in any order:

| Option | Range | Default | Commands |
| --- | ---: | ---: | --- |
| `--base MS` | 1–60000 | 500 | `plan`, `waits`, `breaker` |
| `--factor N` | 1–10 | 2 | `plan`, `waits`, `breaker` |
| `--cap MS` | 1–600000 | 30000 | `plan`, `waits`, `breaker` |
| `--tries N` | 1–10 | 4 | `plan`, `waits`, `breaker` |
| `--deadline MS` | 1–3600000 | none | `plan`, `waits`, `breaker` |
| `--open N` | 1–20 | 3 | `breaker` |
| `--cooldown MS` | 1–600000 | 10000 | `breaker` |
| `--call MS` | 0–60000 | 100 | `breaker` |

Values must be unsigned decimal digits, have no leading zero (except the value 0), and be within the listed range. The full command line is validated from left to right before input is read. For each word after the command, a word that does not start with `-` (including an empty word), or is exactly `-`, is an unexpected argument; an unrecognized word beginning with `-` is an unknown option. A recognized option that the command does not accept is reported as invalid for that command. For an accepted option, checks occur in this order: a duplicate option, a missing value when it is the final word, then a bad value. Options take the next word as their value, even if it looks like an option. `check` and `tally` do not accept options.

The first word `help`, `-h`, or `--help` prints the synopsis and exits 0, ignoring later words. `--version` prints `retrylane 0.1.0` and exits 0, also ignoring later words. With no command, the synopsis goes to standard error and the exit status is 2. Other invalid command lines also exit 2. Any other first word, including an empty word, is an unknown command; command names are case-sensitive.

## Call log

Provide one request per line on standard input. A request contains an ID ending in a colon, followed by one or more call outcomes, for example `chat: 503 503@2000 200`. IDs start with a lowercase letter and contain at most 16 lowercase letters, digits, or hyphens. Outcomes are `timeout`, a three-digit status from 100 through 599, or status 429 or 503 followed by `@` and a Retry-After value in milliseconds from 1 through 600000. Retry-After values use decimal digits only, with no sign or leading zero; the text is split at its first `@`, so `503@3@4` is invalid. A status with Retry-After must be 429 or 503. The outcome classes are: 2xx success; timeout, 429, 500, 502, 503, and 504 retryable; all other statuses fatal.

Input is split at line feeds; a final line without a line feed counts, and one trailing carriage return is removed from each line. Up to 20 lines are allowed, including comments and blank lines. Lines containing only spaces or tabs and lines whose first word starts with `#` are skipped. Words are separated only by ASCII spaces and tabs; other whitespace is not a separator. Each request may contain at most 12 outcomes, and IDs must be unique. Empty input is a valid empty log.

Validation checks the total line limit first, then visits lines in order. For each non-skipped line it checks, in order: request shape (at least an ID and an outcome, with the colon attached to the ID); ID validity; each outcome from left to right; the maximum outcome count; and duplicate IDs. The first problem is reported.

## Retry policy

For each request, outcomes represent successive attempts. The replay stops at the first applicable step:

1. A success ends the request as `ok`.
2. A fatal outcome ends it as `failed`, reporting its status.
3. A retryable outcome on the configured last try gives up with reason `tries`.
4. A Retry-After value greater than the cap gives up with reason `retry-after`.
5. Otherwise, calculate the delay as base × factor to the power of the previous attempt, capped at the configured cap; use the larger of that delay and Retry-After when present. If adding it to waits so far would exceed the deadline, give up with reason `deadline`. A wait exactly equal to the remaining deadline is allowed.
6. If there is no next outcome, mark the request `unfinished`; do not count or report a wait that cannot lead to another attempt.
7. Otherwise commit the wait and replay the next outcome.

`plan`, `waits`, and `breaker` ignore outcomes after replay ends. `tally` counts every outcome in the log.

## Circuit breaker

`breaker` replays requests in log order through one persistent circuit breaker and a virtual clock that starts at zero. The breaker starts closed and counts consecutive retryable failures across requests. Before an attempt while open, the clock advances to the later of its current time and the opening time plus cooldown, then the breaker becomes half-open. Each call advances the clock by `--call`.

A retryable outcome is a failure: it increments the closed-state streak or reopens a half-open breaker. Reaching `--open` failures opens the breaker and resets the streak. Success or fatal outcomes reset the streak; either closes a half-open breaker. After each call, the retry policy determines whether a committed retry wait advances the clock. Cooldown and call duration do not count toward the retry deadline. Each state change is printed; the final line reports state, clock time, and number of openings.

## Output

Successful commands write their results to standard output and nothing to standard error. Log errors and command-line errors write no standard output and one diagnostic line to standard error. Exit statuses below are application statuses (the `go run` launcher may additionally print its own diagnostics when the program exits nonzero).

- `check`: prints `ok N requests` (for example, `ok 1 request` or `ok 0 requests`); exits 0.
- `plan`: prints one request line in log order, then `N requests: O ok, F failed, G gave up, U unfinished`. A request line has one of these forms: `ID: ok after N attempts, waited W ms`, `ID: failed at attempt N with STATUS, waited W ms`, `ID: gave up after N attempts (REASON), waited W ms`, or `ID: unfinished after N attempts, waited W ms`. The attempt count uses singular for one. Exits 0 only if every request is successful (including an empty log); otherwise exits 1.
- `waits`: prints each committed wait as `ID #A MS ms`, then `N waits, T ms total` (including `0 waits, 0 ms total` when there are no committed waits). Exits 0.
- `breaker`: prints each transition as `TIME STATE ID`, followed by `final STATE at TIME ms, N opens`. An empty log prints `final closed at 0 ms, 0 opens`. Exits 1 if it opened at least once, otherwise 0.
- `tally`: prints each numeric status as `STATUS COUNT` in ascending order, then `timeout COUNT` if timeout occurred, then `N outcomes in M requests`. Retry-After variants are counted with their numeric status. With one or more outcomes, it also prints `success S, retryable R, fatal F` and `retryable share P%`; the percentage is rounded half-up to one decimal place. An empty log prints `0 outcomes in 0 requests` and omits the class and share lines. Exits 0.

Counts use singular nouns only for one. Percentages are rounded half-up to tenths (for example, `66.7%`).

These are POSIX-shell examples run in the repository directory with the `retrylane` function defined as above. The final example uses `2>&1` to show diagnostics from standard error alongside standard output.

```text
$ printf 'chat: 503 503@2000 200\nembed: 429 200\n' | retrylane plan
chat: ok after 3 attempts, waited 2500 ms
embed: ok after 2 attempts, waited 500 ms
2 requests: 2 ok, 0 failed, 0 gave up, 0 unfinished
$ printf 'chat: 503 503@2000 200\n' | retrylane waits --base 250
chat #1 250 ms
chat #2 2000 ms
2 waits, 2250 ms total
$ printf 'a: 503 503 503 200\n' | retrylane breaker --open 2 --cooldown 500; printf 'a: 200 500\nb: timeout\n' | retrylane tally
700 open a
1700 half-open a
1800 open a
3800 half-open a
3900 closed a
final closed at 3900 ms, 2 opens
200 1
500 1
timeout 1
3 outcomes in 2 requests
success 1, retryable 2, fatal 0
retryable share 66.7%
$ printf 'a: 200\nb: 20\n' | retrylane check 2>&1; printf 'a: 503\n' | retrylane plan --tries 0 2>&1
retrylane: line 2: bad outcome 20
retrylane: bad value for --tries
```

## Errors

Command-line errors exit 2; log validation errors exit 1. The diagnostic line is prefixed with `retrylane: `. Command-line messages are:

- `unknown command WORD`
- `unexpected argument WORD`
- `unknown option WORD`
- `option OPTION is not valid for COMMAND`
- `given twice: OPTION`
- `missing value for OPTION`
- `bad value for OPTION`

Log validation errors are checked in the precedence described above and exit 1:

- `input is too long (max 20 lines)`
- `line N: expected ID: OUTCOME...`
- `line N: bad id ID`
- `line N: bad outcome WORD`
- `line N: too many outcomes`
- `line N: duplicate request ID`

The no-argument case is the exception to the prefixed diagnostic format: it writes the usage synopsis to standard error and exits 2. Help and version are successful standard-output replies.

## Development

Run these checks from the repository directory:

```sh
go vet ./...
gofmt -l .
go test ./...
```

`gofmt -l .` should print nothing.

---

Built with [BAKED](https://github.com/Gulliby-development/baked), an open-source coding agent (bench round R100).
