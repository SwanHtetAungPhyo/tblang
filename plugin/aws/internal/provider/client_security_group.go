package provider

import (
"fmt"
"context"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateSecurityGroup(ctx context.Context, vpcID, name, description string, tags map[string]string) (*SecurityGroupResult, error) {
	input := &ec2.CreateSecurityGroupInput{
		VpcId:       aws.String(vpcID),
		GroupName:   aws.String(name),
		Description: aws.String(description),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSecurityGroup,
				Tags:         c.buildTags("security-group", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateSecurityGroup(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create Security Group: %w", err)
	}

	if result.GroupId == nil {
		return nil, fmt.Errorf("Security Group creation returned nil")
	}

	return &SecurityGroupResult{
		GroupID: *result.GroupId,
	}, nil
}

// AuthorizeSecurityGroupIngress adds ingress rules to a security group
func (c *AWSClient) AuthorizeSecurityGroupIngress(ctx context.Context, groupID string, rules []types.IpPermission) error {
	if len(rules) == 0 {
		return nil
	}

	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       aws.String(groupID),
		IpPermissions: rules,
	}

	_, err := c.EC2Client.AuthorizeSecurityGroupIngress(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to authorize ingress rules: %w", err)
	}

	return nil
}

// AuthorizeSecurityGroupEgress adds egress rules to a security group
func (c *AWSClient) AuthorizeSecurityGroupEgress(ctx context.Context, groupID string, rules []types.IpPermission) error {
	if len(rules) == 0 {
		return nil
	}

	input := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId:       aws.String(groupID),
		IpPermissions: rules,
	}

	_, err := c.EC2Client.AuthorizeSecurityGroupEgress(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to authorize egress rules: %w", err)
	}

	return nil
}

// DeleteSecurityGroup deletes a security group
func (c *AWSClient) DeleteSecurityGroup(ctx context.Context, groupID string) error {
	input := &ec2.DeleteSecurityGroupInput{
		GroupId: aws.String(groupID),
	}

	_, err := c.EC2Client.DeleteSecurityGroup(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete Security Group %s: %w", groupID, err)
	}

	return nil
}

// DeleteSubnet deletes a subnet
func (c *AWSClient) DeleteSubnet(ctx context.Context, subnetID string) error {
	input := &ec2.DeleteSubnetInput{
		SubnetId: aws.String(subnetID),
	}

	_, err := c.EC2Client.DeleteSubnet(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete Subnet %s: %w", subnetID, err)
	}

	return nil
}

// buildTags creates AWS tags from a map

// EC2 Instance types and methods



