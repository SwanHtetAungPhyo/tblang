package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tblang/core/pkg/plugin/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	proto.UnimplementedProviderServer
	provider GRPCProviderPlugin
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(provider GRPCProviderPlugin) *GRPCServer {
	return &GRPCServer{
		provider: provider,
	}
}

func (s *GRPCServer) Serve() error {

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}
	s.listener = listener

	s.server = grpc.NewServer()
	proto.RegisterProviderServer(s.server, s)

	connectionInfo := map[string]interface{}{
		"network": "tcp",
		"address": listener.Addr().String(),
		"protocol": "grpc",
	}

	if err := json.NewEncoder(os.Stdout).Encode(connectionInfo); err != nil {
		return fmt.Errorf("failed to output connection info: %w", err)
	}
	os.Stdout.Sync()

	log.Printf("gRPC plugin server listening on %s", listener.Addr().String())

	return s.server.Serve(listener)
}

func (s *GRPCServer) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}

func (s *GRPCServer) GetSchema(ctx context.Context, req *proto.GetSchemaRequest) (*proto.GetSchemaResponse, error) {

	interfaceReq := &GetSchemaRequest{}

	resp, err := s.provider.GetSchema(ctx, interfaceReq)
	if err != nil {
		return nil, err
	}

	return GetSchemaResponseToProto(resp), nil
}

func (s *GRPCServer) Configure(ctx context.Context, req *proto.ConfigureRequest) (*proto.ConfigureResponse, error) {

	interfaceReq := &ConfigureRequest{
		TerraformVersion: req.TerraformVersion,
	}

	if req.Config != nil && len(req.Config.Json) > 0 {
		var config interface{}
		if err := json.Unmarshal(req.Config.Json, &config); err == nil {
			interfaceReq.Config = config
		}
	}

	resp, err := s.provider.Configure(ctx, interfaceReq)
	if err != nil {
		return nil, err
	}

	return ConfigureResponseToProto(resp), nil
}

func (s *GRPCServer) ApplyResourceChange(ctx context.Context, req *proto.ApplyResourceChangeRequest) (*proto.ApplyResourceChangeResponse, error) {

	interfaceReq := &ApplyResourceChangeRequest{
		TypeName:       req.TypeName,
		PlannedPrivate: req.PlannedPrivate,
	}

	if req.PriorState != nil && len(req.PriorState.Json) > 0 {
		var state interface{}
		if err := json.Unmarshal(req.PriorState.Json, &state); err == nil {
			interfaceReq.PriorState = state
		}
	}

	if req.PlannedState != nil && len(req.PlannedState.Json) > 0 {
		var state interface{}
		if err := json.Unmarshal(req.PlannedState.Json, &state); err == nil {
			interfaceReq.PlannedState = state
		}
	}

	if req.Config != nil && len(req.Config.Json) > 0 {
		var config interface{}
		if err := json.Unmarshal(req.Config.Json, &config); err == nil {
			interfaceReq.Config = config
		}
	}

	resp, err := s.provider.ApplyResourceChange(ctx, interfaceReq)
	if err != nil {
		return nil, err
	}

	return ApplyResourceChangeResponseToProto(resp), nil
}

func (s *GRPCServer) PlanResourceChange(ctx context.Context, req *proto.PlanResourceChangeRequest) (*proto.PlanResourceChangeResponse, error) {
	return &proto.PlanResourceChangeResponse{}, nil
}

func (s *GRPCServer) ReadResource(ctx context.Context, req *proto.ReadResourceRequest) (*proto.ReadResourceResponse, error) {
	return &proto.ReadResourceResponse{}, nil
}

func (s *GRPCServer) ImportResource(ctx context.Context, req *proto.ImportResourceRequest) (*proto.ImportResourceResponse, error) {
	return &proto.ImportResourceResponse{}, nil
}

