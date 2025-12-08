package provider

import (
	"context"
	"fmt"

	"github.com/tblang/core/pkg/plugin"
)

// AWSProvider implements the TBLang gRPC provider plugin interface for AWS
type AWSProvider struct {
	region    string
	accountID string
	client    *AWSClient
}

// NewAWSProvider creates a new AWS provider
func NewAWSProvider() *AWSProvider {
	return &AWSProvider{}
}

// GetSchema returns the provider and resource schemas
func (p *AWSProvider) GetSchema(ctx context.Context, req *plugin.GetSchemaRequest) (*plugin.GetSchemaResponse, error) {
	return &plugin.GetSchemaResponse{
		Provider: &plugin.Schema{
			Version: 1,
			Block: &plugin.SchemaBlock{
				Attributes: map[string]*plugin.Attribute{
					"region": {
						Type:        "string",
						Description: "AWS region",
						Required:    true,
					},
					"account_id": {
						Type:        "string", 
						Description: "AWS account ID",
						Optional:    true,
					},
				},
			},
		},
		ResourceSchemas: map[string]*plugin.Schema{
			"vpc": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"cidr_block": {
							Type:        "string",
							Description: "CIDR block for VPC",
							Required:    true,
						},
						"enable_dns_hostnames": {
							Type:        "bool",
							Description: "Enable DNS hostnames",
							Optional:    true,
						},
						"enable_dns_support": {
							Type:        "bool", 
							Description: "Enable DNS support",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID",
							Computed:    true,
						},
					},
				},
			},
			"subnet": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID",
							Required:    true,
						},
						"cidr_block": {
							Type:        "string",
							Description: "CIDR block for subnet",
							Required:    true,
						},
						"availability_zone": {
							Type:        "string",
							Description: "Availability zone",
							Required:    true,
						},
						"map_public_ip": {
							Type:        "bool",
							Description: "Map public IP on launch",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"subnet_id": {
							Type:        "string",
							Description: "Subnet ID",
							Computed:    true,
						},
					},
				},
			},
			"security_group": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID",
							Required:    true,
						},
						"name": {
							Type:        "string",
							Description: "Security group name",
							Required:    true,
						},
						"description": {
							Type:        "string",
							Description: "Security group description",
							Optional:    true,
						},
						"ingress_rules": {
							Type:        "list",
							Description: "Ingress rules",
							Optional:    true,
						},
						"egress_rules": {
							Type:        "list",
							Description: "Egress rules",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"group_id": {
							Type:        "string",
							Description: "Security group ID",
							Computed:    true,
						},
					},
				},
			},
			"ec2": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"ami": {
							Type:        "string",
							Description: "AMI ID",
							Required:    true,
						},
						"instance_type": {
							Type:        "string",
							Description: "Instance type",
							Required:    true,
						},
						"subnet_id": {
							Type:        "string",
							Description: "Subnet ID",
							Required:    true,
						},
						"security_groups": {
							Type:        "list",
							Description: "Security group IDs",
							Optional:    true,
						},
						"key_name": {
							Type:        "string",
							Description: "Key pair name",
							Optional:    true,
						},
						"user_data": {
							Type:        "string",
							Description: "User data script",
							Optional:    true,
						},
						"associate_public_ip": {
							Type:        "bool",
							Description: "Associate public IP address",
							Optional:    true,
						},
						"root_volume_size": {
							Type:        "number",
							Description: "Root volume size in GB",
							Optional:    true,
						},
						"root_volume_type": {
							Type:        "string",
							Description: "Root volume type (gp2, gp3, io1, etc.)",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"instance_id": {
							Type:        "string",
							Description: "Instance ID",
							Computed:    true,
						},
						"public_ip": {
							Type:        "string",
							Description: "Public IP address",
							Computed:    true,
						},
						"private_ip": {
							Type:        "string",
							Description: "Private IP address",
							Computed:    true,
						},
						"state": {
							Type:        "string",
							Description: "Instance state",
							Computed:    true,
						},
					},
				},
			},
			"internet_gateway": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID to attach the gateway to",
							Required:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"gateway_id": {
							Type:        "string",
							Description: "Internet Gateway ID",
							Computed:    true,
						},
					},
				},
			},
			"route_table": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID",
							Required:    true,
						},
						"routes": {
							Type:        "list",
							Description: "List of routes",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"route_table_id": {
							Type:        "string",
							Description: "Route Table ID",
							Computed:    true,
						},
					},
				},
			},
			"eip": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"domain": {
							Type:        "string",
							Description: "Domain (vpc or standard)",
							Optional:    true,
						},
						"instance_id": {
							Type:        "string",
							Description: "Instance ID to associate with",
							Optional:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"allocation_id": {
							Type:        "string",
							Description: "Allocation ID",
							Computed:    true,
						},
						"public_ip": {
							Type:        "string",
							Description: "Public IP address",
							Computed:    true,
						},
					},
				},
			},
			"nat_gateway": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"subnet_id": {
							Type:        "string",
							Description: "Subnet ID",
							Required:    true,
						},
						"allocation_id": {
							Type:        "string",
							Description: "EIP Allocation ID",
							Required:    true,
						},
						"tags": {
							Type:        "map",
							Description: "Resource tags",
							Optional:    true,
						},
						"nat_gateway_id": {
							Type:        "string",
							Description: "NAT Gateway ID",
							Computed:    true,
						},
						"state": {
							Type:        "string",
							Description: "NAT Gateway state",
							Computed:    true,
						},
					},
				},
			},
		},
		DataSourceSchemas: map[string]*plugin.Schema{
			"data_ami": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"owners": {
							Type:        "list",
							Description: "List of AMI owners",
							Required:    true,
						},
						"filters": {
							Type:        "list",
							Description: "Filters to apply",
							Optional:    true,
						},
						"most_recent": {
							Type:        "bool",
							Description: "Return most recent AMI",
							Optional:    true,
						},
						"ami_id": {
							Type:        "string",
							Description: "AMI ID",
							Computed:    true,
						},
						"name": {
							Type:        "string",
							Description: "AMI name",
							Computed:    true,
						},
						"architecture": {
							Type:        "string",
							Description: "AMI architecture",
							Computed:    true,
						},
					},
				},
			},
			"data_vpc": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID to look up",
							Optional:    true,
						},
						"filters": {
							Type:        "list",
							Description: "Filters to apply",
							Optional:    true,
						},
						"default": {
							Type:        "bool",
							Description: "Return default VPC",
							Optional:    true,
						},
						"cidr_block": {
							Type:        "string",
							Description: "VPC CIDR block",
							Computed:    true,
						},
						"state": {
							Type:        "string",
							Description: "VPC state",
							Computed:    true,
						},
					},
				},
			},
			"data_subnet": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"subnet_id": {
							Type:        "string",
							Description: "Subnet ID to look up",
							Optional:    true,
						},
						"vpc_id": {
							Type:        "string",
							Description: "VPC ID filter",
							Optional:    true,
						},
						"filters": {
							Type:        "list",
							Description: "Filters to apply",
							Optional:    true,
						},
						"cidr_block": {
							Type:        "string",
							Description: "Subnet CIDR block",
							Computed:    true,
						},
						"availability_zone": {
							Type:        "string",
							Description: "Availability zone",
							Computed:    true,
						},
					},
				},
			},
			"data_availability_zones": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"state": {
							Type:        "string",
							Description: "Filter by state (available, information, impaired, unavailable)",
							Optional:    true,
						},
						"names": {
							Type:        "list",
							Description: "List of availability zone names",
							Computed:    true,
						},
						"zone_ids": {
							Type:        "list",
							Description: "List of availability zone IDs",
							Computed:    true,
						},
					},
				},
			},
			"data_caller_identity": {
				Version: 1,
				Block: &plugin.SchemaBlock{
					Attributes: map[string]*plugin.Attribute{
						"account_id": {
							Type:        "string",
							Description: "AWS account ID",
							Computed:    true,
						},
						"arn": {
							Type:        "string",
							Description: "ARN of the caller",
							Computed:    true,
						},
						"user_id": {
							Type:        "string",
							Description: "User ID",
							Computed:    true,
						},
					},
				},
			},
		},
	}, nil
}

