// Package config
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

package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	external "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

// AWSConfig is a holder for AWS Config type information
type AWSConfig struct {
	Config aws.Config
}

// DefaultAwsConfig returns the aws config struct credentials
func DefaultAwsConfig(config Config) AWSConfig {
	awsConfig := AWSConfig{}
	switch {
	case *config.Iam != "":
		fmt.Println("IAM Role")
		cfg, err := external.LoadDefaultConfig(context.TODO(), external.WithDefaultRegion(*config.Region))
		if err != nil {
			log.Fatalf("Error loading credentials: %v", err)
		}
		stsClient := sts.NewFromConfig(cfg)
		creds := stscreds.NewAssumeRoleProvider(stsClient, *config.Iam)
		cfg.Credentials = aws.NewCredentialsCache(creds)
		awsConfig.Config = cfg
	case *config.Profile != "":
		fmt.Println("Profile")
		cfg, err := external.LoadDefaultConfig(context.TODO(), external.WithSharedConfigProfile(*config.Profile))
		if err != nil {
			log.Fatalf("Error loading credentials: %v", err)
		}
		awsConfig.Config = cfg
	default:
		cfg, err := external.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Error loading credentials: %v", err)
		}
		awsConfig.Config = cfg

	}
	if *config.Region != "" {
		awsConfig.Config.Region = *config.Region
	}
	return awsConfig
}

// Ec2Client returns an ec2 Client
func (config *AWSConfig) Ec2Client() *ec2.Client {
	return ec2.NewFromConfig(config.Config)
}

// ElbClient returns an ec2 Client
func (config *AWSConfig) ElbClient() *elasticloadbalancingv2.Client {
	return elasticloadbalancingv2.NewFromConfig(config.Config)
}

// EcsClient returns an ECS Client
func (config *AWSConfig) EcsClient() *ecs.Client {
	return ecs.NewFromConfig(config.Config)
}

// EventbridgeClient returns and Eventbridge Client
func (config *AWSConfig) EventbridgeClient() *eventbridge.Client {
	return eventbridge.NewFromConfig(config.Config)
}

// DmsClient returns a databasemigrationservice client
func (config *AWSConfig) DmsClient() *databasemigrationservice.Client {
	return databasemigrationservice.NewFromConfig(config.Config)
}

// CdClient returns codedeploy client
func (config *AWSConfig) CdClient() *codedeploy.Client {
	return codedeploy.NewFromConfig(config.Config)
}
