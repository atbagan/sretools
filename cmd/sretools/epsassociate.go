package sretools

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/atbagan/sretools/internal/helpers"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate LB EPS",
	Long:  "Modify Vpc EPS to associate a LB with an endpoint service",
	Run:   associateLoadBalancerEps,
}

func init() {
	epsCmd.AddCommand(associateCmd)
}

//associateLoadBalancerEps associates NLB's with the given endpoint service
func associateLoadBalancerEps(cmd *cobra.Command, args []string) {
	awsConfig := c.DefaultAwsConfig(*settings)
	nlbArns := helpers.GetNlbLoadBalancerArns(awsConfig.ElbClient())

	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	params := &ec2.ModifyVpcEndpointServiceConfigurationInput{
		ServiceId:                  &configuration.Eps.Serviceid,
		AddNetworkLoadBalancerArns: nlbArns,
	}

	_, err = awsConfig.Ec2Client().ModifyVpcEndpointServiceConfiguration(context.TODO(), params)
	if err != nil {
		fmt.Sprintf("failed to load the config, %v", err)
		fmt.Println("failed here")
		os.Exit(1)
	}
	fmt.Printf("Associated Load Balancer(s): %s", nlbArns[0])
}

//func disassociateLoadBalancerEps(cmd *cobra.Command, args []string) {
//
//}
