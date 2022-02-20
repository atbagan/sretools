package sretools

import "github.com/spf13/cobra"

var eventbridgeCmd = &cobra.Command{
	Use:   "eventbridge",
	Short: "EventBridge Service Commands",
	Long:  "Commands related to EventBridge.",
}

func init() {
	rootCmd.AddCommand(eventbridgeCmd)
}
