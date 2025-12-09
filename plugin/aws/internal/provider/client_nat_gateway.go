package provider

import (
"context"
"fmt"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateNATGateway(ctx context.Context, subnetID, allocationID string, tags map[string]string) (*NATGatewayResult, error) {
	input := &ec2.CreateNatGatewayInput{
		SubnetId:     aws.String(subnetID),
		AllocationId: aws.String(allocationID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeNatgateway,
				Tags:         c.buildTags("nat", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateNatGateway(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create NAT gateway: %w", err)
	}

	natGatewayID := *result.NatGateway.NatGatewayId

	// Wait for NAT gateway to be available
	waiter := ec2.NewNatGatewayAvailableWaiter(c.EC2Client)
	err = waiter.Wait(ctx, &ec2.DescribeNatGatewaysInput{
		NatGatewayIds: []string{natGatewayID},
	}, 10*60) // 10 minute timeout

	if err != nil {
		fmt.Printf("Warning: NAT gateway may not be fully available: %v\n", err)
	}

	return &NATGatewayResult{
		NATGatewayID: natGatewayID,
		State:        string(result.NatGateway.State),
	}, nil
}

// DeleteNATGateway deletes a NAT gateway
func (c *AWSClient) DeleteNATGateway(ctx context.Context, natGatewayID string) error {
	_, err := c.EC2Client.DeleteNatGateway(ctx, &ec2.DeleteNatGatewayInput{
		NatGatewayId: aws.String(natGatewayID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete NAT gateway: %w", err)
	}

	// Wait for deletion
	waiter := ec2.NewNatGatewayDeletedWaiter(c.EC2Client)
	err = waiter.Wait(ctx, &ec2.DescribeNatGatewaysInput{
		NatGatewayIds: []string{natGatewayID},
	}, 10*60)

	if err != nil {
		fmt.Printf("Warning: NAT gateway may not be fully deleted: %v\n", err)
	}

	return nil
}

// Data Source types and methods



