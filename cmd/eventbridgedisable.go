package cmd

import (
	"fmt"
	"github.com/atbagan/sretools/config"
	"github.com/atbagan/sretools/helpers"
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
