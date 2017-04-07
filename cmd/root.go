// The MIT License (MIT)
// Copyright (c) 2017 Matt Stratton
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package cmd defines and implements command-line commands and flags
// used by probablyfine. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// webdir is the path to the source files for the Hugo website
var webdir string = setWebdir()

// const webdir = "/Users/mattstratton/src/devopsdays-web"

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "probablyfine",
	Short: "Run maintenance tasks for the devopsdays.org website",
	Long: `Command-line utilities for the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/mattstratton/probablyfine`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.probablyfine.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".probablyfine") // name of config file (without extension)
	viper.AddConfigPath("$HOME")         // adding home directory as first search path
	viper.AutomaticEnv()                 // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func setWebdir() string {
	if os.Getenv("DODPATH") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return pwd
	} else {
		s := os.Getenv("DODPATH")
		s = strings.TrimSuffix(s, "/")
		s = strings.TrimSuffix(s, "\\")
		return s
	}
}
