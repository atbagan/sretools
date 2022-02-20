package sretools

import "github.com/spf13/cobra"

var EventbridgeCmd = &cobra.Command{
	Use:   "eventbridge",
	Short: "EventBridge Service Commands",
	Long:  "Commands related to EventBridge.",
}

func init() {
	rootCmd.AddCommand(EventbridgeCmd)
}
