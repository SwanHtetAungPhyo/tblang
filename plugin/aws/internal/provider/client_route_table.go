package provider

import (
"fmt"
"context"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) CreateRouteTable(ctx context.Context, vpcID string, tags map[string]string) (*RouteTableResult, error) {
	input := &ec2.CreateRouteTableInput{
		VpcId: aws.String(vpcID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags:         c.buildTags("rtb", tags),
			},
		},
	}

	result, err := c.EC2Client.CreateRouteTable(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create route table: %w", err)
	}

	return &RouteTableResult{
		RouteTableID: *result.RouteTable.RouteTableId,
	}, nil
}

// CreateRoute creates a route in a route table
func (c *AWSClient) CreateRoute(ctx context.Context, routeTableID, destCIDR, gatewayID, natGatewayID string) error {
	input := &ec2.CreateRouteInput{
		RouteTableId:         aws.String(routeTableID),
		DestinationCidrBlock: aws.String(destCIDR),
	}

	if gatewayID != "" {
		input.GatewayId = aws.String(gatewayID)
	}
	if natGatewayID != "" {
		input.NatGatewayId = aws.String(natGatewayID)
	}

	_, err := c.EC2Client.CreateRoute(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create route: %w", err)
	}

	return nil
}

// DeleteRouteTable deletes a route table
func (c *AWSClient) DeleteRouteTable(ctx context.Context, routeTableID string) error {
	_, err := c.EC2Client.DeleteRouteTable(ctx, &ec2.DeleteRouteTableInput{
		RouteTableId: aws.String(routeTableID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete route table: %w", err)
	}

	return nil
}

// EIP types and methods


