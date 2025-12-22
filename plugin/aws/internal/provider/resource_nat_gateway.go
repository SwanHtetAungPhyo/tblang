package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
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
