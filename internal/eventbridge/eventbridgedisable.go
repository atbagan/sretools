package eventbridge

import (
	"fmt"
	tools "github.com/atbagan/sretools/cmd/sretools"
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
	tools.EventbridgeCmd.AddCommand(eventDisableCmd)
}

func disableEventRules(cmd *cobra.Command, args []string) {
	awsConfig := config.DefaultAwsConfig(*tools.Settings)
	rules := helpers.GetEventbridgeRules(awsConfig.EventbridgeClient())

	for _, v := range rules {
		//params := &eventbridge.DisableRuleInput{
		//	Name:         aws.String(v),
		//
		//}
		fmt.Println(v)
	}

}
