// Package helpers
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

// GetEventbridgeRules gets the NLB arns and returns them
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
