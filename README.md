# pixar

[![GitHub Actions](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml/badge.svg)](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml)
[![Codacy Grade Badge](https://app.codacy.com/project/badge/Grade/a2731a9c8e33458baea3e9ad9c362d8c)](https://www.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=andrewmolyuk/pixar&amp;utm_campaign=Badge_Grade)
[![Codacy Coverage Badge](https://app.codacy.com/project/badge/Coverage/a2731a9c8e33458baea3e9ad9c362d8c)](https://www.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=github.com&utm_medium=referral&utm_content=andrewmolyuk/pixar&utm_campaign=Badge_Coverage)
![GitHub release](https://img.shields.io/github/v/release/andrewmolyuk/pixar)

Pixar is just a pics archiver written for my personal needs. I sometimes use it to archive my photos to an external
drive and to the cloud.

## Installation

You can download the latest binary archive suitable to your operating system
from [GitHub releases](https://github.com/andrewmolyuk/pixar/releases/latest). Extract the archive to any preferable
place, and you are ready to run the binary.

For more comfortable usage you can move binary to `/usr/local/bin` on your MacOS or `/usr/bin` on your Linux system.
For Windows, you can move it to `C:\Program Files\Pixar` and add the folder to the `%PATH%` environment variable.

## Usage

## Commands

### make lint

Run both linters, prune any no-longer-needed dependencies from `go.mod` and perform internal code formatter.

### make test

Run all tests in the project and print the results to the console.

### make build

Generate binary file suitable for the local OS.

## Development prerequisites

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
