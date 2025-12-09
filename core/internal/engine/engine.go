package engine

// This file serves as the main entry point for the engine package.
// All functionality has been split into separate files for better organization:
//
// - types.go: Type definitions and color constants
// - constructor.go: Engine constructor
// - initialize.go: Initialization and shutdown
// - plan.go: Plan command and display
// - apply.go: Apply command and changes
// - destroy.go: Destroy command and resource ordering
// - show.go: Show command
// - graph.go: Graph visualization
// - plugin_loader.go: Plugin loading and configuration
// - changes.go: Change calculation
// - resource_plugin.go: Plugin-based resource operations
// - resource_resolver.go: Resource reference resolution
// - aws_cli.go: AWS CLI integration (legacy)
