package provider

import (
"context"

"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/ec2"
"github.com/aws/aws-sdk-go-v2/service/sts"
)

func NewAWSClient(region string) (*AWSClient, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	return &AWSClient{
		EC2Client: ec2.NewFromConfig(cfg),
		STSClient: sts.NewFromConfig(cfg),
		Region:    region,
	}, nil
}
