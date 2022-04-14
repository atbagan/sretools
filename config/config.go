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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

// Config holds the global configuration settings
type Config struct {
	Verbose     *bool
	Profile     *string
	Iam         *string
	Region      *string
	NameFile    *string
	Dms         DmsConfiguration
	Eps         EpsConfiguration
	Eventbridge EventbridgeConfiguration
	Ecs         EcsConfiguration
	Codedeploy  CodedeployConfiguration
	Credentials aws.CredentialsProvider
	ErrorCode   bool
}

// EpsConfiguration config struct that holds config values for EPS
type EpsConfiguration struct {
	Serviceid string
}

// EventbridgeConfiguration config struct that holds config values for Eventbridge
type EventbridgeConfiguration struct {
	Nameprefix string
}

// EcsConfiguration config struct that holds config values for ECS
type EcsConfiguration struct {
	Cluster string
}

// CodedeployConfiguration config struct that holds config values for codedeploy
type CodedeployConfiguration struct {
	ApplicationName               string
	AutoRollbackConfiguration     *types.AutoRollbackConfiguration
	DeploymentConfigName          string
	DeploymentGroupName           string
	Description                   string
	FileExistsBehavior            types.FileExistsBehavior
	IgnoreApplicationStopFailures bool
	Revision                      *types.RevisionLocation
	TargetInstances               *types.TargetInstances
	UpdateOutdatedInstancesOnly   bool
	Bucket                        string
	Key                           string
	Etag                          string
	Version                       string
}

type DmsConfiguration struct {
	TaskArn string
}
