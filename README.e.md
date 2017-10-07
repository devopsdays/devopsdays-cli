---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---

# {{.Name}}

{{template "badge/travis" .}} [![Build status](https://ci.appveyor.com/api/projects/status/u7pu7ins2csxngxu?svg=true)](https://ci.appveyor.com/project/DevOpsDays/devopsdays-cli) [![Coveralls](https://img.shields.io/coveralls/devopsdays/devopsdays-cli.svg)]()
 {{template "badge/goreport" .}} {{template "badge/godoc" .}} [![GitHub release](https://img.shields.io/github/release/devopsdays/devopsdays-cli.svg)](https://github.com/devopsdays/devopsdays-cli/releases) {{template "license/shields" .}}

Command-line utilities for the [devopsdays](https://www.devopsdays.org) website built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).

![devopsdays gopher](gopher.png)![devopsdays yak](yak.png)

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. We appreciate your contribution. Please refer to the [contributing guidelines](CONTRIBUTING.md) for details on how to help.

# {{toc 5}}

# Demo

[![asciicast](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3.png)](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3)

# Install

{{template "gh/releases" .}}

#### Bintray
{{template "choco_bintray/install" .}}

#### homebrew

```sh
brew install devopsdays/tap/devopsdays-cli
```

#### via Go
{{template "go/install" .}}


# Usage

#### $ {{exec "devopsdays-cli" | color "sh"}}

#### $ {{exec "devopsdays-cli" "create" "speaker" "--help"| color "sh"}}

#### $ {{exec "devopsdays-cli" "create" "event" "--help"| color "sh"}}


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
$ changelog prepare
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