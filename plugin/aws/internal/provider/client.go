package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AWSClient wraps AWS SDK clients for the plugin
type AWSClient struct {
	EC2    *ec2.Client
	STS    *sts.Client
	Config aws.Config
	Region string
}

// VPCResult represents the result of VPC creation
type VPCResult struct {
	VpcID string
	State string
}

// NewAWSClient creates a new AWS client for the plugin
func NewAWSClient(region string) (*AWSClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &AWSClient{
		EC2:    ec2.NewFromConfig(cfg),
		STS:    sts.NewFromConfig(cfg),
		Config: cfg,
		Region: region,
	}, nil
}

// CreateVPC creates a VPC and returns the result
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

	result, err := c.EC2.CreateVpc(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create VPC: %w", err)
	}

	if result.Vpc == nil {
		return nil, fmt.Errorf("VPC creation returned nil")
	}

	// Enable DNS hostnames and support
	vpcID := *result.Vpc.VpcId
	if err := c.enableVPCDNS(ctx, vpcID); err != nil {
		// Log warning but don't fail
		fmt.Printf("Warning: failed to enable DNS for VPC %s: %v\n", vpcID, err)
	}

	return &VPCResult{
		VpcID: vpcID,
		State: string(result.Vpc.State),
	}, nil
}

// enableVPCDNS enables DNS hostnames and support for a VPC
func (c *AWSClient) enableVPCDNS(ctx context.Context, vpcID string) error {
	// Enable DNS hostnames
	_, err := c.EC2.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:              aws.String(vpcID),
		EnableDnsHostnames: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS hostnames: %w", err)
	}

	// Enable DNS support
	_, err = c.EC2.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId:           aws.String(vpcID),
		EnableDnsSupport: &types.AttributeBooleanValue{Value: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("failed to enable DNS support: %w", err)
	}

	return nil
}

// SubnetResult represents the result of Subnet creation
type SubnetResult struct {
	SubnetID string
	State    string
}

// CreateSubnet creates a subnet and returns the result
func (c *AWSClient) CreateSubnet(ctx context.Context, vpcID, cidrBlock, availabilityZone string, tags map[string]string) (*SubnetResult, error) {
	input := &ec2.CreateSubnetInput{
		VpcId:            aws.String(vpcID),
		CidrBlock:        aws.String(cidrBlock),
		AvailabilityZone: aws.String(availabilityZone),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags:         c.buildTags("subnet", tags),
			},
		},
	}

	result, err := c.EC2.CreateSubnet(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create Subnet: %w", err)
	}

	if result.Subnet == nil {
		return nil, fmt.Errorf("Subnet creation returned nil")
	}

	return &SubnetResult{
		SubnetID: *result.Subnet.SubnetId,
		State:    string(result.Subnet.State),
	}, nil
}

// ConfigureSubnetPublicIP configures public IP mapping for a subnet
func (c *AWSClient) ConfigureSubnetPublicIP(ctx context.Context, subnetID string, mapPublicIP bool) error {
	input := &ec2.ModifySubnetAttributeInput{
		SubnetId: aws.String(subnetID),
		MapPublicIpOnLaunch: &types.AttributeBooleanValue{
			Value: aws.Bool(mapPublicIP),
		},
	}

	_, err := c.EC2.ModifySubnetAttribute(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to configure public IP mapping: %w", err)
	}

	return nil
}

// DeleteVPC deletes a VPC
func (c *AWSClient) DeleteVPC(ctx context.Context, vpcID string) error {
	input := &ec2.DeleteVpcInput{
		VpcId: aws.String(vpcID),
	}

	_, err := c.EC2.DeleteVpc(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete VPC %s: %w", vpcID, err)
	}

	return nil
}

// SecurityGroupResult represents the result of Security Group creation
type SecurityGroupResult struct {
	GroupID string
}

// CreateSecurityGroup creates a security group and returns the result
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

	result, err := c.EC2.CreateSecurityGroup(ctx, input)
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

	_, err := c.EC2.AuthorizeSecurityGroupIngress(ctx, input)
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

	_, err := c.EC2.AuthorizeSecurityGroupEgress(ctx, input)
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

	_, err := c.EC2.DeleteSecurityGroup(ctx, input)
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

	_, err := c.EC2.DeleteSubnet(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete Subnet %s: %w", subnetID, err)
	}

	return nil
}

// buildTags creates AWS tags from a map
func (c *AWSClient) buildTags(resourceName string, additionalTags map[string]string) []types.Tag {
	// Start with additional tags
	tagMap := make(map[string]string)
	for key, value := range additionalTags {
		tagMap[key] = value
	}
	
	// Add default tags only if not already present
	if _, exists := tagMap["ManagedBy"]; !exists {
		tagMap["ManagedBy"] = "TBLang"
	}
	
	// Convert to AWS tags
	var tags []types.Tag
	for key, value := range tagMap {
		tags = append(tags, types.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		})
	}

	return tags
}

// EC2 Instance types and methods

