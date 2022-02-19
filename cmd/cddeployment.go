package cmd

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	cd "github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var deploymentCmd = &cobra.Command{
Use:   "deployment",
Short: "Codedeploy Deployment",
Long:  "Create a codedeploy deployment",
Run:  createDeployment,
}

func init() {
	codedeployCmd.AddCommand(deploymentCmd)
}

func createDeployment(cmd *cobra.Command, args []string)  {
	awsConfig := c.DefaultAwsConfig(*settings)
	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &cd.CreateDeploymentInput{
		ApplicationName:               aws.String(configuration.Codedeploy.ApplicationName),
		AutoRollbackConfiguration:     nil,
		DeploymentConfigName:          aws.String(configuration.Codedeploy.DeploymentConfigName),
		DeploymentGroupName:           aws.String(configuration.Codedeploy.DeploymentGroupName),
		Description:                   aws.String(configuration.Codedeploy.Description),
		FileExistsBehavior:            "",
		IgnoreApplicationStopFailures: false,
		Revision:                      &types.RevisionLocation{
			S3Location: &types.S3Location{
				Bucket: aws.String(configuration.Codedeploy.Bucket),
				ETag: aws.String("ETag"),
				Key: aws.String(configuration.Codedeploy.Key),
				Version: aws.String("VersionID"),
			},
		},
		TargetInstances:               nil,
		UpdateOutdatedInstancesOnly:   false,
	}
	_ , err = awsConfig.CdClient().CreateDeployment(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		fmt.Println("failed here")
		os.Exit(1)
	}
}

//func createDeploymentConfig(cmd *cobra.Command, args []string)  {
//	awsConfig := c.DefaultAwsConfig(*settings)
//	var configuration c.Config
//	err := viper.Unmarshal(&configuration)
//	if err != nil {
//		fmt.Printf( "Unable to decode into struct, %v", err)
//	}
//
//	params := &cd.CreateDeploymentConfigInput{
//		DeploymentConfigName: nil,
//		ComputePlatform:      "",
//		MinimumHealthyHosts:  nil,
//		TrafficRoutingConfig: nil,
//	}
//
//}