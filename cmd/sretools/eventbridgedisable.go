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
	"github.com/atbagan/sretools/internal/helpers"
	"github.com/spf13/cobra"
)

var eventDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Modify EPS",
	Long:  "Modify Vpc EPS to associate LB with EPS",
	Run:   disableEventRules,
}

func init() {
	eventbridgeCmd.AddCommand(eventDisableCmd)
}

func disableEventRules(cmd *cobra.Command, args []string) {
	awsConfig := config.DefaultAwsConfig(*settings)
	rules := helpers.GetEventbridgeRules(awsConfig.EventbridgeClient())

	for _, v := range rules {
		//params := &eventbridge.DisableRuleInput{
		//	Name:         aws.String(v),
		//
		//}
		fmt.Println(v)
	}

}
