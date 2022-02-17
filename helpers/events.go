package helpers

import (
	"context"
	"fmt"
	"github.com/atbagan/sretools/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"os"
)

var settings = new(config.Config)

//GetEventbridgeRules gets the NLB arns and returns them
func GetEventbridgeRules(svc *eventbridge.Client) []string {
	params := &eventbridge.ListRulesInput{
		NamePrefix: aws.String(settings.Eventbridge.Nameprefix),
	}
	result, err := svc.ListRules(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}
	rules := make([]string, 0)

	for _, v := range result.Rules {
		if *v.ScheduleExpression != "" {
			rules = append(rules, *v.Name)
		}
	}
	return rules
}
