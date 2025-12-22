package provider

import (
"fmt"
"context"

"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) AllocateEIP(ctx context.Context, tags map[string]string) (*EIPResult, error) {
	input := &ec2.AllocateAddressInput{
		Domain: types.DomainTypeVpc,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeElasticIp,
				Tags:         c.buildTags("eip", tags),
			},
		},
	}

	result, err := c.EC2Client.AllocateAddress(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to allocate EIP: %w", err)
	}

	return &EIPResult{
		AllocationID: *result.AllocationId,
		PublicIP:     *result.PublicIp,
	}, nil
}

func (c *AWSClient) ReleaseEIP(ctx context.Context, allocationID string) error {
	_, err := c.EC2Client.ReleaseAddress(ctx, &ec2.ReleaseAddressInput{
		AllocationId: aws.String(allocationID),
	})
	if err != nil {
		return fmt.Errorf("failed to release EIP: %w", err)
	}

	return nil
}
