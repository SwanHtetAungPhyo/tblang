package provider

import (
	"context"

	"github.com/tblang/core/pkg/plugin"
)

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

	if region, exists := config["region"]; exists {
		if regionStr, ok := region.(string); ok {
			p.region = regionStr
		}
	}

	if accountID, exists := config["account_id"]; exists {
		if accountIDStr, ok := accountID.(string); ok {

			_ = accountIDStr
		}
	}

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
