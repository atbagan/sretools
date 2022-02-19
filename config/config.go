package config

import (
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

// Config holds the global configuration settings
type Config struct {
	Verbose     *bool
	Profile     *string
	Region      *string
	NameFile    *string
	Eps         EpsConfiguration
	Eventbridge EventbridgeConfiguration
	Ecs         EcsConfiguration
	Codedeploy  CodedeployConfiguration

	ErrorCode bool
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
	Revision                      Revision
	TargetInstances               *types.TargetInstances
	UpdateOutdatedInstancesOnly   bool
	Bucket                        string
	Key                           string
	Etag                          string
	Version                       string
}

type Revision struct {
	RevisionType   string `json:"revisionType"`
	AppSpecContent struct {
		Content string `json:"content"`
		Sha256  string `json:"sha256"`
	} `json:"appSpecContent"`
}
