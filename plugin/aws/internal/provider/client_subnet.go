package provider

import (
"fmt"
"context"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateSubnet(ctx context.Context, vpcID, cidrBlock, availabilityZone string, tags map[string]string) (*SubnetResult, error) {
	input := &ec2.CreateSubnetInput{
		VpcId:            aws.String(vpcID),
		CidrBlock:        aws.String(cidrBlock),
		AvailabilityZone: aws.String(availabilityZone),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags:         c.buildTags("subnet", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateSubnet(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create Subnet: %w", err)
	}

	if result.Subnet == nil {
		return nil, fmt.Errorf("Subnet creation returned nil")
	}

	return &SubnetResult{
		SubnetID: *result.Subnet.SubnetId,
		State:    string(result.Subnet.State),
	}, nil
}

func (c *AWSClient) ConfigureSubnetPublicIP(ctx context.Context, subnetID string, mapPublicIP bool) error {
	input := &ec2.ModifySubnetAttributeInput{
		SubnetId: aws.String(subnetID),
		MapPublicIpOnLaunch: &types.AttributeBooleanValue{
			Value: aws.Bool(mapPublicIP),
		},
	}

	_, err := c.EC2Client.ModifySubnetAttribute(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to configure public IP mapping: %w", err)
	}

	return nil
}

func (c *AWSClient) DeleteVPC(ctx context.Context, vpcID string) error {
	input := &ec2.DeleteVpcInput{
		VpcId: aws.String(vpcID),
	}

	_, err := c.EC2Client.DeleteVpc(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete VPC %s: %w", vpcID, err)
	}

	return nil
}
