package sretools

import "github.com/spf13/cobra"

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "ECS Service Commands",
	Long:  "Commands related to Elastic Container Service.",
}

func init() {
	rootCmd.AddCommand(ecsCmd)
}
