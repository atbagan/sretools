// Package sretools
// Copyright (c) 2022, Andrew Bagan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sretools

import (
	"fmt"
	"github.com/atbagan/sretools/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	settings = new(config.Config)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sretools",
	Short: "Various tools for sre's dealing with AWS",
}

// Execute executes the commands. Called in main.Main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Execute func error: ", err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&debug, "verbose", "v", false, "verbose logging")
	settings.Profile = rootCmd.PersistentFlags().StringP("profile", "p", "", "Use a specific profile")
	settings.Region = rootCmd.PersistentFlags().StringP("region", "r", "", "Use a specific region")
	settings.Iam = rootCmd.PersistentFlags().String("iam", "", "Use a specific iam role to assume")
	settings.Arn = rootCmd.PersistentFlags().String("arn", "", "Use a specific arn")
	settings.NameFile = rootCmd.PersistentFlags().StringP("namefile", "n", "", "Use this file to provide names")
	//settings.Ecs.Cluster = rootCmd.PersistentFlags().String("CLUSTER","", "cluster name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName(".sretools") // name of config file (without extension)
	viper.SetConfigType("yml")       // set the config file type
	viper.AddConfigPath(".")         // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
