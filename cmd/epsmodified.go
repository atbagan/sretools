package cmd

import (
	"automation/config"
	"automation/helpers"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
	"os"
)

var associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Modify EPS",
	Long:  "Modify Vpc EPS to associate LB with EPS",
	Run:   associateLoadBalancerEps,
}

func init() {
	epsCmd.AddCommand(associateCmd)
}

//associateVpcEndpointService modifies the endpoint service
func associateLoadBalancerEps(cmd *cobra.Command, args []string) {
	awsConfig := config.DefaultAwsConfig(*settings)
	nlbArns := helpers.GetNlbLoadBalancerArns(awsConfig.ElbClient())
	params := &ec2.ModifyVpcEndpointServiceConfigurationInput{
		ServiceId:                  aws.String(settings.Eps.Serviceid),
		AddNetworkLoadBalancerArns: nlbArns,
	}
	_, err := awsConfig.Ec2Client().ModifyVpcEndpointServiceConfiguration(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		os.Exit(1)
	}
	fmt.Printf("Associated Load Balancer(s): %s", nlbArns[0])
}

//func disassociateLoadBalancerEps(cmd *cobra.Command, args []string) {
//
//}
