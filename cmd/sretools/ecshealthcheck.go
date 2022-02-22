package sretools

import (
	"context"
	"fmt"
	c "github.com/atbagan/sretools/config"
	"github.com/atbagan/sretools/internal/helpers"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var healthCmd = &cobra.Command{
	Use:   "health-check",
	Short: "Health Check for ECS Service",
	Long:  "Checks health for a given target group to determine if your service is healthy or not",
	Run:   getHealthCheck,
}

func init() {
	ecsCmd.AddCommand(healthCmd)
}

//List struct for arn and current health status
type List struct {
	ARN    string
	Status string
}

// AtomicInt state variable
type AtomicInt struct {
	mu sync.Mutex // A lock than can be held by one goroutine at a time.
	n  int
}

var wg = sync.WaitGroup{}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add() {
	a.mu.Lock() // Wait for the lock to be free and then take it.
	a.n++
	//fmt.Println("mutex Var:", a.n)
	a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func getHealthCheck(cmd *cobra.Command, args []string) {
	awsConfig := c.DefaultAwsConfig(*settings)

	var configuration c.Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	tgs, err := helpers.GetServices(awsConfig.EcsClient())
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	numOfTargets := len(tgs.TargetGroup)

	var n AtomicInt

	ch := make(chan *List, numOfTargets)

	wg.Add(numOfTargets)
	for i := 0; i < numOfTargets; i++ {
		go GetTargetHealth(tgs.TargetGroup[i], &n, ch, awsConfig.ElbClient())
	}

	go func(ch chan<- *List) {
		defer close(ch)
		wg.Wait()
		fmt.Println("")
	}(ch)

	for i := range ch {
		fmt.Println(i.ARN)
	}

	select {
	case <-ch:
		if n.Value() == numOfTargets {
			fmt.Println("(4/4) All target groups are healthy")
			os.Exit(0)
		} else {
			fmt.Println("(4/4) Not all target groups are healthy. Please log in to your AWS console to verify")
			if configuration.ErrorCode {
				os.Exit(1)
			}
		}
	case <-time.After(5 * time.Second):
		fmt.Println("(4/4) TIMED OUT")
		os.Exit(1)
	}
}

//GetTargetHealth get tg health
func GetTargetHealth(arn string, n *AtomicInt, ch chan<- *List, svc *elasticloadbalancingv2.Client) {
	defer wg.Done()
	var params elasticloadbalancingv2.DescribeTargetHealthInput

	var listing List
	listing.ARN = arn
	listing.Status = "unhealthy"
	attempt := 0
	params.TargetGroupArn = &arn

	for attempt < 5 {

		result, err := svc.DescribeTargetHealth(context.TODO(), &params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, v := range result.TargetHealthDescriptions {
			if v.TargetHealth.State == "healthy" {
				listing.Status = "healthy"
				break
			}
		}
		if listing.Status == "healthy" {
			n.Add()
			break
		} else {
			time.Sleep(time.Duration(10) * time.Second)
			attempt++
		}

	}
	ch <- &listing
}
