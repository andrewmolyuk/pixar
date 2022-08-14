# pixar

[![GitHub Actions](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml/badge.svg)](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml)
[![Codacy Grade Badge](https://app.codacy.com/project/badge/Grade/a2731a9c8e33458baea3e9ad9c362d8c)](https://www.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=andrewmolyuk/pixar&amp;utm_campaign=Badge_Grade)
[![Codacy Coverage Badge](https://app.codacy.com/project/badge/Coverage/a2731a9c8e33458baea3e9ad9c362d8c)](https://www.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=github.com&utm_medium=referral&utm_content=andrewmolyuk/pixar&utm_campaign=Badge_Coverage)
![GitHub release](https://img.shields.io/github/v/release/andrewmolyuk/pixar)

Pixar is a pics archiver written for my personal needs. I sometimes use it to archive my photos to an external drive and
to the cloud.

## Features

- Automatically move or copy photos and videos to structured folders formatted by Year, Month and Date based on the
  file's EXIF information when the file was created

- Run from a batch file or the command line on Mac, Linux or Windows
- Choose which file extensions to process
- Simulation run for testing and checking purposes

## Roadmap and User Suggestions

- Log the list of the processed files in the CSV format
- Move or copy unhandled files into separate folder
- Optionally move or copy files without exif data to structured folders using the file date
- Control how duplicate files are handled when found. Skip, Rename, Overwrite or Move to a separate folder
- Create sub folder for events when more than specific amount of pictures where created during specific time interval
- Add files to delete during processing, like Thumbs.db or .DS_Store files
- Add ability to sync/upload photos with cloud drives: S3, iDrive etc
- Add custom format for output folder names

Any useful idea or suggestion is welcomed.

## Installation

Download the latest binary archive suitable to your operating system
from [GitHub releases](https://github.com/andrewmolyuk/pixar/releases/latest). Extract the archive to any preferable
place, and you are ready to run the binary.

For more comfortable usage you can move binary to `/usr/local/bin` on your MacOS or `/usr/bin` on your Linux system.
For Windows, you can move it to `C:\Program Files\Pixar` and add the folder to the `%PATH%` environment variable.

## Usage

in order to get help on how to use the application, run:

```shell
pixar --help
```

You have to get the similar to following output in your console:

```shell
pixar [OPTIONS]

Scan folders and move photos and videos into folders according to their EXIF information

Application Options:
  -i, --input=      Input folder (default: .)
  -o, --output=     Output folder (default: output)
  -m, --move        Move files instead of copying them
  -d, --debug       Debug mode
  -v, --version     Show Pixar version info
  -e, --extensions= File extensions to process (default: .jpeg,.jpg,.tiff,.png)

Help Options:
  -h, --help     Show this help message
```

## Examples

```shell
pixar -v

pixar --input ./photos

rm -Rf ./testdata/output && go pixar --debug --input ./testdata/input --output ./testdata/output
```

## Development

### Prerequisites

#### Staticcheck

Staticcheck is a state-of-the-art linter for the Go programming language. Beginning with Go 1.17, the simplest way of
installing Staticcheck is by running:

```shell
go install honnef.co/go/tools/cmd/staticcheck@latest
```

#### golangci-lint

Golangci-lint is a Go linters aggregator. You can install a binary release on macOS using brew:

```shell
brew install golangci-lint
brew upgrade golangci-lint
```

#### gocyclo

Gocyclo calculates cyclomatic complexities of functions in Go source code.

```shell
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

### Commands

#### make lint

Run both linters, prune any no-longer-needed dependencies from `go.mod` and perform internal code formatter.

#### make test

Run all tests in the project and print the results to the console.

#### make build

Generate binary file suitable for the local OS.

#### make run

Execute the application with files from the `testdata` folder. 