// Configure configures the provider with the given configuration
func (p *AWSProvider) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ConfigureResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	// Extract region
	if region, exists := config["region"]; exists {
		if regionStr, ok := region.(string); ok {
			p.region = regionStr
		}
	}

	// Extract account ID
	if accountID, exists := config["account_id"]; exists {
		if accountIDStr, ok := accountID.(string); ok {
			p.accountID = accountIDStr
		}
	}

	// Initialize AWS client
	client, err := NewAWSClient(p.region)
	if err != nil {
		return &plugin.ConfigureResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to initialize AWS client",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	p.client = client

	return &plugin.ConfigureResponse{}, nil
}

// PlanResourceChange plans changes for a resource
func (p *AWSProvider) PlanResourceChange(ctx context.Context, req *plugin.PlanResourceChangeRequest) (*plugin.PlanResourceChangeResponse, error) {
	// For now, just return the proposed state
	// In a real implementation, this would validate the configuration
	// and determine what changes are needed
	return &plugin.PlanResourceChangeResponse{
		PlannedState: req.ProposedNewState,
	}, nil
}

// ApplyResourceChange applies changes to a resource
func (p *AWSProvider) ApplyResourceChange(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	if p.client == nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Provider not configured",
					Detail:   "AWS provider must be configured before use",
				},
			},
		}, nil
	}

	// Check if this is a destroy operation (planned state is nil/empty but prior state exists)
	isDestroy := req.PlannedState == nil && req.PriorState != nil

	if isDestroy {
		// Handle destroy operations
		switch req.TypeName {
		case "vpc":
			return p.destroyVPC(ctx, req)
		case "subnet":
			return p.destroySubnet(ctx, req)
		case "security_group":
			return p.destroySecurityGroup(ctx, req)
		case "ec2":
			return p.destroyEC2(ctx, req)
		case "internet_gateway":
			return p.destroyInternetGateway(ctx, req)
		case "route_table":
			return p.destroyRouteTable(ctx, req)
		case "eip":
			return p.destroyEIP(ctx, req)
		case "nat_gateway":
			return p.destroyNATGateway(ctx, req)
		// Data sources don't need destroy
		case "data_ami", "data_vpc", "data_subnet", "data_availability_zones", "data_caller_identity":
			return &plugin.ApplyResourceChangeResponse{NewState: nil}, nil
		default:
			return &plugin.ApplyResourceChangeResponse{
				Diagnostics: []*plugin.Diagnostic{
					{
						Severity: "error",
						Summary:  "Unsupported resource type",
						Detail:   fmt.Sprintf("Resource type %s is not supported", req.TypeName),
					},
				},
			}, nil
		}
	}

	// Handle create/update operations
	switch req.TypeName {
	case "vpc":
		return p.applyVPC(ctx, req)
	case "subnet":
		return p.applySubnet(ctx, req)
	case "security_group":
		return p.applySecurityGroup(ctx, req)
	case "ec2":
		return p.applyEC2(ctx, req)
	case "internet_gateway":
		return p.applyInternetGateway(ctx, req)
	case "route_table":
		return p.applyRouteTable(ctx, req)
	case "eip":
		return p.applyEIP(ctx, req)
	case "nat_gateway":
		return p.applyNATGateway(ctx, req)
	// Data sources
	case "data_ami":
		return p.readDataAMI(ctx, req)
	case "data_vpc":
		return p.readDataVPC(ctx, req)
	case "data_subnet":
		return p.readDataSubnet(ctx, req)
	case "data_availability_zones":
		return p.readDataAvailabilityZones(ctx, req)
	case "data_caller_identity":
		return p.readDataCallerIdentity(ctx, req)
	default:
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Unsupported resource type",
					Detail:   fmt.Sprintf("Resource type %s is not supported", req.TypeName),
				},
			},
		}, nil
	}
}

