package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateEC2Instance(ctx context.Context, config *EC2InstanceConfig) (*EC2InstanceResult, error) {
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(config.AMI),
		InstanceType: types.InstanceType(config.InstanceType),
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		SubnetId:     aws.String(config.SubnetID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags:         c.buildTags("ec2", config.Tags),
			},
		},
	}

	if len(config.SecurityGroupIDs) > 0 {
		input.SecurityGroupIds = config.SecurityGroupIDs
	}

	if config.KeyName != "" {
		input.KeyName = aws.String(config.KeyName)
	}

	if config.UserData != "" {
		encoded := base64.StdEncoding.EncodeToString([]byte(config.UserData))
		input.UserData = aws.String(encoded)
	}

	if config.AssociatePublicIP {
		input.NetworkInterfaces = []types.InstanceNetworkInterfaceSpecification{
			{
				DeviceIndex:              aws.Int32(0),
				SubnetId:                 aws.String(config.SubnetID),
				AssociatePublicIpAddress: aws.Bool(true),
				Groups:                   config.SecurityGroupIDs,
			},
		}

		input.SubnetId = nil
		input.SecurityGroupIds = nil
	}

	if config.RootVolumeSize > 0 {
		input.BlockDeviceMappings = []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/xvda"),
				Ebs: &types.EbsBlockDevice{
					VolumeSize:          aws.Int32(config.RootVolumeSize),
					VolumeType:          types.VolumeType(config.RootVolumeType),
					DeleteOnTermination: aws.Bool(true),
				},
			},
		}
	}

	result, err := c.EC2Client.RunInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create EC2 instance: %w", err)
	}

	if len(result.Instances) == 0 {
		return nil, fmt.Errorf("no instances created")
	}

	instance := result.Instances[0]

	waiter := ec2.NewInstanceRunningWaiter(c.EC2Client)
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{*instance.InstanceId},
	}, 5*60)

	if err != nil {
		fmt.Printf("Warning: instance may not be fully running: %v\n", err)
	}

	descResult, err := c.EC2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{*instance.InstanceId},
	})

	var publicIP, privateIP, state string
	if err == nil && len(descResult.Reservations) > 0 && len(descResult.Reservations[0].Instances) > 0 {
		inst := descResult.Reservations[0].Instances[0]
		if inst.PublicIpAddress != nil {
			publicIP = *inst.PublicIpAddress
		}
		if inst.PrivateIpAddress != nil {
			privateIP = *inst.PrivateIpAddress
		}
		state = string(inst.State.Name)
	}

	return &EC2InstanceResult{
		InstanceID: *instance.InstanceId,
		PublicIP:   publicIP,
		PrivateIP:  privateIP,
		State:      state,
	}, nil
}

func (c *AWSClient) TerminateEC2Instance(ctx context.Context, instanceID string) error {
	_, err := c.EC2Client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return fmt.Errorf("failed to terminate instance %s: %w", instanceID, err)
	}

	fmt.Printf("  Waiting for instance %s to terminate...\n", instanceID)

	waiter := ec2.NewInstanceTerminatedWaiter(c.EC2Client)
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}, 5*time.Minute)

	if err != nil {

		fmt.Printf("  Waiter timeout, polling for termination status...\n")
		for i := 0; i < 30; i++ {
			time.Sleep(2 * time.Second)
			result, descErr := c.EC2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
				InstanceIds: []string{instanceID},
			})
			if descErr != nil {
				continue
			}
			if len(result.Reservations) > 0 && len(result.Reservations[0].Instances) > 0 {
				state := result.Reservations[0].Instances[0].State.Name
				if state == types.InstanceStateNameTerminated {
					fmt.Printf("  Instance %s terminated successfully\n", instanceID)
					return nil
				}
				fmt.Printf("  Instance state: %s\n", state)
			}
		}
		return fmt.Errorf("timeout waiting for instance %s to terminate", instanceID)
	}

	fmt.Printf("  Instance %s terminated successfully\n", instanceID)
	return nil
}
