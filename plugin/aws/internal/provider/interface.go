package provider

import (
	"context"

	"github.com/tblang/core/pkg/plugin"
)

func (p *AWSProvider) PlanResourceChange(ctx context.Context, req *plugin.PlanResourceChangeRequest) (*plugin.PlanResourceChangeResponse, error) {

	return &plugin.PlanResourceChangeResponse{
		PlannedState: req.ProposedNewState,
	}, nil
}

func (p *AWSProvider) ReadResource(ctx context.Context, req *plugin.ReadResourceRequest) (*plugin.ReadResourceResponse, error) {

	return &plugin.ReadResourceResponse{
		NewState: req.CurrentState,
	}, nil
}

func (p *AWSProvider) ImportResource(ctx context.Context, req *plugin.ImportResourceRequest) (*plugin.ImportResourceResponse, error) {

	return &plugin.ImportResourceResponse{}, nil
}

func (p *AWSProvider) ValidateResourceConfig(ctx context.Context, req *plugin.ValidateResourceConfigRequest) (*plugin.ValidateResourceConfigResponse, error) {

	return &plugin.ValidateResourceConfigResponse{}, nil
}
