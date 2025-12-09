package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateVPC(ctx context.Context, cidrBlock string, tags map[string]string) (*VPCResult, error) {
	input := &ec2.CreateVpcInput{
		CidrBlock: aws.String(cidrBlock),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeVpc,
				Tags:         c.buildTags("vpc", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateVpc(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create VPC: %w", err)
	}

	if result.Vpc == nil {
		return nil, fmt.Errorf("VPC creation returned nil")
	}

	// Enable DNS hostnames and support
	vpcID := *result.Vpc.VpcId
	if err := c.enableVPCDNS(ctx, vpcID); err != nil {
		// Log warning but don't fail
		fmt.Printf("Warning: failed to enable DNS for VPC %s: %v\n", vpcID, err)
	}

	return &VPCResult{
		VPCID: vpcID,
		State: string(result.Vpc.State),
	}, nil
}

// enableVPCDNS enables DNS hostnames and support for a VPC
func (c *AWSClient) enableVPCDNS(ctx context.Context, vpcID string) error {
	// Enable DNS hostnames
	_, err := c.EC2Client.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:              aws.String(vpcID),
		EnableDnsHostnames: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS hostnames: %w", err)
	}

	// Enable DNS support
	_, err = c.EC2Client.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:           aws.String(vpcID),
		EnableDnsSupport: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS support: %w", err)
	}

	return nil
}

// SubnetResult represents the result of Subnet creation

// CreateSubnet creates a subnet and returns the result

// ConfigureSubnetPublicIP configures public IP mapping for a subnet

// DeleteVPC deletes a VPC