type EC2InstanceConfig struct {
	AMI               string
	InstanceType      string
	SubnetID          string
	SecurityGroups    []string
	KeyName           string
	UserData          string
	AssociatePublicIP bool
	RootVolumeSize    int32
	RootVolumeType    string
	Tags              map[string]string
}

type EC2InstanceResult struct {
	InstanceID string
	PublicIP   string
	PrivateIP  string
	State      string
}

// CreateEC2Instance creates an EC2 instance
func (c *AWSClient) CreateEC2Instance(ctx context.Context, config *EC2InstanceConfig) (*EC2InstanceResult, error) {
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(config.AMI),
		InstanceType: types.InstanceType(config.InstanceType),
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		SubnetId:     aws.String(config.SubnetID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags:         c.buildTags("ec2", config.Tags),
			},
		},
	}

	// Add security groups if specified
	if len(config.SecurityGroups) > 0 {
		input.SecurityGroupIds = config.SecurityGroups
	}

	// Add key name if specified
	if config.KeyName != "" {
		input.KeyName = aws.String(config.KeyName)
	}

	// Add user data if specified
	if config.UserData != "" {
		input.UserData = aws.String(config.UserData)
	}

	// Configure network interface for public IP
	if config.AssociatePublicIP {
		input.NetworkInterfaces = []types.InstanceNetworkInterfaceSpecification{
			{
				DeviceIndex:              aws.Int32(0),
				SubnetId:                 aws.String(config.SubnetID),
				AssociatePublicIpAddress: aws.Bool(true),
				Groups:                   config.SecurityGroups,
			},
		}
		// Clear these as they're now in network interface
		input.SubnetId = nil
		input.SecurityGroupIds = nil
	}

	// Configure root volume
	if config.RootVolumeSize > 0 {
		input.BlockDeviceMappings = []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/xvda"),
				Ebs: &types.EbsBlockDevice{
					VolumeSize:          aws.Int32(config.RootVolumeSize),
					VolumeType:          types.VolumeType(config.RootVolumeType),
					DeleteOnTermination: aws.Bool(true),
				},
			},
		}
	}

	result, err := c.EC2.RunInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create EC2 instance: %w", err)
	}

	if len(result.Instances) == 0 {
		return nil, fmt.Errorf("no instances created")
	}

	instance := result.Instances[0]
	
	// Wait for instance to be running to get IP addresses
	waiter := ec2.NewInstanceRunningWaiter(c.EC2)
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{*instance.InstanceId},
	}, 5*60) // 5 minute timeout
	
	if err != nil {
		fmt.Printf("Warning: instance may not be fully running: %v\n", err)
	}

	// Describe instance to get updated info
	descResult, err := c.EC2.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{*instance.InstanceId},
	})
	
	var publicIP, privateIP, state string
	if err == nil && len(descResult.Reservations) > 0 && len(descResult.Reservations[0].Instances) > 0 {
		inst := descResult.Reservations[0].Instances[0]
		if inst.PublicIpAddress != nil {
			publicIP = *inst.PublicIpAddress
		}
		if inst.PrivateIpAddress != nil {
			privateIP = *inst.PrivateIpAddress
		}
		state = string(inst.State.Name)
	}

	return &EC2InstanceResult{
		InstanceID: *instance.InstanceId,
		PublicIP:   publicIP,
		PrivateIP:  privateIP,
		State:      state,
	}, nil
}

// TerminateEC2Instance terminates an EC2 instance and waits for full termination
func (c *AWSClient) TerminateEC2Instance(ctx context.Context, instanceID string) error {
	_, err := c.EC2.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return fmt.Errorf("failed to terminate instance %s: %w", instanceID, err)
	}

	fmt.Printf("  Waiting for instance %s to terminate...\n", instanceID)

	// Wait for termination with proper timeout (5 minutes)
	waiter := ec2.NewInstanceTerminatedWaiter(c.EC2)
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}, 5*time.Minute)
	
	if err != nil {
		// If waiter fails, poll manually
		fmt.Printf("  Waiter timeout, polling for termination status...\n")
		for i := 0; i < 30; i++ { // Try for up to 60 more seconds
			time.Sleep(2 * time.Second)
			result, descErr := c.EC2.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
				InstanceIds: []string{instanceID},
			})
			if descErr != nil {
				continue
			}
			if len(result.Reservations) > 0 && len(result.Reservations[0].Instances) > 0 {
				state := result.Reservations[0].Instances[0].State.Name
				if state == types.InstanceStateNameTerminated {
					fmt.Printf("  Instance %s terminated successfully\n", instanceID)
					return nil
				}
				fmt.Printf("  Instance state: %s\n", state)
			}
		}
		return fmt.Errorf("timeout waiting for instance %s to terminate", instanceID)
	}

	fmt.Printf("  Instance %s terminated successfully\n", instanceID)
	return nil
}

// Internet Gateway types and methods

type InternetGatewayResult struct {
	GatewayID string
}

