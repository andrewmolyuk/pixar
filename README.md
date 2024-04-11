# pixar

[![GitHub Actions](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml/badge.svg)](https://github.com/andrewmolyuk/pixar/actions/workflows/ci.yml)
[![Go Report Badge](https://goreportcard.com/badge/github.com/andrewmolyuk/pixar)](https://goreportcard.com/report/github.com/andrewmolyuk/pixar)
[![Codacy Grade Badge](https://app.codacy.com/project/badge/Grade/a2731a9c8e33458baea3e9ad9c362d8c)](https://app.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/a2731a9c8e33458baea3e9ad9c362d8c)](https://app.codacy.com/gh/andrewmolyuk/pixar/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)
[![GitHub release](https://img.shields.io/github/v/release/andrewmolyuk/pixar)](https://github.com/andrewmolyuk/pixar/releases)
[![License Badge](https://img.shields.io/github/license/andrewmolyuk/pixar)](https://opensource.org/licenses/MIT)

## Description

Pixar is a command line tool to organize photos and videos according to their EXIF information. It's written in Go and
can be run on Mac, Linux or Windows.

The main idea is to copy or move photos and videos to structured folders according embedded EXIF information. The
application is lossless, so it doesn't modify the original files.

## Features

The main features of Pixar are:

- Nothing is overridable during the run of the program, so it's lossless
- Copy or move photos and videos to structured folders according embedded EXIF information
- Run from a batch file or the command line on Mac, Linux or Windows
- Choose which file extensions to process
- Simulation run for testing and checking purposes
- Log the list of the processed files in the CSV file
- Limit the number of actions to be performed at the same time
- Skip duplicate files or place them into a separate folder

## Roadmap and User Suggestions

- Move or copy unhandled files into separate folder
- Optionally move or copy files without exif data to structured folders using the file modification date
- Create sub folder for events when more than specific amount of pictures where created during specific time interval
- Define files to delete during processing, like Thumbs.db or .DS_Store files
- Sync/upload photos with cloud drives: S3, iDrive etc
- Custom format for output folder names
- Apply CSV file with list of actions have to be performed
- Load configuration from a JSON file

Any useful idea or suggestion is welcomed.

## Installation

The easiest way to install Pixar is to download the binary from
the [releases page](https://github.com/andrewmolyuk/pixar/releases/latest). Extract the binary from the archive and run
it from the command line.

For more comfortable usage you can move binary to `/usr/local/bin` on your MacOS or `/usr/bin` on your Linux system.
For Windows, you can move it to `C:\Program Files\Pixar` and add the folder to the `%PATH%` environment variable.

## Usage

Get the list of available options with `pixar --help` command.

```shell
pixar --help
```

The following output in the console is expected:

```markdown
pixar [OPTIONS]

Scan folders and move photos and videos into folders according to their EXIF information

Application Options:
-i, --input= Input folder (default: .)
-o, --output= Output folder (default: output)
-m, --move Move files instead of copying them
-d, --debug Debug mode
-v, --version Show Pixar version info
-e, --extensions= File extensions to process (default: .jpeg,.jpg,.tiff,.png)
-s, --simulation Simulation mode
-c, --csv= CSV file name for actions output
-n, --concurrent= Maximum number of concurrent operations (default: 100)
-p, --policy= Policy for duplicates: skip, folder (default: skip)

Help Options:
-h, --help Show this help message
```

### Options

#### -i, --input

Input folder where the application will start to scan for files to process. Default value is the current folder. Can be
an absolute or relative path.

#### -o, --output

Output folder where the application will put the processed files. Default value is `output` folder in the current
folder. Can be an absolute or relative path.

#### -m, --move

If this option is set, the application will move the files instead of copying them.

#### -d, --debug

If this option is set, the application will run in debug mode and provide more detailed output.

#### -v, --version

If this option is set, the application will show the version info.

#### -e, --extensions

File extensions to process. Default value is `.jpeg,.jpg,.tiff,.png`.

#### -s, --simulation

If this option is set, the application will run in simulation mode and don't perform any actions.

#### -c, --csv

CSV file name for actions output. If this option is set, the application will create a CSV file with the actions for
detailed check and review.

#### -n, --concurrent

Limit of actions amount to perform at the same time. Default value is `100`.

#### -p, --policy

Define policy for duplicates: skip, folder. Default value is `skip`. When folder policy is set, the application will
copy or move duplicates into separate folder named `Duplicates` under `output` folder with all `input` folders structure
where the duplicate was found. In order if file already exists in duplication folder the application is failed.

## Examples

```shell
pixar -v

pixar --input ./photos -s -p folder

rm -Rf ./testdata/output && go pixar --debug --input ./testdata/input --output ./testdata/output
```

## Development

### Prerequisites

#### Staticcheck

Staticcheck is a state-of-the-art linter for the Go programming language. Beginning with Go 1.17, the simplest way of
installing Staticcheck is by running:

```shell
brew install staticcheck
```

or

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

#### markdownlint

We are using markdownlint to lint all markdown files in the project.

```shell
brew install markdownlint-cli
```

### Commands

#### make lint

Run both linters, prune any no-longer-needed dependencies from `go.mod` and perform internal code formatter.

#### make test

Run all tests in the project and print the results to the console.

#### make build

Generate binary file suitable for the local OS.

#### make dev

Execute the application with files from the `testdata` folder.

#### make sim

Execute the application in simulation mode with files from the `testdata` folder.
