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

This option defines the policy for handling duplicates: `skip` or `folder`. The default value is `skip`. If the `folder`
policy is selected, the application will copy or move duplicate files into a separate `Duplicates` folder located within
the `output` folder, maintaining the original `input` folder structure. If a duplicate file already exists in the
`Duplicates` folder, the application will terminate with an error.

## Examples

```shell
pixar -v

pixar --input ./photos -s -p folder

rm -Rf ./testdata/output && go pixar --debug --input ./testdata/input --output ./testdata/output
```

## Development

To contribute to the development of this project, follow the steps below:

1. Fork the repository and clone it to your local machine
2. Navigate to your local repo and create a new branch
3. Make your changes or additions to the code
4. Push your changes to your fork
5. Submit a pull request detailing the changes you made

Before submitting your pull request, please ensure your code adheres to the following guidelines:

- Code must be properly formatted and linted.
- Include comments for complex sections of code.
- Write tests for new features or bug fixes.
- Update the README.md file if necessary.

Your pull request will be reviewed by the maintainers and merged if it meets the standards of the project.

### Prerequisites

#### Staticcheck

Staticcheck is a cutting-edge linter for the Go programming language. Starting from Go 1.17, the easiest method to
install Staticcheck is by executing the following commands:

```shell
brew install staticcheck
```

or

```shell
go install honnef.co/go/tools/cmd/staticcheck@latest
```

#### golangci-lint

Golangci-lint is an aggregator of linters for Go programming language. It simplifies the process of using multiple
linters and provides a unified output. For macOS users, a binary release can be installed using the brew package
manager:

```shell
brew install golangci-lint
brew upgrade golangci-lint
```

#### gocyclo

Gocyclo is a tool that calculates the cyclomatic complexities of functions within Go source code. This complexity is a
quantitative measure of the number of linearly independent paths through a program's source code. It is a useful metric
for understanding the complexity and maintainability of a particular function.

```shell
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

#### markdownlint

We are using markdownlint to lint all markdown files in the project. You can install it using brew:

```shell
brew install markdownlint-cli
```

### Commands

#### make lint

This command run all the linters, including Staticcheck and Golangci-lint, to catch any potential issues in the code
that could lead to bugs or make the code harder to read and maintain.

In addition to linting, this command will also prune any dependencies from `go.mod` that are no longer needed. This is
an important step in managing the project's dependencies, as it helps to keep the project clean, reduces the size of the
project, and can even improve the build time and runtime efficiency.

Lastly, this command will run the internal code formatter. Code formatting is a critical aspect of any project as it
ensures that all the code in the project follows a consistent style.

In addition, this command is used to execute the application in simulation mode using files from the testdata folder.
The simulation mode allows you to run the application without performing any actual actions, which is particularly
useful for testing and debugging. It enables you to observe the application's behavior and identify any potential issues
or anomalies in its operation without affecting any real data. It's recommended to use this command during the
development process to ensure that the application behaves as expected in different scenarios.

In summary, the `make lint` command is a comprehensive tool for maintaining the quality, efficiency, and readability of
the project's code.

#### make test

Run all tests in the project and print the results to the console. This command is crucial for ensuring the stability
and reliability of the codebase. It executes all unit tests and integration tests, providing a comprehensive overview of
the project's health. If any test fails, the command will return an error, making it easy to spot and fix issues early
in the development process.

#### make build

This command is used to compile the source code of the project into an executable binary file. The `make build` command
will also ensure that all dependencies are correctly linked and that the resulting binary is optimized for the local
operating system. This command is particularly useful when you're ready to test your application locally or prepare it
for distribution. It's recommended to run this command after making any changes to the source code to ensure that the
latest version of the application is always available for testing or deployment.

#### make dev

This command is used to execute the application using files from the `testdata` folder. This can help to identify any
issues or bugs in the application before it's used with real data. It's recommended to use this command during the
development process to ensure that the application behaves as expected.

## License

This project is licensed under the terms of the MIT license. The MIT license is a permissive free software license that
allows for reuse within proprietary software provided that all copies of the licensed software include a copy of the MIT
License terms and the copyright notice. For more details, see the `LICENSE` file in the project root.
