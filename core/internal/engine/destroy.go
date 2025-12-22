package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
)

func (e *Engine) Destroy(ctx context.Context, filename string) error {
	fmt.Println("Destroying infrastructure...")

	program, err := e.compiler.CompileFile(filename)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	if err := e.loadAndConfigurePlugins(ctx, program); err != nil {
		return fmt.Errorf("failed to load plugins: %w", err)
	}

	currentState, err := e.stateManager.LoadState()
	if err != nil {
		fmt.Println("No state found, nothing to destroy")
		return nil
	}

	fmt.Println("\nThe following resources will be destroyed:")
	for name, resource := range currentState.Resources {
		fmt.Printf("  - %s (%s)\n", name, resource.Type)
	}

	fmt.Print("\nDo you really want to destroy all resources? (yes/no): ")
	var response string
	fmt.Scanln(&response)

	if response != "yes" && response != "y" {
		fmt.Println("Destroy cancelled.")
		return nil
	}

	if err := e.destroyResources(ctx, currentState); err != nil {
		return fmt.Errorf("failed to destroy resources: %w", err)
	}

	fmt.Println("Destroy complete!")
	return nil
}

func (e *Engine) destroyResources(ctx context.Context, currentState *state.State) error {

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

	orderedResources := append([]*state.ResourceState{}, ec2Instances...)
	orderedResources = append(orderedResources, natGateways...)
	orderedResources = append(orderedResources, eips...)
	orderedResources = append(orderedResources, routeTables...)
	orderedResources = append(orderedResources, securityGroups...)
	orderedResources = append(orderedResources, subnets...)
	orderedResources = append(orderedResources, internetGateways...)
	orderedResources = append(orderedResources, others...)
	orderedResources = append(orderedResources, vpcs...)

	orderedResources = append(orderedResources, dataSources...)

	for _, resource := range orderedResources {
		warningColor.Printf("Destroying %s (%s)...\n", resource.Name, resource.Type)

		if err := e.destroyResourceWithPlugin(ctx, resource); err != nil {
			errorColor.Printf("  Error: failed to destroy %s: %v\n", resource.Name, err)

		} else {
			successColor.Printf("Deleted %s (%s)\n", resource.Name, resource.Type)
		}

		delete(currentState.Resources, resource.Name)
		if err := e.stateManager.SaveState(currentState); err != nil {
			return fmt.Errorf("failed to save state: %w", err)
		}
	}

	return nil
}
