package provider

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AWSProvider implements the TBLang provider interface for AWS
type AWSProvider struct {
	ec2Client *ec2.Client
	stsClient *sts.Client
	region    string
	profile   string
	client    *AWSClient
}
