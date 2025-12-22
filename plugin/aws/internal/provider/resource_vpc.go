package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
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

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["vpc_id"] = vpc.VPCID
	newState["state"] = vpc.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

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

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}
