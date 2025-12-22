package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
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
