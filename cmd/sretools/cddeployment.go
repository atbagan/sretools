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
	"github.com/aws/aws-sdk-go-v2/aws"
	cd "github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Codedeploy Deployment",
	Long:  "Create a codedeploy deployment",
	Run:   createDeployment,
}

func init() {
	codedeployCmd.AddCommand(deploymentCmd)
}

func createDeployment(cmd *cobra.Command, args []string) {
	awsConfig := c.DefaultAwsConfig(*settings)
	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Error("Unable to decode into struct")

	}

	params := &cd.CreateDeploymentInput{
		ApplicationName: aws.String(configuration.Codedeploy.ApplicationName),
		AutoRollbackConfiguration: &types.AutoRollbackConfiguration{
			Enabled: true,
			Events: []types.AutoRollbackEvent{
				"DEPLOYMENT_FAILURE",
				"DEPLOYMENT_STOP_ON_ALARM",
				"DEPLOYMENT_STOP_ON_REQUEST",
			},
		},
		DeploymentConfigName:          aws.String(configuration.Codedeploy.DeploymentConfigName),
		DeploymentGroupName:           aws.String(configuration.Codedeploy.DeploymentGroupName),
		Description:                   aws.String(configuration.Codedeploy.Description),
		FileExistsBehavior:            "",
		IgnoreApplicationStopFailures: false,
		Revision: &types.RevisionLocation{
			RevisionType: "S3",
			S3Location: &types.S3Location{
				Bucket:     aws.String(configuration.Codedeploy.Bucket),
				ETag:       aws.String(configuration.Codedeploy.Etag),
				Key:        aws.String(configuration.Codedeploy.Key),
				Version:    aws.String(configuration.Codedeploy.Version),
				BundleType: "JSON",
			},
		},
		TargetInstances:             nil,
		UpdateOutdatedInstancesOnly: false,
	}
	result, err := awsConfig.CdClient().CreateDeployment(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		log.Fatal(err)
	}
	fmt.Println("deployment in progress with deployment id: ", &result.DeploymentId)
}
