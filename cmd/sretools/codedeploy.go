package sretools

import "github.com/spf13/cobra"

var CodedeployCmd = &cobra.Command{
	Use:   "codedeploy",
	Short: "CodeDeploy Commands",
	Long:  "Commands related to CodeDeploy.",
}

func init() {
	rootCmd.AddCommand(CodedeployCmd)
}
