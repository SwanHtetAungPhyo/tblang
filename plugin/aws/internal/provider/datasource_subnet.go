package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) readDataSubnet(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid data source configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	subnetID, _ := config["subnet_id"].(string)
	vpcID, _ := config["vpc_id"].(string)

	subnet, err := p.client.DescribeSubnet(ctx, subnetID, vpcID)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to find Subnet",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["subnet_id"] = subnet.SubnetID
	newState["cidr_block"] = subnet.CIDRBlock
	newState["availability_zone"] = subnet.AvailabilityZone

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

