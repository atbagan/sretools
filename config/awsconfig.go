package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	external "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

//AWSConfig is a holder for AWS Config type information
type AWSConfig struct {
	Config aws.Config
}

// DefaultAwsConfig aws config
func DefaultAwsConfig(config Config) AWSConfig {
	awsConfig := AWSConfig{}

	if *config.Profile != "" {
		cfg, err := external.LoadDefaultConfig(context.TODO(), external.WithSharedConfigProfile(*config.Profile))
		if err != nil {
			panic(err)
		}
		awsConfig.Config = cfg
	} else {
		cfg, err := external.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic(err)
		}
		awsConfig.Config = cfg
	}
	if *config.Region != "" {
		awsConfig.Config.Region = *config.Region
	}
	return awsConfig

}

//Ec2Client returns an ec2 Client
func (config *AWSConfig) Ec2Client() *ec2.Client {
	return ec2.NewFromConfig(config.Config)
}

//ElbClient returns an ec2 Client
func (config *AWSConfig) ElbClient() *elasticloadbalancingv2.Client {
	return elasticloadbalancingv2.NewFromConfig(config.Config)
}

//EcsClient returns an ECS Client
func (config *AWSConfig) EcsClient() *ecs.Client {
	return ecs.NewFromConfig(config.Config)
}

//EventbridgeClient returns and Eventbridge Client
func (config *AWSConfig) EventbridgeClient() *eventbridge.Client {
	return eventbridge.NewFromConfig(config.Config)
}

//DmsClient returns a databasemigrationservice client
func (config *AWSConfig) DmsClient() *databasemigrationservice.Client {
	return databasemigrationservice.NewFromConfig(config.Config)
}

//CdClient returns codedeploy client
func (config *AWSConfig) CdClient() *codedeploy.Client {
	return codedeploy.NewFromConfig(config.Config)
}
