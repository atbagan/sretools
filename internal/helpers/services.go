package helpers

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/spf13/viper"
	"os"
)

// TargetGroups struct for tgs being returned
type TargetGroups struct {
	TargetGroup []string
}

// GetServices return tgs
func GetServices(svc *ecs.Client) (*TargetGroups, error) {

	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &ecs.ListServicesInput{
		Cluster: &configuration.Ecs.Cluster,
	}

	result, err := svc.ListServices(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}

	tgs := GetTargetGroupArn(svc, result.ServiceArns)
	return tgs, nil
}

// GetTargetGroupArn return tg arns
func GetTargetGroupArn(svc *ecs.Client, va []string) *TargetGroups {
	var tgs TargetGroups

	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &ecs.DescribeServicesInput{
		Services: va,
		Cluster:  &configuration.Ecs.Cluster,
	}
	result, err := svc.DescribeServices(context.TODO(), params)

	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}

	for _, va := range result.Services {
		if &va.LoadBalancers != nil {
			for _, va := range va.LoadBalancers {
				tgs.TargetGroup = append(tgs.TargetGroup, *va.TargetGroupArn)
			}
		}
	}
	return &tgs
}
