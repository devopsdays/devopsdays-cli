
# devopsdays-cli

[![travis Status](https://travis-ci.org/devopsdays/devopsdays-cli.svg?branch=master)](https://travis-ci.org/devopsdays/devopsdays-cli) [![Build status](https://ci.appveyor.com/api/projects/status/u7pu7ins2csxngxu?svg=true)](https://ci.appveyor.com/project/DevOpsDays/devopsdays-cli)
 [![Go Report Card](https://goreportcard.com/badge/github.com/devopsdays/devopsdays-cli)](https://goreportcard.com/report/github.com/devopsdays/devopsdays-cli) [![GoDoc](https://godoc.org/github.com/devopsdays/devopsdays-cli?status.svg)](http://godoc.org/github.com/devopsdays/devopsdays-cli) [![GitHub release](https://img.shields.io/github/release/devopsdays/devopsdays-cli.svg)](https://github.com/devopsdays/devopsdays-cli/releases) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) [![ZenHub](https://raw.githubusercontent.com/ZenHubIO/support/master/zenhub-badge.png)](https://app.zenhub.com/workspace/o/devopsdays/devopsdays-cli/)<br />
 [![Coveralls](https://img.shields.io/coveralls/devopsdays/devopsdays-cli.svg)]() [![BCH compliance](https://bettercodehub.com/edge/badge/devopsdays/devopsdays-cli?branch=master)](https://bettercodehub.com/results/devopsdays/devopsdays-cli) [![Ebert](https://ebertapp.io/github/devopsdays/devopsdays-cli.svg)](https://ebertapp.io/github/devopsdays/devopsdays-cli) [![codebeat badge](https://codebeat.co/badges/9acd2699-7397-45fb-8219-e0e5a68c8399)](https://codebeat.co/projects/github-com-devopsdays-devopsdays-cli-master)

Command-line utilities for the [devopsdays](https://www.devopsdays.org) website built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).

![devopsdays gopher](gopher.png)![devopsdays yak](yak.png)

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. We appreciate your contribution. Please refer to the [contributing guidelines](CONTRIBUTING.md) for details on how to help.

# TOC
- [Demo](#demo)
- [Install](#install)
  - [Bintray](#bintray)
  - [homebrew](#homebrew)
  - [via Go](#via-go)
- [Usage](#usage)
  - [$ devopsdays-cli --help](#-devopsdays-cli---help)
  - [$ devopsdays-cli create speaker --help](#-devopsdays-cli-create-speaker---help)
  - [$ devopsdays-cli create event --help](#-devopsdays-cli-create-event---help)
  - [$ devopsdays-cli create sponsor --help](#-devopsdays-cli-create-sponsor---help)
  - [$ devopsdays-cli show speaker --help](#-devopsdays-cli-show-speaker---help)
  - [$ devopsdays-cli show talk --help](#-devopsdays-cli-show-talk---help)
- [Reference](#reference)
  - [Matrix of commands](#matrix-of-commands)
  - [Contributing](#contributing)
- [History](#history)
- [How to release](#how-to-release)
  - [Tools needed for release](#tools-needed-for-release)
- [License](#license)

# Demo

[![asciicast](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3.png)](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3)

# Install

Check the [release page](https://github.com/devopsdays/devopsdays-cli/releases)!

#### Bintray
```sh
choco source add -n=devopsdays -s="https://api.bintray.com/nuget/devopsdays/choco"
choco install devopsdays-cli
```

#### homebrew

```sh
brew install devopsdays/tap/devopsdays-cli
```

#### via Go
```sh
go get github.com/devopsdays/devopsdays-cli
```


# Usage

#### $ devopsdays-cli --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


Command-line utilities for the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/devopsdays/devopsdays-cli

Usage:
  devopsdays-cli [flags]
  devopsdays-cli [command]

Available Commands:
  add         Add items to talks, programs, or events
  create      Create a new event, organizer, program, speaker, sponsor, or talk
  edit        Edit an existing item
  help        Help about any command
  remove      Remove items from an event, a talk, or a program
  show        Show details about various items

Flags:
  -d, --debug   enable debug mode
  -h, --help    help for devopsdays-cli

Use "devopsdays-cli [command] --help" for more information about a command.
```

#### $ devopsdays-cli create speaker --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Creates a new speaker for an event.

Usage:
  devopsdays-cli create speaker [flags]

Examples:
  devopsdays-cli create speaker
  devopsdays-cli create speaker --city new-york
  devopsdays-cli create speaker -c "New York" --year "2017"

Flags:
  -c, --city string   city to use
  -h, --help          help for speaker
  -y, --year string   year to use

Global Flags:
  -d, --debug   enable debug mode
```

#### $ devopsdays-cli create event --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Create a new event.

Usage:
  devopsdays-cli create event create [flags]

Examples:
  devopsdays-cli create event
  devopsdays-cli create event -c New York --year 2017

Flags:
  -c, --city string   city to use
  -h, --help          help for event
  -y, --year string   year to use

Global Flags:
  -d, --debug   enable debug mode
```

#### $ devopsdays-cli create sponsor --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Create a new sponsor file add the sponsor's image.
The name argument must not contain any spaces.

Usage:
  devopsdays-cli create sponsor [name] [flags]

Examples:
  devopsdays-cli create sponsor
  devopsdays-cli create sponsor bluth-company

Flags:
  -h, --help   help for sponsor

Global Flags:
  -d, --debug   enable debug mode
```

#### $ devopsdays-cli show speaker --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Show a speaker from an event.

Usage:
  devopsdays-cli show speaker [flags]

Examples:
  devopsdays-cli show speaker george-bluth
  devopsdays-cli show speaker --city new-york --year 2017 --all
  devopsdays-cli show speaker george-bluth -c "New York" --year "2017"

Flags:
  -a, --all           show all
  -c, --city string   city to use
  -h, --help          help for speaker
  -y, --year string   year to use

Global Flags:
  -d, --debug   enable debug mode
```

#### $ devopsdays-cli show talk --help
```sh
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|d|e|v|o|p|s|d|a|y|s|-|c|l|i|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Show a speaker from an event.

Usage:
  devopsdays-cli show talk [flags]

Examples:
  devopsdays-cli show talk
  devopsdays-cli show talk --city new-york --year 2017 --all
  devopsdays-cli show talk -c "New York" --year "2017"

Flags:
  -a, --all           show all NOT IMPLEMENTED
  -c, --city string   city to use
  -h, --help          help for talk
  -y, --year string   year to use

Global Flags:
  -d, --debug   enable debug mode
```


# Reference

## Matrix of commands


|        | config | event | organizer | program | speaker | sponsor | talk | version |
|--------|--------|-------|-----------|---------|---------|---------|------|---------|
| add    |        |       |           |         | x       | x       | x    |         |
| create |        | x     | x         | x       | x       | x       | x    |         |
| edit   |        | x     | x         | x       | x       | x       | x    |         |
| remove |        |       | x         |         | x       | x       | x    |         |
| show   | x      | x     | x         | x       | x       | x       | x    | x       |


## Contributing

See [CONTRIBUTING.md](contributing.md) for details on how to contribute to this project.

# History

[CHANGELOG](CHANGELOG.md)

# How to release

```sh
$ gump <patch|minor|major>
```

## Tools needed for release

- [commit](https://github.com/mh-cbon/commit)
- [changelog](https://github.com/mh-cbon/changelog)
- [emd](https://github.com/mh-cbon/emd)

# License

devopsdays-cli - Command-line utilities for the devopsdays.org website

|                      |                                          |
|:---------------------|:-----------------------------------------|
| **Author:**          | Matt Stratton (<matt.stratton@gmail.com>)
| **Copyright:**       | Copyright 2017, Matt Stratton
| **License:**         | The MIT License

```
The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

```
