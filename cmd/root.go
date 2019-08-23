/*
Copyright © 2019 We Just Do Stuff <hello@wejustdostuff.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/wejustdostuff/carnot/crawler"
)

var conf crawler.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "carnot",
	Short: "File sorting tool to organize files into folders by year, month and day.",
	Long: `This software will move all files from the input directory into the output
    directory without changing the files content. It will only rename the
    files and place them in the proper directory for year, month and day.`,
	Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&conf.File, "config", "c", "", "config file (default is $HOME/.carnot.yaml)")

	rootCmd.Flags().BoolVar(&conf.Dry, "dry", false, "Run in dry mode")
	rootCmd.Flags().BoolVarP(&conf.Verbose, "verbose", "v", false, "Set debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if conf.File != "" {
		// Use config file from the flag.
		viper.SetConfigFile(conf.File)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".carnot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".carnot")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		exit("Source directory is missing")
	}
	if len(args) < 2 {
		exit("Target directory is missing")
	}

	conf.Source = args[0]
	conf.Target = args[1]

	if err := crawler.Run(&conf); err != nil {
		exit(err.Error())
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
