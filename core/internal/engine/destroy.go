package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
)

// Destroy removes all infrastructure
func (e *Engine) Destroy(ctx context.Context, filename string) error {
	fmt.Println("Destroying infrastructure...")

	// Compile the tblang file to get cloud vendor configuration (including profile)
	program, err := e.compiler.CompileFile(filename)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	// Load and configure required plugins (sets AWS profile and loads plugin)
	if err := e.loadAndConfigurePlugins(ctx, program); err != nil {
		return fmt.Errorf("failed to load plugins: %w", err)
	}

	// Load current state
	currentState, err := e.stateManager.LoadState()
	if err != nil {
		fmt.Println("No state found, nothing to destroy")
		return nil
	}

	// Show what will be destroyed
	fmt.Println("\nThe following resources will be destroyed:")
	for name, resource := range currentState.Resources {
		fmt.Printf("  - %s (%s)\n", name, resource.Type)
	}

	// Ask for confirmation
	fmt.Print("\nDo you really want to destroy all resources? (yes/no): ")
	var response string
	fmt.Scanln(&response)

	if response != "yes" && response != "y" {
		fmt.Println("Destroy cancelled.")
		return nil
	}

	// Destroy resources using plugins
	if err := e.destroyResources(ctx, currentState); err != nil {
		return fmt.Errorf("failed to destroy resources: %w", err)
	}

	fmt.Println("Destroy complete!")
	return nil
}

// destroyResources destroys all resources in proper order
func (e *Engine) destroyResources(ctx context.Context, currentState *state.State) error {
	// Sort resources by type to ensure proper destruction order
	// Order: ec2 -> nat_gateway -> eip -> route_table -> security_group -> subnet -> internet_gateway -> vpc
	var ec2Instances []*state.ResourceState
	var natGateways []*state.ResourceState
	var eips []*state.ResourceState
	var routeTables []*state.ResourceState
	var securityGroups []*state.ResourceState
	var subnets []*state.ResourceState
	var internetGateways []*state.ResourceState
	var vpcs []*state.ResourceState
	var dataSources []*state.ResourceState
	var others []*state.ResourceState

	for _, resource := range currentState.Resources {
		switch resource.Type {
		case "ec2":
			ec2Instances = append(ec2Instances, resource)
		case "nat_gateway":
			natGateways = append(natGateways, resource)
		case "eip":
			eips = append(eips, resource)
		case "route_table":
			routeTables = append(routeTables, resource)
		case "security_group":
			securityGroups = append(securityGroups, resource)
		case "subnet":
			subnets = append(subnets, resource)
		case "internet_gateway":
			internetGateways = append(internetGateways, resource)
		case "vpc":
			vpcs = append(vpcs, resource)
		case "data_ami", "data_vpc", "data_subnet", "data_availability_zones", "data_caller_identity":
			dataSources = append(dataSources, resource)
		default:
			others = append(others, resource)
		}
	}

	// Destroy in proper order: EC2 first, then dependent resources, then VPCs last
	orderedResources := append([]*state.ResourceState{}, ec2Instances...)
	orderedResources = append(orderedResources, natGateways...)
	orderedResources = append(orderedResources, eips...)
	orderedResources = append(orderedResources, routeTables...)
	orderedResources = append(orderedResources, securityGroups...)
	orderedResources = append(orderedResources, subnets...)
	orderedResources = append(orderedResources, internetGateways...)
	orderedResources = append(orderedResources, others...)
	orderedResources = append(orderedResources, vpcs...)
	// Data sources don't need to be destroyed, but remove from state
	orderedResources = append(orderedResources, dataSources...)

	// Destroy each resource
	for _, resource := range orderedResources {
		warningColor.Printf("Destroying %s (%s)...\n", resource.Name, resource.Type)

		// Use plugin to destroy resource
		if err := e.destroyResourceWithPlugin(ctx, resource); err != nil {
			errorColor.Printf("  Error: failed to destroy %s: %v\n", resource.Name, err)
			// Continue with other resources even if one fails
		} else {
			successColor.Printf("Deleted %s (%s)\n", resource.Name, resource.Type)
		}

		// Remove from state
		delete(currentState.Resources, resource.Name)
		if err := e.stateManager.SaveState(currentState); err != nil {
			return fmt.Errorf("failed to save state: %w", err)
		}
	}

	return nil
}
