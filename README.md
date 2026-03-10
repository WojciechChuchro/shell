[![progress-banner](https://backend.codecrafters.io/progress/shell/30a941f8-a07b-46be-b00e-68a940a404fb)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# A small shell in Go

This is my personal implementation of a small POSIX‑style shell written in Go.
The project started as my solution to the
["Build Your Own Shell" challenge](https://app.codecrafters.io/courses/shell/overview)
on [Codecrafters](https://codecrafters.io), and I'm iterating on it here as a
stand‑alone project.

Right now the shell is in an early stage: it reads a command from standard
input and reports it as "command not found". Over time, the goal is to support:

- Basic REPL-style interaction
- Running external programs
- Built-in commands like `cd`, `pwd`, `echo`, etc.
- More POSIX-style shell features as the project evolves

## Getting started

### Prerequisites

- Go `1.26` or newer installed locally

### Running locally

The main entry point for the shell is in `app/main.go`.

To run it directly with Go:

```sh
go run ./app
```

If you're working through the Codecrafters track, you can also use the helper
script (this is what their test runner uses under the hood):

```sh
./your_program.sh
```

### Example interactions

The shell aims to behave like a typical POSIX shell. For example, the `type`
builtin should be able to distinguish between external commands and builtins:

```sh
$ type grep
grep is /usr/bin/grep
$ type invalid_command
invalid_command: not found
$ type echo
echo is a shell builtin
```

As more features are implemented, this section will grow with additional
examples.

## Running the Codecrafters tests

If you're viewing this as part of the Codecrafters challenge and have the
`codecrafters` CLI installed, you can run their test suite against this
repository with:

```sh
codecrafters submit
```

Behind the scenes, the tester will invoke your program with a custom `PATH`
that looks roughly like this:

```sh
PATH="/usr/bin:/usr/local/bin:$PATH" ./your_program.sh
```

This ensures standard binaries like `grep` are available while still giving your
shell control over how commands are resolved.

The tests themselves run on Codecrafters' infrastructure, and output is streamed
back to your terminal.

## Acknowledgements

This project was built in collaboration with
[Codecrafters](https://codecrafters.io), whose "Build Your Own Shell" course
provides the original spec, test harness and progression of stages. The code
and design decisions here are my own, but heavily inspired by that course.
