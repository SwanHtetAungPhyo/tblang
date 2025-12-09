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

	// Create Security Group using AWS client
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

	// Add ingress rules if specified
	if ingressRules, exists := config["ingress_rules"]; exists {
		if rules, ok := ingressRules.([]interface{}); ok && len(rules) > 0 {
			// Convert rules to AWS format
			// For now, skip rule creation - would need proper conversion
			fmt.Printf("  Note: Ingress rules configuration found but not yet implemented\n")
		}
	}

	// Add egress rules if specified
	if egressRules, exists := config["egress_rules"]; exists {
		if rules, ok := egressRules.([]interface{}); ok && len(rules) > 0 {
			// Convert rules to AWS format
			// For now, skip rule creation - would need proper conversion
			fmt.Printf("  Note: Egress rules configuration found but not yet implemented\n")
		}
	}

	// Return new state
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

	// Delete Security Group using AWS client
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

	// Return nil state to indicate resource is destroyed
	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}

