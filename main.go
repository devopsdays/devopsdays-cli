// Command-line utilities for the [devopsdays](https://www.devopsdays.org) website
package main

import (
	"github.com/devopsdays/devopsdays-cli/cmd"
	_ "github.com/dimiro1/banner/autoload"
)

func main() {
	cmd.Execute()
}
