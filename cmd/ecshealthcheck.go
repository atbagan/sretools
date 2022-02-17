package cmd

import "github.com/spf13/cobra"

var healthCmd = &cobra.Command{
	Use:   "health-check",
	Short: "Health Check for ECS Service",
	Long:  "Checks health for a given target group to determine if your service is healthy or not",
	Run:   getHealthCheck,
}

func init() {
	ecsCmd.AddCommand(healthCmd)
}

func getHealthCheck(cmd *cobra.Command, args []string) {

}
