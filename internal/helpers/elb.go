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
