package sretools

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var dmsStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stopping of DMS",
	Long:  "Stops DMS service",
	Run:   stopDms,
}

func init() {
	dmsCmd.AddCommand(dmsStopCmd)
}

func stopDms(cmd *cobra.Command, args []string) {
	awsConfig := c.DefaultAwsConfig(*settings)
	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &databasemigrationservice.StopReplicationTaskInput{
		ReplicationTaskArn: &configuration.Dms.TaskArn,
	}

	_, err = awsConfig.DmsClient().StopReplicationTask(context.TODO(), params)

	if err != nil {
		fmt.Sprintf("failed to stop dms task, %v", err)
		os.Exit(1)
	}
}
