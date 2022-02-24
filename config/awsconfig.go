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
)

//AWSConfig is a holder for AWS Config type information
type AWSConfig struct {
	Config aws.Config
}

// DefaultAwsConfig aws config
func DefaultAwsConfig(config Config) AWSConfig {
	awsConfig := AWSConfig{}

	if *config.Iam != "" {
		fmt.Println("IAM Role")
		cfg, err := external.LoadDefaultConfig(context.TODO(), external.WithDefaultRegion(*config.Region))
		if err != nil {
			panic(err)
		}
		stsClient := sts.NewFromConfig(cfg)
		creds := stscreds.NewAssumeRoleProvider(stsClient, *config.Iam)
		cfg.Credentials = aws.NewCredentialsCache(creds)
		awsConfig.Config = cfg

	} else if *config.Profile != "" {
		fmt.Println("Profile")
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

func AssumeIamRole(config Config) AWSConfig {
	awsConfig := AWSConfig{}
	//input := &sts.AssumeRoleInput{
	//	RoleArn:         config.Iam,
	//	RoleSessionName: aws.String("assume_role_session" ),
	//}
	//result, err := TakeRole(context.TODO(), stsClient, input)

	//appCreds := stscreds.NewAssumeRoleProvider(stsClient, *config.Iam)
	//value, err := appCreds.Retrieve(context.TODO())
	//awsConfig.Credentials = value
	//if err != nil {
	//	panic(err)
	//}
	//if *config.Profile != "" {
	//	fmt.Println("profile")
	//	cfg, err := external.LoadDefaultConfig(context.TODO(), external.WithSharedConfigProfile(*config.Profile))
	//	if err != nil {
	//		panic(err)
	//	}
	//	awsConfig.Config = cfg
	//}
	//} else {
	//	cfg, err := external.LoadDefaultConfig(context.TODO())
	//	if err != nil {
	//		panic(err)
	//	}
	//	awsConfig.Config = cfg
	//}
	//if *config.Region != "" {
	//	awsConfig.Config.Region = *config.Region
	//}
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
