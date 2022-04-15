package sretools

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var dmsStartCmd = &cobra.Command{
	Use:   "start",
	Short: "starting of DMS",
	Long:  "Starts DMS replication task service",
	Run:   startDms,
}

func init() {
	dmsCmd.AddCommand(dmsStartCmd)
}

func startDms(cmd *cobra.Command, args []string) {
	awsConfig := c.DefaultAwsConfig(*settings)
	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &databasemigrationservice.StartReplicationTaskInput{
		ReplicationTaskArn:       settings.Arn,
		StartReplicationTaskType: types.StartReplicationTaskTypeValue("resume-processing"),
	}

	resp, err := awsConfig.DmsClient().StartReplicationTask(context.TODO(), params)

	if err != nil {
		fmt.Printf("failed to start dms task, %v", err)
		os.Exit(1)
	}
	fmt.Printf("starting dms replication task: %s ", *resp.ReplicationTask.ReplicationTaskIdentifier)
}
