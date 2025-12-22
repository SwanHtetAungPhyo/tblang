package provider

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type AWSClient struct {
	EC2Client *ec2.Client
	STSClient *sts.Client
	Region    string
}

type VPCResult struct {
	VPCID string
	State string
}

type SubnetResult struct {
	SubnetID string
	State    string
}

type SecurityGroupResult struct {
	GroupID string
}

type EC2InstanceConfig struct {
	AMI                string
	InstanceType       string
	SubnetID           string
	SecurityGroupIDs   []string
	KeyName            string
	UserData           string
	AssociatePublicIP  bool
	RootVolumeSize     int32
	RootVolumeType     string
	Tags               map[string]string
}

type EC2InstanceResult struct {
	InstanceID string
	PublicIP   string
	PrivateIP  string
	State      string
}

type InternetGatewayResult struct {
	GatewayID string
}

type RouteTableResult struct {
	RouteTableID string
}

type EIPResult struct {
	AllocationID string
	PublicIP     string
}

type NATGatewayResult struct {
	NATGatewayID string
	State        string
}

type AMIFilter struct {
	Name   string
	Values []string
}

type AMIResult struct {
	AMIID        string
	Name         string
	Architecture string
}

type VPCDataResult struct {
	VPCID     string
	CIDRBlock string
	State     string
}

type SubnetDataResult struct {
	SubnetID         string
	VPCID            string
	CIDRBlock        string
	AvailabilityZone string
}

type AvailabilityZonesResult struct {
	Names   []string
	ZoneIDs []string
}

type CallerIdentityResult struct {
	AccountID string
	ARN       string
	UserID    string
}
