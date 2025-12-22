package provider

import (
	"context"
	"fmt"

	"github.com/tblang/core/pkg/plugin"
)

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

	isDestroy := req.PlannedState == nil && req.PriorState != nil

	if isDestroy {
		return p.handleDestroy(ctx, req)
	}

	return p.handleCreateOrUpdate(ctx, req)
}

func (p *AWSProvider) handleDestroy(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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

func (p *AWSProvider) handleCreateOrUpdate(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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
