package provider

import (
"context"
"fmt"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
"github.com/aws/aws-sdk-go-v2/service/sts"
)
func (c *AWSClient) DescribeAMI(ctx context.Context, owners []string, filters []AMIFilter, mostRecent bool) (*AMIResult, error) {
	input := &ec2.DescribeImagesInput{
		Owners: owners,
	}

	// Add filters
	for _, f := range filters {
		input.Filters = append(input.Filters, types.Filter{
			Name:   aws.String(f.Name),
			Values: f.Values,
		})
	}

	result, err := c.EC2Client.DescribeImages(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe AMIs: %w", err)
	}

	if len(result.Images) == 0 {
		return nil, fmt.Errorf("no AMIs found matching criteria")
	}

	// Sort by creation date if most_recent is true
	images := result.Images
	if mostRecent && len(images) > 1 {
		// Simple sort - find the most recent
		var mostRecentImage types.Image
		var mostRecentDate string
		for _, img := range images {
			if img.CreationDate != nil && *img.CreationDate > mostRecentDate {
				mostRecentDate = *img.CreationDate
				mostRecentImage = img
			}
		}
		if mostRecentImage.ImageId != nil {
			images = []types.Image{mostRecentImage}
		}
	}

	img := images[0]
	var name, arch string
	if img.Name != nil {
		name = *img.Name
	}
	if img.Architecture != "" {
		arch = string(img.Architecture)
	}

	return &AMIResult{
		AMIID:        *img.ImageId,
		Name:         name,
		Architecture: arch,
	}, nil
}


// DescribeVPC finds a VPC
func (c *AWSClient) DescribeVPC(ctx context.Context, vpcID string, isDefault bool) (*VPCDataResult, error) {
	input := &ec2.DescribeVpcsInput{}

	if vpcID != "" {
		input.VpcIds = []string{vpcID}
	}

	if isDefault {
		input.Filters = []types.Filter{
			{
				Name:   aws.String("isDefault"),
				Values: []string{"true"},
			},
		}
	}

	result, err := c.EC2Client.DescribeVpcs(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe VPCs: %w", err)
	}

	if len(result.Vpcs) == 0 {
		return nil, fmt.Errorf("no VPCs found")
	}

	vpc := result.Vpcs[0]
	return &VPCDataResult{
		VPCID:     *vpc.VpcId,
		CIDRBlock: *vpc.CidrBlock,
		State:     string(vpc.State),
	}, nil
}


// DescribeSubnet finds a subnet
func (c *AWSClient) DescribeSubnet(ctx context.Context, subnetID, vpcID string) (*SubnetDataResult, error) {
	input := &ec2.DescribeSubnetsInput{}

	if subnetID != "" {
		input.SubnetIds = []string{subnetID}
	}

	if vpcID != "" {
		input.Filters = []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
		}
	}

	result, err := c.EC2Client.DescribeSubnets(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe subnets: %w", err)
	}

	if len(result.Subnets) == 0 {
		return nil, fmt.Errorf("no subnets found")
	}

	subnet := result.Subnets[0]
	return &SubnetDataResult{
		SubnetID:         *subnet.SubnetId,
		CIDRBlock:        *subnet.CidrBlock,
		AvailabilityZone: *subnet.AvailabilityZone,
	}, nil
}


// DescribeAvailabilityZones lists availability zones
func (c *AWSClient) DescribeAvailabilityZones(ctx context.Context, state string) (*AvailabilityZonesResult, error) {
	input := &ec2.DescribeAvailabilityZonesInput{}

	if state != "" {
		input.Filters = []types.Filter{
			{
				Name:   aws.String("state"),
				Values: []string{state},
			},
		}
	}

	result, err := c.EC2Client.DescribeAvailabilityZones(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe availability zones: %w", err)
	}

	var names, zoneIDs []string
	for _, az := range result.AvailabilityZones {
		if az.ZoneName != nil {
			names = append(names, *az.ZoneName)
		}
		if az.ZoneId != nil {
			zoneIDs = append(zoneIDs, *az.ZoneId)
		}
	}

	return &AvailabilityZonesResult{
		Names:   names,
		ZoneIDs: zoneIDs,
	}, nil
}


// GetCallerIdentity gets the caller's AWS identity
func (c *AWSClient) GetCallerIdentity(ctx context.Context) (*CallerIdentityResult, error) {
	result, err := c.STSClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to get caller identity: %w", err)
	}

	return &CallerIdentityResult{
		AccountID: *result.Account,
		ARN:       *result.Arn,
		UserID:    *result.UserId,
	}, nil
}

// Helper to suppress unused import warning
