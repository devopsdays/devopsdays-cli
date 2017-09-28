---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---

# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

# {{toc 5}}

# Install

{{template "gh/releases" .}}

#### Bintray
{{template "choco_bintray/install" .}}

# Usage

#### $ {{exec "emd" "-help" | color "sh"}}

#### $ {{shell "emd gen -help" | color "sh"}}

#### $ {{shell "emd init -help" | color "sh"}}

# History

[CHANGELOG](CHANGELOG.md)
