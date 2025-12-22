package plugin

import (
	"context"
)

type ProviderPlugin interface {

	GetSchema(ctx context.Context, req *GetSchemaRequest) (*GetSchemaResponse, error)

	Configure(ctx context.Context, req *ConfigureRequest) (*ConfigureResponse, error)

	PlanResourceChange(ctx context.Context, req *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error)

	ApplyResourceChange(ctx context.Context, req *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error)

	ReadResource(ctx context.Context, req *ReadResourceRequest) (*ReadResourceResponse, error)

	ImportResource(ctx context.Context, req *ImportResourceRequest) (*ImportResourceResponse, error)

	ValidateResourceConfig(ctx context.Context, req *ValidateResourceConfigRequest) (*ValidateResourceConfigResponse, error)
}

type Schema struct {
	Version   int64             `json:"version"`
	Block     *SchemaBlock      `json:"block"`
}

type SchemaBlock struct {
	Attributes map[string]*Attribute `json:"attributes"`
	BlockTypes map[string]*BlockType `json:"block_types"`
}

type Attribute struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Optional    bool   `json:"optional"`
	Computed    bool   `json:"computed"`
	Sensitive   bool   `json:"sensitive"`
}

type BlockType struct {
	NestingMode string       `json:"nesting_mode"`
	Block       *SchemaBlock `json:"block"`
	MinItems    int64        `json:"min_items"`
	MaxItems    int64        `json:"max_items"`
}

type GetSchemaRequest struct{}

type GetSchemaResponse struct {
	Provider          *Schema            `json:"provider"`
	ResourceSchemas   map[string]*Schema `json:"resource_schemas"`
	DataSourceSchemas map[string]*Schema `json:"data_source_schemas"`
	Diagnostics       []*Diagnostic      `json:"diagnostics"`
}

type ConfigureRequest struct {
	TerraformVersion string      `json:"terraform_version"`
	Config           interface{} `json:"config"`
}

type ConfigureResponse struct {
	Diagnostics []*Diagnostic `json:"diagnostics"`
}

type PlanResourceChangeRequest struct {
	TypeName         string      `json:"type_name"`
	PriorState       interface{} `json:"prior_state"`
	ProposedNewState interface{} `json:"proposed_new_state"`
	Config           interface{} `json:"config"`
	PriorPrivate     []byte      `json:"prior_private"`
}

type PlanResourceChangeResponse struct {
	PlannedState    interface{}   `json:"planned_state"`
	RequiresReplace []string      `json:"requires_replace"`
	PlannedPrivate  []byte        `json:"planned_private"`
	Diagnostics     []*Diagnostic `json:"diagnostics"`
}

type ApplyResourceChangeRequest struct {
	TypeName       string      `json:"type_name"`
	PriorState     interface{} `json:"prior_state"`
	PlannedState   interface{} `json:"planned_state"`
	Config         interface{} `json:"config"`
	PlannedPrivate []byte      `json:"planned_private"`
}

type ApplyResourceChangeResponse struct {
	NewState    interface{}   `json:"new_state"`
	Private     []byte        `json:"private"`
	Diagnostics []*Diagnostic `json:"diagnostics"`
}

type ReadResourceRequest struct {
	TypeName     string      `json:"type_name"`
	CurrentState interface{} `json:"current_state"`
	Private      []byte      `json:"private"`
}

type ReadResourceResponse struct {
	NewState    interface{}   `json:"new_state"`
	Private     []byte        `json:"private"`
	Diagnostics []*Diagnostic `json:"diagnostics"`
}

type ImportResourceRequest struct {
	TypeName string `json:"type_name"`
	Id       string `json:"id"`
}

type ImportResourceResponse struct {
	ImportedResources []*ImportedResource `json:"imported_resources"`
	Diagnostics       []*Diagnostic       `json:"diagnostics"`
}

type ImportedResource struct {
	TypeName string      `json:"type_name"`
	State    interface{} `json:"state"`
	Private  []byte      `json:"private"`
}

type ValidateResourceConfigRequest struct {
	TypeName string      `json:"type_name"`
	Config   interface{} `json:"config"`
}

type ValidateResourceConfigResponse struct {
	Diagnostics []*Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Severity string `json:"severity"`
	Summary  string `json:"summary"`
	Detail   string `json:"detail"`
}
