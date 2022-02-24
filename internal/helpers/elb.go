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
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"os"
)

// GetNlbLoadBalancerArns gets the NLB arns and returns them
func GetNlbLoadBalancerArns(svc *elasticloadbalancingv2.Client) []string {
	params := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
	result, err := svc.DescribeLoadBalancers(context.TODO(), params)

	arns := make([]string, 0)

	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}

	for _, v := range result.LoadBalancers {
		if v.Type == "network" {
			arns = append(arns, *v.LoadBalancerArn)
		}
	}
	return arns
}

// GetAlbLoadBalancerArns gets the NLB arns and returns them
func GetAlbLoadBalancerArns(svc *elasticloadbalancingv2.Client) []string {
	params := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
	result, err := svc.DescribeLoadBalancers(context.TODO(), params)

	arns := make([]string, 0)

	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}

	for _, v := range result.LoadBalancers {
		if v.Type == "application" {
			arns = append(arns, *v.LoadBalancerArn)
		}
	}
	return arns
}

// GetAllLoadBalancerArns gets the NLB arns and returns them
func GetAllLoadBalancerArns(svc *elasticloadbalancingv2.Client) []string {
	params := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
	result, err := svc.DescribeLoadBalancers(context.TODO(), params)

	arns := make([]string, 0)

	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}

	for _, v := range result.LoadBalancers {
		arns = append(arns, *v.LoadBalancerArn)
	}
	return arns
}
