
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
  - [$ devopsdays-cli --help](#-devopsdays-cli---help)
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

#### $ devopsdays-cli --help
```sh
Command-line utilities for the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/devopsdays/devopsdays-cli

Usage:
  devopsdays-cli [command]

Available Commands:
  add         Add a thing to another thing
  create      Make a new thing
  edit        Edit an existing thing
  help        Help about any command
  remove      Remove a thing to another thing
  show        Show details about a thing

Flags:
  -c, --city string   city name
  -h, --help          help for devopsdays-cli
  -y, --year string   year

Use "devopsdays-cli [command] --help" for more information about a command.
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
