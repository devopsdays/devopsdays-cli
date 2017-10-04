
# devopsdays-cli

[![travis Status](https://travis-ci.org/devopsdays/devopsdays-cli.svg?branch=master)](https://travis-ci.org/devopsdays/devopsdays-cli) [![Build status](https://ci.appveyor.com/api/projects/status/u7pu7ins2csxngxu?svg=true)](https://ci.appveyor.com/project/DevOpsDays/devopsdays-cli) [![Coveralls](https://img.shields.io/coveralls/devopsdays/devopsdays-cli.svg)]()
 [![Go Report Card](https://goreportcard.com/badge/github.com/devopsdays/devopsdays-cli)](https://goreportcard.com/report/github.com/devopsdays/devopsdays-cli) [![GoDoc](https://godoc.org/github.com/devopsdays/devopsdays-cli?status.svg)](http://godoc.org/github.com/devopsdays/devopsdays-cli) [![GitHub release](https://img.shields.io/github/release/devopsdays/devopsdays-cli.svg)](https://github.com/devopsdays/devopsdays-cli/releases) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Command-line utilities for the [devopsdays](https://www.devopsdays.org) website built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).


# TOC
- [Install](#install)
  - [Bintray](#bintray)
  - [homebrew](#homebrew)
  - [via Go](#via-go)
- [Usage](#usage)
  - [$ devopsdays-cli](#-devopsdays-cli)
- [History](#history)
- [How to release](#how-to-release)
  - [Tools needed for release](#tools-needed-for-release)

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

#### $ devopsdays-cli
```sh
_                               _                            _ _
   __| | _____   _____  _ __  ___  __| | __ _ _   _ ___        ___| (_)
  / _` |/ _ \ \ / / _ \| '_ \/ __|/ _` |/ _` | | | / __|_____ / __| | |
 | (_| |  __/\ V / (_) | |_) \__ \ (_| | (_| | |_| \__ \_____| (__| | |
  \__,_|\___| \_/ \___/| .__/|___/\__,_|\__,_|\__, |___/      \___|_|_|
                       |_|                    |___/

Command-line utilities for the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/devopsdays/devopsdays-cli

Usage:
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
_                               _                            _ _
   __| | _____   _____  _ __  ___  __| | __ _ _   _ ___        ___| (_)
  / _` |/ _ \ \ / / _ \| '_ \/ __|/ _` |/ _` | | | / __|_____ / __| | |
 | (_| |  __/\ V / (_) | |_) \__ \ (_| | (_| | |_| \__ \_____| (__| | |
  \__,_|\___| \_/ \___/| .__/|___/\__,_|\__,_|\__, |___/      \___|_|_|
                       |_|                    |___/
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

# History

[CHANGELOG](CHANGELOG.md)

# How to release

```sh
$ changelog prepare
$ gump <patch|minor|major>
```

## Tools needed for release

- [commit](https://github.com/mh-cbon/commit)
- [changelog](https://github.com/mh-cbon/changelog)
- [emd](https://github.com/mh-cbon/emd)
