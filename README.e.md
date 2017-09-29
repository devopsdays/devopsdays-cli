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

# Usage

#### $ {{exec "devopsdays-cli" "--help" | color "sh"}}

{{/* #### $ {{shell "emd gen -help" | color "sh"}}

#### $ {{shell "emd init -help" | color "sh"}} */}}

# History

[CHANGELOG](CHANGELOG.md)
