package provider

import "github.com/tblang/core/pkg/plugin"

// Data source schema definitions

func getDataAMISchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"owners":       {Type: "list", Description: "List of AMI owners", Required: true},
				"filters":      {Type: "list", Description: "Filters to apply", Optional: true},
				"most_recent":  {Type: "bool", Description: "Return most recent AMI", Optional: true},
				"ami_id":       {Type: "string", Description: "AMI ID", Computed: true},
				"name":         {Type: "string", Description: "AMI name", Computed: true},
				"architecture": {Type: "string", Description: "AMI architecture", Computed: true},
			},
		},
	}
}

func getDataVPCSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"vpc_id":     {Type: "string", Description: "VPC ID to look up", Optional: true},
				"filters":    {Type: "list", Description: "Filters to apply", Optional: true},
				"default":    {Type: "bool", Description: "Return default VPC", Optional: true},
				"cidr_block": {Type: "string", Description: "VPC CIDR block", Computed: true},
				"state":      {Type: "string", Description: "VPC state", Computed: true},
			},
		},
	}
}

func getDataSubnetSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"subnet_id":         {Type: "string", Description: "Subnet ID to look up", Optional: true},
				"vpc_id":            {Type: "string", Description: "VPC ID filter", Optional: true},
				"filters":           {Type: "list", Description: "Filters to apply", Optional: true},
				"cidr_block":        {Type: "string", Description: "Subnet CIDR block", Computed: true},
				"availability_zone": {Type: "string", Description: "Availability zone", Computed: true},
			},
		},
	}
}

func getDataAvailabilityZonesSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"state":    {Type: "string", Description: "Filter by state (available, information, impaired, unavailable)", Optional: true},
				"names":    {Type: "list", Description: "List of availability zone names", Computed: true},
				"zone_ids": {Type: "list", Description: "List of availability zone IDs", Computed: true},
			},
		},
	}
}

func getDataCallerIdentitySchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"account_id": {Type: "string", Description: "AWS account ID", Computed: true},
				"arn":        {Type: "string", Description: "ARN of the caller", Computed: true},
				"user_id":    {Type: "string", Description: "User ID", Computed: true},
			},
		},
	}
}
