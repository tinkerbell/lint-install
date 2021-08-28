# lint-install

[![GoReport Widget]][GoReport Status]
![](https://img.shields.io/badge/Stability-Experimental-red.svg)

[GoReport Status]: https://goreportcard.com/report/github.com/tinkerbell/lint-install
[GoReport Widget]: https://goreportcard.com/badge/github.com/tinkerbell/lint-install

Install well-configured linters to your project in a consistent and repeatable way. This tool specifically supports creating and updating `Makefile`
targets, and lints the following:

- Go
- Shell
- Dockerfile

## Philosophy

Catch as many errors as possible, but be idiomatic to the language. 

## Usage

```
go get github.com/tinkerbell/lint-install
$HOME/go/bin/lint-install <repo>
```

## Options

* `--dry-run`: Log what changes would be made if any
* `--shell=warn`: Make shell warnings non-fatal
* `--dockerfile=warn`:  Make Dockerfile warnings non-fatal
* `--go=warn`:  Make Dockerfile warnings non-fatal
