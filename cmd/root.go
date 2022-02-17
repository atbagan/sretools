package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/atbagan/sretools/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var settings = new(config.Config)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sretools",
	Short: "Various tools for sre's dealing with AWS",
}

//Execute executes the commands. Called in main.Main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	settings.Verbose = rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	settings.Profile = rootCmd.PersistentFlags().StringP("profile", "p", "", "Use a specific profile")
	settings.Region = rootCmd.PersistentFlags().StringP("region", "r", "", "Use a specific region")
	settings.NameFile = rootCmd.PersistentFlags().StringP("namefile", "n", "", "Use this file to provide names")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName(".sretools") // name of config file (without extension)
	viper.SetConfigType("yml")       // set the config file type
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getName(id string) string {
	if *settings.NameFile != "" {
		nameFile, err := ioutil.ReadFile(*settings.NameFile)
		if err != nil {
			panic(err)
		}
		values := make(map[string]string)
		err = json.Unmarshal(nameFile, &values)
		if err != nil {
			panic(err)
		}
		if val, ok := values[id]; ok {
			return val
		}
	}
	return id
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
