// Package cmd defines and implements command-line commands and flags
// used by devopsdays. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// webdir is the path to the source files for the Hugo website
var webdir = setWebdir()

// const webdir = "/Users/mattstratton/src/devopsdays-web"

var cfgFile string
var cityFlag string
var yearFlag string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "devopsdays-cli",
	Short: "Run maintenance tasks for the devopsdays.org website",
	Long: `Command-line utilities for the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/devopsdays/devopsdays-cli`,
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
	RootCmd.PersistentFlags().StringVarP(&cityFlag, "city", "c", "", "city name")
	RootCmd.PersistentFlags().StringVarP(&yearFlag, "year", "y", "", "year")

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devopsdays.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".devopsdays-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")           // adding home directory as first search path
	viper.AutomaticEnv()                   // read in environment variables that match

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
