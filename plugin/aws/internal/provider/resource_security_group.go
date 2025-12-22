package provider

import (
"fmt"
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) applySecurityGroup(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid Security Group configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	vpcID, _ := config["vpc_id"].(string)
	name, _ := config["name"].(string)
	description, _ := config["description"].(string)

	if vpcID == "" || name == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "vpc_id and name are required for Security Group",
				},
			},
		}, nil
	}

	if description == "" {
		description = "Managed by TBLang"
	}

	sg, err := p.client.CreateSecurityGroup(ctx, vpcID, name, description, extractTags(config))
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create Security Group",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	if ingressRules, exists := config["ingress_rules"]; exists {
		if rules, ok := ingressRules.([]interface{}); ok && len(rules) > 0 {

			fmt.Printf("  Note: Ingress rules configuration found but not yet implemented\n")
		}
	}

	if egressRules, exists := config["egress_rules"]; exists {
		if rules, ok := egressRules.([]interface{}); ok && len(rules) > 0 {

			fmt.Printf("  Note: Egress rules configuration found but not yet implemented\n")
		}
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["group_id"] = sg.GroupID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroySecurityGroup(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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

	groupID, ok := priorState["group_id"].(string)
	if !ok || groupID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Security Group ID",
					Detail:   "group_id is required to destroy Security Group",
				},
			},
		}, nil
	}

	if err := p.client.DeleteSecurityGroup(ctx, groupID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to delete Security Group",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}
