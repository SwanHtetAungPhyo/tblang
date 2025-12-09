package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) readDataAMI(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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

	// Extract owners
	var owners []string
	if ownersList, exists := config["owners"]; exists {
		if ownerArr, ok := ownersList.([]interface{}); ok {
			for _, o := range ownerArr {
				if oStr, ok := o.(string); ok {
					owners = append(owners, oStr)
				}
			}
		}
	}

	// Extract filters
	var filters []AMIFilter
	if filtersList, exists := config["filters"]; exists {
		if filterArr, ok := filtersList.([]interface{}); ok {
			for _, f := range filterArr {
				if fMap, ok := f.(map[string]interface{}); ok {
					name, _ := fMap["name"].(string)
					var values []string
					if valList, exists := fMap["values"]; exists {
						if valArr, ok := valList.([]interface{}); ok {
							for _, v := range valArr {
								if vStr, ok := v.(string); ok {
									values = append(values, vStr)
								}
							}
						}
					}
					filters = append(filters, AMIFilter{Name: name, Values: values})
				}
			}
		}
	}

	mostRecent, _ := config["most_recent"].(bool)

	ami, err := p.client.DescribeAMI(ctx, owners, filters, mostRecent)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to find AMI",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["ami_id"] = ami.AMIID
	newState["name"] = ami.Name
	newState["architecture"] = ami.Architecture

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