// ReadResource reads the current state of a resource
func (p *AWSProvider) ReadResource(ctx context.Context, req *plugin.ReadResourceRequest) (*plugin.ReadResourceResponse, error) {
	// TODO: Implement resource reading from AWS
	return &plugin.ReadResourceResponse{
		NewState: req.CurrentState,
	}, nil
}

// ImportResource imports an existing resource
func (p *AWSProvider) ImportResource(ctx context.Context, req *plugin.ImportResourceRequest) (*plugin.ImportResourceResponse, error) {
	// TODO: Implement resource import
	return &plugin.ImportResourceResponse{}, nil
}

// ValidateResourceConfig validates a resource configuration
func (p *AWSProvider) ValidateResourceConfig(ctx context.Context, req *plugin.ValidateResourceConfigRequest) (*plugin.ValidateResourceConfigResponse, error) {
	// TODO: Implement configuration validation
	return &plugin.ValidateResourceConfigResponse{}, nil
}

// Resource-specific apply methods

func (p *AWSProvider) applyVPC(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid VPC configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	cidrBlock, _ := config["cidr_block"].(string)
	if cidrBlock == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required field",
					Detail:   "cidr_block is required for VPC",
				},
			},
		}, nil
	}

	// Create VPC using AWS client
	vpc, err := p.client.CreateVPC(ctx, cidrBlock, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create VPC",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Return new state
	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["vpc_id"] = vpc.VpcID
	newState["state"] = vpc.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) applySubnet(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Subnet configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	cidrBlock, _ := config["cidr_block"].(string)
	availabilityZone, _ := config["availability_zone"].(string)
	
	if vpcID == "" || cidrBlock == "" || availabilityZone == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "vpc_id, cidr_block, and availability_zone are required for Subnet",
				},
			},
		}, nil
	}

	// Create Subnet using AWS client
	subnet, err := p.client.CreateSubnet(ctx, vpcID, cidrBlock, availabilityZone, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Subnet",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Configure public IP mapping if specified
	if mapPublicIP, exists := config["map_public_ip"]; exists {
		if mapPublic, ok := mapPublicIP.(bool); ok && mapPublic {
			if err := p.client.ConfigureSubnetPublicIP(ctx, subnet.SubnetID, true); err != nil {
				// Log warning but don't fail
				fmt.Printf("Warning: failed to configure public IP mapping: %v\n", err)
			}
		}
	}

	// Return new state
	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["subnet_id"] = subnet.SubnetID
	newState["state"] = subnet.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) applySecurityGroup(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Security Group configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	name, _ := config["name"].(string)
	description, _ := config["description"].(string)
	
	if vpcID == "" || name == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "vpc_id and name are required for Security Group",
				},
			},
		}, nil
	}

	if description == "" {
		description = "Managed by TBLang"
	}

	// Create Security Group using AWS client
	sg, err := p.client.CreateSecurityGroup(ctx, vpcID, name, description, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Security Group",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Add ingress rules if specified
	if ingressRules, exists := config["ingress_rules"]; exists {
		if rules, ok := ingressRules.([]interface{}); ok && len(rules) > 0 {
			// Convert rules to AWS format
			// For now, skip rule creation - would need proper conversion
			fmt.Printf("  Note: Ingress rules configuration found but not yet implemented\n")
		}
	}

	// Add egress rules if specified
	if egressRules, exists := config["egress_rules"]; exists {
		if rules, ok := egressRules.([]interface{}); ok && len(rules) > 0 {
			// Convert rules to AWS format
			// For now, skip rule creation - would need proper conversion
			fmt.Printf("  Note: Egress rules configuration found but not yet implemented\n")
		}
	}

	// Return new state
	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["group_id"] = sg.GroupID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) applyEC2(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid EC2 configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	ami, _ := config["ami"].(string)
	instanceType, _ := config["instance_type"].(string)
	subnetID, _ := config["subnet_id"].(string)

	if ami == "" || instanceType == "" || subnetID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "ami, instance_type, and subnet_id are required for EC2",
				},
			},
		}, nil
	}

	// Extract optional parameters
	keyName, _ := config["key_name"].(string)
	userData, _ := config["user_data"].(string)
	associatePublicIP, _ := config["associate_public_ip"].(bool)
	
	// Extract security groups
	var securityGroups []string
	if sgs, exists := config["security_groups"]; exists {
		if sgList, ok := sgs.([]interface{}); ok {
			for _, sg := range sgList {
				if sgStr, ok := sg.(string); ok {
					securityGroups = append(securityGroups, sgStr)
				}
			}
		}
	}

	// Extract root volume configuration
	var rootVolumeSize int32 = 8 // default
	var rootVolumeType string = "gp3" // default
	if size, exists := config["root_volume_size"]; exists {
		if sizeFloat, ok := size.(float64); ok {
			rootVolumeSize = int32(sizeFloat)
		}
	}
	if volType, exists := config["root_volume_type"]; exists {
		if volTypeStr, ok := volType.(string); ok {
			rootVolumeType = volTypeStr
		}
	}

	// Create EC2 instance using AWS client
	instance, err := p.client.CreateEC2Instance(ctx, &EC2InstanceConfig{
		AMI:               ami,
		InstanceType:      instanceType,
		SubnetID:          subnetID,
		SecurityGroups:    securityGroups,
		KeyName:           keyName,
		UserData:          userData,
		AssociatePublicIP: associatePublicIP,
		RootVolumeSize:    rootVolumeSize,
		RootVolumeType:    rootVolumeType,
		Tags:              extractTags(config),
	})
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create EC2 instance",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Return new state
	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["instance_id"] = instance.InstanceID
	newState["public_ip"] = instance.PublicIP
	newState["private_ip"] = instance.PrivateIP
	newState["state"] = instance.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

