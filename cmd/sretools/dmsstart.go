// Package sretools
// Copyright (c) 2022, Andrew Bagan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
