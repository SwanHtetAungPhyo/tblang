package provider

import (
"fmt"
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) applyRouteTable(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Route Table configuration",
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
					Detail:   "vpc_id is required for Route Table",
				},
			},
		}, nil
	}

	rt, err := p.client.CreateRouteTable(ctx, vpcID, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Route Table",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	if routes, exists := config["routes"]; exists {
		if routeList, ok := routes.([]interface{}); ok {
			for _, route := range routeList {
				if routeMap, ok := route.(map[string]interface{}); ok {
					destCIDR, _ := routeMap["destination_cidr"].(string)
					gatewayID, _ := routeMap["gateway_id"].(string)
					natGatewayID, _ := routeMap["nat_gateway_id"].(string)

					if destCIDR != "" {
						if err := p.client.CreateRoute(ctx, rt.RouteTableID, destCIDR, gatewayID, natGatewayID); err != nil {
							fmt.Printf("Warning: failed to create route: %v\n", err)
						}
					}
				}
			}
		}
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["route_table_id"] = rt.RouteTableID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyRouteTable(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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

	routeTableID, _ := priorState["route_table_id"].(string)
	if routeTableID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Route Table ID",
					Detail:   "route_table_id is required to delete Route Table",
				},
			},
		}, nil
	}

	if err := p.client.DeleteRouteTable(ctx, routeTableID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Route Table",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}
