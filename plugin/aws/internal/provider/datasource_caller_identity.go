package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) readDataCallerIdentity(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	identity, err := p.client.GetCallerIdentity(ctx)
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to get caller identity",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	newState["account_id"] = identity.AccountID
	newState["arn"] = identity.ARN
	newState["user_id"] = identity.UserID

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}