// Resource-specific destroy methods

func (p *AWSProvider) destroyVPC(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	vpcID, ok := priorState["vpc_id"].(string)
	if !ok || vpcID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing VPC ID",
					Detail:   "vpc_id is required to destroy VPC",
				},
			},
		}, nil
	}

	// Delete VPC using AWS client
	if err := p.client.DeleteVPC(ctx, vpcID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete VPC",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Return nil state to indicate resource is destroyed
	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

func (p *AWSProvider) destroySubnet(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	subnetID, ok := priorState["subnet_id"].(string)
	if !ok || subnetID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Subnet ID",
					Detail:   "subnet_id is required to destroy Subnet",
				},
			},
		}, nil
	}

	// Delete Subnet using AWS client
	if err := p.client.DeleteSubnet(ctx, subnetID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Subnet",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Return nil state to indicate resource is destroyed
	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

func (p *AWSProvider) destroySecurityGroup(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	groupID, ok := priorState["group_id"].(string)
	if !ok || groupID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Security Group ID",
					Detail:   "group_id is required to destroy Security Group",
				},
			},
		}, nil
	}

	// Delete Security Group using AWS client
	if err := p.client.DeleteSecurityGroup(ctx, groupID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Security Group",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Return nil state to indicate resource is destroyed
	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

func (p *AWSProvider) destroyEC2(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	instanceID, ok := priorState["instance_id"].(string)
	if !ok || instanceID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Instance ID",
					Detail:   "instance_id is required to terminate EC2 instance",
				},
			},
		}, nil
	}

	// Terminate EC2 instance using AWS client
	if err := p.client.TerminateEC2Instance(ctx, instanceID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to terminate EC2 instance",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

// Helper functions

func extractTags(config map[string]interface{}) map[string]string {
	tags := make(map[string]string)
	
	if tagsInterface, exists := config["tags"]; exists {
		if tagsMap, ok := tagsInterface.(map[string]interface{}); ok {
			for k, v := range tagsMap {
				if str, ok := v.(string); ok {
					tags[k] = str
				}
			}
		}
	}
	
	return tags
}

// Internet Gateway methods

func (p *AWSProvider) applyInternetGateway(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Internet Gateway configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	if vpcID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required field",
					Detail:   "vpc_id is required for Internet Gateway",
				},
			},
		}, nil
	}

	igw, err := p.client.CreateInternetGateway(ctx, vpcID, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Internet Gateway",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["gateway_id"] = igw.GatewayID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyInternetGateway(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	gatewayID, _ := priorState["gateway_id"].(string)
	vpcID, _ := priorState["vpc_id"].(string)

	if gatewayID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Gateway ID",
					Detail:   "gateway_id is required to delete Internet Gateway",
				},
			},
		}, nil
	}

	if err := p.client.DeleteInternetGateway(ctx, gatewayID, vpcID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Internet Gateway",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

// Route Table methods

func (p *AWSProvider) applyRouteTable(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Route Table configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	if vpcID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required field",
					Detail:   "vpc_id is required for Route Table",
				},
			},
		}, nil
	}

	rt, err := p.client.CreateRouteTable(ctx, vpcID, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Route Table",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	// Add routes if specified
	if routes, exists := config["routes"]; exists {
		if routeList, ok := routes.([]interface{}); ok {
			for _, route := range routeList {
				if routeMap, ok := route.(map[string]interface{}); ok {
					destCIDR, _ := routeMap["destination_cidr"].(string)
					gatewayID, _ := routeMap["gateway_id"].(string)
					natGatewayID, _ := routeMap["nat_gateway_id"].(string)
					
					if destCIDR != "" {
						if err := p.client.CreateRoute(ctx, rt.RouteTableID, destCIDR, gatewayID, natGatewayID); err != nil {
							fmt.Printf("Warning: failed to create route: %v\n", err)
						}
					}
				}
			}
		}
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["route_table_id"] = rt.RouteTableID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyRouteTable(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	routeTableID, _ := priorState["route_table_id"].(string)
	if routeTableID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Route Table ID",
					Detail:   "route_table_id is required to delete Route Table",
				},
			},
		}, nil
	}

	if err := p.client.DeleteRouteTable(ctx, routeTableID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Route Table",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

// EIP methods

func (p *AWSProvider) applyEIP(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid EIP configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	eip, err := p.client.AllocateEIP(ctx, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to allocate EIP",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["allocation_id"] = eip.AllocationID
	newState["public_ip"] = eip.PublicIP

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyEIP(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	allocationID, _ := priorState["allocation_id"].(string)
	if allocationID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Allocation ID",
					Detail:   "allocation_id is required to release EIP",
				},
			},
		}, nil
	}

	if err := p.client.ReleaseEIP(ctx, allocationID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to release EIP",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

// NAT Gateway methods

func (p *AWSProvider) applyNATGateway(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid NAT Gateway configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	subnetID, _ := config["subnet_id"].(string)
	allocationID, _ := config["allocation_id"].(string)

	if subnetID == "" || allocationID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "subnet_id and allocation_id are required for NAT Gateway",
				},
			},
		}, nil
	}

	natGW, err := p.client.CreateNATGateway(ctx, subnetID, allocationID, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create NAT Gateway",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["nat_gateway_id"] = natGW.NATGatewayID
	newState["state"] = natGW.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyNATGateway(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	priorState, ok := req.PriorState.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid prior state",
					Detail:   "Prior state must be a map",
				},
			},
		}, nil
	}

	natGatewayID, _ := priorState["nat_gateway_id"].(string)
	if natGatewayID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing NAT Gateway ID",
					Detail:   "nat_gateway_id is required to delete NAT Gateway",
				},
			},
		}, nil
	}

	if err := p.client.DeleteNATGateway(ctx, natGatewayID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete NAT Gateway",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

// Data Source methods

func (p *AWSProvider) readDataAMI(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid data source configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	// Extract owners
	var owners []string
	if ownersList, exists := config["owners"]; exists {
		if ownerArr, ok := ownersList.([]interface{}); ok {
			for _, o := range ownerArr {
				if oStr, ok := o.(string); ok {
					owners = append(owners, oStr)
				}
			}
		}
	}

	// Extract filters
	var filters []AMIFilter
	if filtersList, exists := config["filters"]; exists {
		if filterArr, ok := filtersList.([]interface{}); ok {
			for _, f := range filterArr {
				if fMap, ok := f.(map[string]interface{}); ok {
					name, _ := fMap["name"].(string)
					var values []string
					if valList, exists := fMap["values"]; exists {
						if valArr, ok := valList.([]interface{}); ok {
							for _, v := range valArr {
								if vStr, ok := v.(string); ok {
									values = append(values, vStr)
								}
							}
						}
					}
					filters = append(filters, AMIFilter{Name: name, Values: values})
				}
			}
		}
	}

	mostRecent, _ := config["most_recent"].(bool)

	ami, err := p.client.DescribeAMI(ctx, owners, filters, mostRecent)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to find AMI",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["ami_id"] = ami.AMIID
	newState["name"] = ami.Name
	newState["architecture"] = ami.Architecture

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) readDataVPC(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid data source configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	isDefault, _ := config["default"].(bool)

	vpc, err := p.client.DescribeVPC(ctx, vpcID, isDefault)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to find VPC",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["vpc_id"] = vpc.VpcID
	newState["cidr_block"] = vpc.CIDRBlock
	newState["state"] = vpc.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) readDataSubnet(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid data source configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	subnetID, _ := config["subnet_id"].(string)
	vpcID, _ := config["vpc_id"].(string)

	subnet, err := p.client.DescribeSubnet(ctx, subnetID, vpcID)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to find Subnet",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["subnet_id"] = subnet.SubnetID
	newState["cidr_block"] = subnet.CIDRBlock
	newState["availability_zone"] = subnet.AvailabilityZone

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) readDataAvailabilityZones(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		config = make(map[string]interface{})
	}

	state, _ := config["state"].(string)
	if state == "" {
		state = "available"
	}

	azs, err := p.client.DescribeAvailabilityZones(ctx, state)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to describe availability zones",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["names"] = azs.Names
	newState["zone_ids"] = azs.ZoneIDs

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) readDataCallerIdentity(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	identity, err := p.client.GetCallerIdentity(ctx)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to get caller identity",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	newState["account_id"] = identity.AccountID
	newState["arn"] = identity.ARN
	newState["user_id"] = identity.UserID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}