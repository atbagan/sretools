// Package sretools
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

// associateLoadBalancerEps associates NLB's with the given endpoint service
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
