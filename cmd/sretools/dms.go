package sretools

import "github.com/spf13/cobra"

var DmsCmd = &cobra.Command{
	Use:   "dms",
	Short: "DMS Service Commands",
	Long:  "Commands related to Database Migration Service.",
}

func init() {
	rootCmd.AddCommand(DmsCmd)
}
