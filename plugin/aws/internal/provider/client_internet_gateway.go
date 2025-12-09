package provider

import (
"fmt"
"context"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateInternetGateway(ctx context.Context, vpcID string, tags map[string]string) (*InternetGatewayResult, error) {
	input := &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags:         c.buildTags("igw", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateInternetGateway(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create internet gateway: %w", err)
	}

	gatewayID := *result.InternetGateway.InternetGatewayId

	// Attach to VPC
	_, err = c.EC2Client.AttachInternetGateway(ctx, &ec2.AttachInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
		VpcId:             aws.String(vpcID),
	})
	if err != nil {
		// Try to delete the gateway if attach fails
		c.EC2Client.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
			InternetGatewayId: aws.String(gatewayID),
		})
		return nil, fmt.Errorf("failed to attach internet gateway to VPC: %w", err)
	}

	return &InternetGatewayResult{
		GatewayID: gatewayID,
	}, nil
}

// DeleteInternetGateway detaches and deletes an internet gateway
func (c *AWSClient) DeleteInternetGateway(ctx context.Context, gatewayID, vpcID string) error {
	// Detach from VPC first
	if vpcID != "" {
		_, err := c.EC2Client.DetachInternetGateway(ctx, &ec2.DetachInternetGatewayInput{
			InternetGatewayId: aws.String(gatewayID),
			VpcId:             aws.String(vpcID),
		})
		if err != nil {
			return fmt.Errorf("failed to detach internet gateway: %w", err)
		}
	}

	// Delete the gateway
	_, err := c.EC2Client.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete internet gateway: %w", err)
	}

	return nil
}

// Route Table types and methods


