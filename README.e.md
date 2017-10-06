---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---

# {{.Name}}

{{template "badge/travis" .}} [![Build status](https://ci.appveyor.com/api/projects/status/u7pu7ins2csxngxu?svg=true)](https://ci.appveyor.com/project/DevOpsDays/devopsdays-cli) [![Coveralls](https://img.shields.io/coveralls/devopsdays/devopsdays-cli.svg)]()
 {{template "badge/goreport" .}} {{template "badge/godoc" .}} [![GitHub release](https://img.shields.io/github/release/devopsdays/devopsdays-cli.svg)](https://github.com/devopsdays/devopsdays-cli/releases) {{template "license/shields" .}}

Command-line utilities for the [devopsdays](https://www.devopsdays.org) website built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).

![devopsdays gopher](https://raw.githubusercontent.com/devopsdays/devopsdays-cli/master/gopher.png)

[![asciicast](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3.png)](https://asciinema.org/a/htDFVyRFgai6p8yq2QAplLTd3)

# {{toc 5}}

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
