package provider

import "github.com/tblang/core/pkg/plugin"

func getVPCSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"cidr_block":           {Type: "string", Description: "CIDR block for VPC", Required: true},
				"enable_dns_hostnames": {Type: "bool", Description: "Enable DNS hostnames", Optional: true},
				"enable_dns_support":   {Type: "bool", Description: "Enable DNS support", Optional: true},
				"tags":                 {Type: "map", Description: "Resource tags", Optional: true},
				"vpc_id":               {Type: "string", Description: "VPC ID", Computed: true},
			},
		},
	}
}

func getSubnetSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"vpc_id":            {Type: "string", Description: "VPC ID", Required: true},
				"cidr_block":        {Type: "string", Description: "CIDR block for subnet", Required: true},
				"availability_zone": {Type: "string", Description: "Availability zone", Required: true},
				"map_public_ip":     {Type: "bool", Description: "Map public IP on launch", Optional: true},
				"tags":              {Type: "map", Description: "Resource tags", Optional: true},
				"subnet_id":         {Type: "string", Description: "Subnet ID", Computed: true},
			},
		},
	}
}

func getSecurityGroupSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"vpc_id":        {Type: "string", Description: "VPC ID", Required: true},
				"name":          {Type: "string", Description: "Security group name", Required: true},
				"description":   {Type: "string", Description: "Security group description", Optional: true},
				"ingress_rules": {Type: "list", Description: "Ingress rules", Optional: true},
				"egress_rules":  {Type: "list", Description: "Egress rules", Optional: true},
				"tags":          {Type: "map", Description: "Resource tags", Optional: true},
				"group_id":      {Type: "string", Description: "Security group ID", Computed: true},
			},
		},
	}
}

func getEC2Schema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"ami":                  {Type: "string", Description: "AMI ID", Required: true},
				"instance_type":        {Type: "string", Description: "Instance type", Required: true},
				"subnet_id":            {Type: "string", Description: "Subnet ID", Required: true},
				"security_groups":      {Type: "list", Description: "Security group IDs", Optional: true},
				"key_name":             {Type: "string", Description: "Key pair name", Optional: true},
				"user_data":            {Type: "string", Description: "User data script", Optional: true},
				"associate_public_ip":  {Type: "bool", Description: "Associate public IP address", Optional: true},
				"root_volume_size":     {Type: "number", Description: "Root volume size in GB", Optional: true},
				"root_volume_type":     {Type: "string", Description: "Root volume type (gp2, gp3, io1, etc.)", Optional: true},
				"tags":                 {Type: "map", Description: "Resource tags", Optional: true},
				"instance_id":          {Type: "string", Description: "Instance ID", Computed: true},
				"public_ip":            {Type: "string", Description: "Public IP address", Computed: true},
				"private_ip":           {Type: "string", Description: "Private IP address", Computed: true},
				"state":                {Type: "string", Description: "Instance state", Computed: true},
			},
		},
	}
}

func getInternetGatewaySchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"vpc_id":     {Type: "string", Description: "VPC ID to attach the gateway to", Required: true},
				"tags":       {Type: "map", Description: "Resource tags", Optional: true},
				"gateway_id": {Type: "string", Description: "Internet Gateway ID", Computed: true},
			},
		},
	}
}

func getRouteTableSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"vpc_id":         {Type: "string", Description: "VPC ID", Required: true},
				"routes":         {Type: "list", Description: "List of routes", Optional: true},
				"tags":           {Type: "map", Description: "Resource tags", Optional: true},
				"route_table_id": {Type: "string", Description: "Route Table ID", Computed: true},
			},
		},
	}
}

func getEIPSchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"domain":        {Type: "string", Description: "Domain (vpc or standard)", Optional: true},
				"instance_id":   {Type: "string", Description: "Instance ID to associate with", Optional: true},
				"tags":          {Type: "map", Description: "Resource tags", Optional: true},
				"allocation_id": {Type: "string", Description: "Allocation ID", Computed: true},
				"public_ip":     {Type: "string", Description: "Public IP address", Computed: true},
			},
		},
	}
}

func getNATGatewaySchema() *plugin.Schema {
	return &plugin.Schema{
		Version: 1,
		Block: &plugin.SchemaBlock{
			Attributes: map[string]*plugin.Attribute{
				"subnet_id":      {Type: "string", Description: "Subnet ID", Required: true},
				"allocation_id":  {Type: "string", Description: "EIP Allocation ID", Required: true},
				"tags":           {Type: "map", Description: "Resource tags", Optional: true},
				"nat_gateway_id": {Type: "string", Description: "NAT Gateway ID", Computed: true},
				"state":          {Type: "string", Description: "NAT Gateway state", Computed: true},
			},
		},
	}
}
