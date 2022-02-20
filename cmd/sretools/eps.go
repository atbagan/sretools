package sretools

import (
	"github.com/spf13/cobra"
)

var EpsCmd = &cobra.Command{
	Use:   "eps",
	Short: "Endpoint Service Commands",
	Long:  "Commands related to vpc endpoint services.",
}

func init() {
	rootCmd.AddCommand(EpsCmd)
}
