package cmd

import "github.com/spf13/cobra"

var dmsRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart of DMS",
	Long:  "Reboots DMS service and does health checks",
	Run:   restartDms,
}

func init() {
	dmsCmd.AddCommand(dmsRestartCmd)
}

func restartDms(cmd *cobra.Command, args []string) {

}