// CreateInternetGateway creates an internet gateway and attaches it to a VPC
func (c *AWSClient) CreateInternetGateway(ctx context.Context, vpcID string, tags map[string]string) (*InternetGatewayResult, error) {
	input := &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags:         c.buildTags("igw", tags),
			},
		},
	}

	result, err := c.EC2.CreateInternetGateway(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create internet gateway: %w", err)
	}

	gatewayID := *result.InternetGateway.InternetGatewayId

	// Attach to VPC
	_, err = c.EC2.AttachInternetGateway(ctx, &ec2.AttachInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
		VpcId:             aws.String(vpcID),
	})
	if err != nil {
		// Try to delete the gateway if attach fails
		c.EC2.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
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
		_, err := c.EC2.DetachInternetGateway(ctx, &ec2.DetachInternetGatewayInput{
			InternetGatewayId: aws.String(gatewayID),
			VpcId:             aws.String(vpcID),
		})
		if err != nil {
			return fmt.Errorf("failed to detach internet gateway: %w", err)
		}
	}

	// Delete the gateway
	_, err := c.EC2.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete internet gateway: %w", err)
	}

	return nil
}

// Route Table types and methods

type RouteTableResult struct {
	RouteTableID string
}

// CreateRouteTable creates a route table
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

	result, err := c.EC2.CreateRouteTable(ctx, input)
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

	_, err := c.EC2.CreateRoute(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create route: %w", err)
	}

	return nil
}

// DeleteRouteTable deletes a route table
func (c *AWSClient) DeleteRouteTable(ctx context.Context, routeTableID string) error {
	_, err := c.EC2.DeleteRouteTable(ctx, &ec2.DeleteRouteTableInput{
		RouteTableId: aws.String(routeTableID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete route table: %w", err)
	}

	return nil
}

// EIP types and methods

type EIPResult struct {
	AllocationID string
	PublicIP     string
}

// AllocateEIP allocates an Elastic IP
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

	result, err := c.EC2.AllocateAddress(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to allocate EIP: %w", err)
	}

	return &EIPResult{
		AllocationID: *result.AllocationId,
		PublicIP:     *result.PublicIp,
	}, nil
}

// ReleaseEIP releases an Elastic IP
func (c *AWSClient) ReleaseEIP(ctx context.Context, allocationID string) error {
	_, err := c.EC2.ReleaseAddress(ctx, &ec2.ReleaseAddressInput{
		AllocationId: aws.String(allocationID),
	})
	if err != nil {
		return fmt.Errorf("failed to release EIP: %w", err)
	}

	return nil
}

// NAT Gateway types and methods

type NATGatewayResult struct {
	NATGatewayID string
	State        string
}

// CreateNATGateway creates a NAT gateway
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

	result, err := c.EC2.CreateNatGateway(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create NAT gateway: %w", err)
	}

	natGatewayID := *result.NatGateway.NatGatewayId

	// Wait for NAT gateway to be available
	waiter := ec2.NewNatGatewayAvailableWaiter(c.EC2)
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
	_, err := c.EC2.DeleteNatGateway(ctx, &ec2.DeleteNatGatewayInput{
		NatGatewayId: aws.String(natGatewayID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete NAT gateway: %w", err)
	}

	// Wait for deletion
	waiter := ec2.NewNatGatewayDeletedWaiter(c.EC2)
	err = waiter.Wait(ctx, &ec2.DescribeNatGatewaysInput{
		NatGatewayIds: []string{natGatewayID},
	}, 10*60)

	if err != nil {
		fmt.Printf("Warning: NAT gateway may not be fully deleted: %v\n", err)
	}

	return nil
}

// Data Source types and methods

type AMIFilter struct {
	Name   string
	Values []string
}

type AMIResult struct {
	AMIID        string
	Name         string
	Architecture string
}

// DescribeAMI finds an AMI based on filters
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

	result, err := c.EC2.DescribeImages(ctx, input)
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

type VPCDataResult struct {
	VpcID     string
	CIDRBlock string
	State     string
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

	result, err := c.EC2.DescribeVpcs(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe VPCs: %w", err)
	}

	if len(result.Vpcs) == 0 {
		return nil, fmt.Errorf("no VPCs found")
	}

	vpc := result.Vpcs[0]
	return &VPCDataResult{
		VpcID:     *vpc.VpcId,
		CIDRBlock: *vpc.CidrBlock,
		State:     string(vpc.State),
	}, nil
}

type SubnetDataResult struct {
	SubnetID         string
	CIDRBlock        string
	AvailabilityZone string
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

	result, err := c.EC2.DescribeSubnets(ctx, input)
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

type AvailabilityZonesResult struct {
	Names   []string
	ZoneIDs []string
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

	result, err := c.EC2.DescribeAvailabilityZones(ctx, input)
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

type CallerIdentityResult struct {
	AccountID string
	ARN       string
	UserID    string
}

// GetCallerIdentity gets the caller's AWS identity
func (c *AWSClient) GetCallerIdentity(ctx context.Context) (*CallerIdentityResult, error) {
	result, err := c.STS.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
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
var _ = time.Second