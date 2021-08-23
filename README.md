# lint-install

Install sanely configured linters to your project.

This tool specifically supports creating and updating `Makefile` targets, with configured linters for the following languages:

- Go
- Shell
- Dockerfile

## Philosophy

Catch as much as possible, but don't warn about issues that the language authors themselves do not believe in.

## Usage

```
go get github.com/tinkerbell/lint-install
lint-install <repo>
```

## Options

* `--dry-run`: Log what changes would be made if any
* `--shell=warn`: Make shell warnings non-fatal
* `--dockerfile=warn`:  Make Dockerfile warnings non-fatal
* `--go=warn`:  Make Dockerfile warnings non-fatal

## Languages

- Go
- Shell
- Dockerfile