func (s *GRPCServer) ValidateResourceConfig(ctx context.Context, req *proto.ValidateResourceConfigRequest) (*proto.ValidateResourceConfigResponse, error) {
	return &proto.ValidateResourceConfigResponse{}, nil
}

func GetSchemaResponseToProto(resp *GetSchemaResponse) *proto.GetSchemaResponse {
	protoResp := &proto.GetSchemaResponse{
		ResourceSchemas:   make(map[string]*proto.Schema),
		DataSourceSchemas: make(map[string]*proto.Schema),
		Diagnostics:       make([]*proto.Diagnostic, len(resp.Diagnostics)),
	}

	if resp.Provider != nil {
		protoResp.Provider = SchemaToProto(resp.Provider)
	}

	for name, schema := range resp.ResourceSchemas {
		protoResp.ResourceSchemas[name] = SchemaToProto(schema)
	}

	for name, schema := range resp.DataSourceSchemas {
		protoResp.DataSourceSchemas[name] = SchemaToProto(schema)
	}

	for i, diag := range resp.Diagnostics {
		protoResp.Diagnostics[i] = DiagnosticToProto(diag)
	}

	return protoResp
}

func SchemaToProto(schema *Schema) *proto.Schema {
	protoSchema := &proto.Schema{
		Version: schema.Version,
	}

	if schema.Block != nil {
		protoSchema.Block = SchemaBlockToProto(schema.Block)
	}

	return protoSchema
}

func SchemaBlockToProto(block *SchemaBlock) *proto.SchemaBlock {
	protoBlock := &proto.SchemaBlock{
		Attributes: make(map[string]*proto.Attribute),
		BlockTypes: make(map[string]*proto.BlockType),
	}

	for name, attr := range block.Attributes {
		protoBlock.Attributes[name] = AttributeToProto(attr)
	}

	for name, blockType := range block.BlockTypes {
		protoBlock.BlockTypes[name] = BlockTypeToProto(blockType)
	}

	return protoBlock
}

func AttributeToProto(attr *Attribute) *proto.Attribute {
	return &proto.Attribute{
		Type:        attr.Type,
		Description: attr.Description,
		Required:    attr.Required,
		Optional:    attr.Optional,
		Computed:    attr.Computed,
		Sensitive:   attr.Sensitive,
	}
}

func BlockTypeToProto(blockType *BlockType) *proto.BlockType {
	protoBlockType := &proto.BlockType{
		NestingMode: blockType.NestingMode,
		MinItems:    blockType.MinItems,
		MaxItems:    blockType.MaxItems,
	}

	if blockType.Block != nil {
		protoBlockType.Block = SchemaBlockToProto(blockType.Block)
	}

	return protoBlockType
}

func DiagnosticToProto(diag *Diagnostic) *proto.Diagnostic {
	return &proto.Diagnostic{
		Severity: diag.Severity,
		Summary:  diag.Summary,
		Detail:   diag.Detail,
	}
}

func ConfigureResponseToProto(resp *ConfigureResponse) *proto.ConfigureResponse {
	protoResp := &proto.ConfigureResponse{
		Diagnostics: make([]*proto.Diagnostic, len(resp.Diagnostics)),
	}

	for i, diag := range resp.Diagnostics {
		protoResp.Diagnostics[i] = DiagnosticToProto(diag)
	}

	return protoResp
}

func ApplyResourceChangeResponseToProto(resp *ApplyResourceChangeResponse) *proto.ApplyResourceChangeResponse {
	protoResp := &proto.ApplyResourceChangeResponse{
		Private:     resp.Private,
		Diagnostics: make([]*proto.Diagnostic, len(resp.Diagnostics)),
	}

	if resp.NewState != nil {
		if jsonData, err := json.Marshal(resp.NewState); err == nil {
			protoResp.NewState = &proto.DynamicValue{Json: jsonData}
		}
	}

	for i, diag := range resp.Diagnostics {
		protoResp.Diagnostics[i] = DiagnosticToProto(diag)
	}

	return protoResp
}
