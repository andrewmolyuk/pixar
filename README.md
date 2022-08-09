# go-photos

![GitHub Actions](https://github.com/andrewmolyuk/go-photos/actions/workflows/ci.yml/badge.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a2731a9c8e33458baea3e9ad9c362d8c)](https://www.codacy.com/gh/andrewmolyuk/go-photos/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=andrewmolyuk/go-photos&amp;utm_campaign=Badge_Grade)

## Prerequisites

### Staticcheck

Staticcheck is a state-of-the-art linter for the Go programming language. Beginning with Go 1.17, the simplest way of
installing Staticcheck is by running:

```shell
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### golangci-lint

Golangci-lint is a Go linters aggregator. You can install a binary release on macOS using brew:

```shell
brew install golangci-lint
brew upgrade golangci-lint
```

### gocyclo

Gocyclo calculates cyclomatic complexities of functions in Go source code.

```shell
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

## Commands

### make lint

Run both linters, prune any no-longer-needed dependencies from `go.mod` and perform internal code formatter.

### make build

Generate binary file suitable for the local OS.