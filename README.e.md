---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---

# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

Command-line utilities for the [devopsdays](https://www.devopsdays.org) website built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).


# {{toc 5}}

# Install

{{template "gh/releases" .}}

#### Bintray
{{template "choco_bintray/install" .}}

#### homebrew

```sh
brew install devopsdays/homebrew-devopsdays-cli/devopsdays-cli
```

#### via Go
{{template "go/install" .}}


# Usage

#### $ {{exec "devopsdays-cli" "--help" | color "sh"}}

{{/* #### $ {{shell "emd gen -help" | color "sh"}}

#### $ {{shell "emd init -help" | color "sh"}} */}}

# History

[CHANGELOG](CHANGELOG.md)

# How to release

```sh
$ changelog prepare
$ gump <patch|minor|major>
```

## Tools needed for release

- [https://github.com/mh-cbon/commit](commit)
- [https://github.com/mh-cbon/changelog](changelog)
- [https://github.com/mh-cbon/emd#toc](emd)
