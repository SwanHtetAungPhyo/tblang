package provider

import (
	"context"

	"github.com/tblang/core/pkg/plugin"
)

// GetSchema returns the provider and resource schemas
func (p *AWSProvider) GetSchema(ctx context.Context, req *plugin.GetSchemaRequest) (*plugin.GetSchemaResponse, error) {
	return &plugin.GetSchemaResponse{
		Provider:          getProviderSchema(),
		ResourceSchemas:   getResourceSchemas(),
		DataSourceSchemas: getDataSourceSchemas(),
	}, nil
}

func getProviderSchema() *plugin.Schema {
	return &plugin.Schema{
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
	}
}

func getResourceSchemas() map[string]*plugin.Schema {
	return map[string]*plugin.Schema{
		"vpc":              getVPCSchema(),
		"subnet":           getSubnetSchema(),
		"security_group":   getSecurityGroupSchema(),
		"ec2":              getEC2Schema(),
		"internet_gateway": getInternetGatewaySchema(),
		"route_table":      getRouteTableSchema(),
		"eip":              getEIPSchema(),
		"nat_gateway":      getNATGatewaySchema(),
	}
}

func getDataSourceSchemas() map[string]*plugin.Schema {
	return map[string]*plugin.Schema{
		"data_ami":                  getDataAMISchema(),
		"data_vpc":                  getDataVPCSchema(),
		"data_subnet":               getDataSubnetSchema(),
		"data_availability_zones":   getDataAvailabilityZonesSchema(),
		"data_caller_identity":      getDataCallerIdentitySchema(),
	}
}
