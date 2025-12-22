package provider

import (
"fmt"
"context"

"github.com/tblang/core/pkg/plugin"
)
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

	if mapPublicIP, exists := config["map_public_ip"]; exists {
		if mapPublic, ok := mapPublicIP.(bool); ok && mapPublic {
			if err := p.client.ConfigureSubnetPublicIP(ctx, subnet.SubnetID, true); err != nil {

				fmt.Printf("Warning: failed to configure public IP mapping: %v\n", err)
			}
		}
	}

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

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}
