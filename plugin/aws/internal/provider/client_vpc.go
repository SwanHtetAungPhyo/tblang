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

	vpcID := *result.Vpc.VpcId
	if err := c.enableVPCDNS(ctx, vpcID); err != nil {

		fmt.Printf("Warning: failed to enable DNS for VPC %s: %v\n", vpcID, err)
	}

	return &VPCResult{
		VPCID: vpcID,
		State: string(result.Vpc.State),
	}, nil
}

func (c *AWSClient) enableVPCDNS(ctx context.Context, vpcID string) error {

	_, err := c.EC2Client.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:              aws.String(vpcID),
		EnableDnsHostnames: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS hostnames: %w", err)
	}

	_, err = c.EC2Client.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:           aws.String(vpcID),
		EnableDnsSupport: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS support: %w", err)
	}

	return nil
}
