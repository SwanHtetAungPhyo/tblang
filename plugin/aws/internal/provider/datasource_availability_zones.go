package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
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
