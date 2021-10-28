# lint-install

[![GoReport Widget]][GoReport Status]
[![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)](https://github.com/emersion/stability-badges#stable)

[GoReport Status]: https://goreportcard.com/report/github.com/tinkerbell/lint-install
[GoReport Widget]: https://goreportcard.com/badge/github.com/tinkerbell/lint-install

Idiomatic linters for opinionated projects.

This tool installs well-configured linters to any project, open-source or
otherwise. The linters can be used in a repeatable and consistent way across CI,
local tests, and IDE's.

lint-install adds linter configuration to the root of your project, and Makefile
rules to install a consistently versioned set of linters to be used in any
environment. These Makefile rules can also be upgrading by lint-install, updating
all environments simultaneously.

Currently supported languages:

- Go
- Shell
- Dockerfile
- YAML

## Philosophy

- Catch all the bugs!
- Improve readability as much as possible.
- Be idiomatic: only raise issues that the language authors would flag

## Usage

Installation:

`go get github.com/tinkerbell/lint-install`

Add Makefile rules for a git repository:

`$HOME/go/bin/lint-install <repo>`

Users can then lint the project using:

`make lint`

Other options:

```
  -dockerfile string
     Level to lint Dockerfile with: [ignore, warn, error] (default "error")
  -dry-run
     Display changes to make
  -go string
     Level to lint Go with: [ignore, warn, error] (default "error")
  -makefile string
     name of Makefile to update (default "Makefile")
  -shell string
     Level to lint Shell with: [ignore, warn, error] (default "error")
  -yaml string
     Level to lint YAML with: [ignore, warn, error] (default "error")
```
