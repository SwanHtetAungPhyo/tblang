package engine

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/tblang/core/internal/state"
)

// AWS CLI integration methods for testing

func (e *Engine) createVPCWithAWSCLI(resource *state.ResourceState) error {
	cidrBlock, ok := resource.Attributes["cidr_block"].(string)
	if !ok {
		return fmt.Errorf("cidr_block not found in VPC configuration")
	}

	infoColor.Printf("  Creating VPC with CIDR: %s\n", cidrBlock)

	// Create VPC using AWS CLI
	cmd := exec.Command("aws", "ec2", "create-vpc", "--cidr-block", cidrBlock, "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("aws cli error: %w", err)
	}

	// Parse the output to get VPC ID
	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return fmt.Errorf("failed to parse AWS CLI output: %w", err)
	}

	vpc, ok := result["Vpc"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid AWS CLI response format")
	}

	vpcID, ok := vpc["VpcId"].(string)
	if !ok {
		return fmt.Errorf("VPC ID not found in response")
	}

	// Update resource attributes with actual VPC ID
	resource.Attributes["vpc_id"] = vpcID
	resource.Attributes["state"] = vpc["State"]

	successColor.Printf("  VPC created with ID: %s\n", vpcID)

	// Add tags if specified
	if tags, exists := resource.Attributes["tags"]; exists {
		if err := e.tagVPCWithAWSCLI(vpcID, tags); err != nil {
			fmt.Printf("  Warning: failed to tag VPC: %v\n", err)
		}
	}

	return nil
}

func (e *Engine) tagVPCWithAWSCLI(vpcID string, tags interface{}) error {
	tagsMap, ok := tags.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid tags format")
	}

	for key, value := range tagsMap {
		tagSpec := fmt.Sprintf("Key=%s,Value=%s", key, value)
		cmd := exec.Command("aws", "ec2", "create-tags", "--resources", vpcID, "--tags", tagSpec)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to tag VPC with %s=%s: %w", key, value, err)
		}
	}

	successColor.Printf("  VPC tagged successfully\n")
	return nil
}

func (e *Engine) deleteVPCWithAWSCLI(resource *state.ResourceState) error {
	vpcID, ok := resource.Attributes["vpc_id"].(string)
	if !ok {
		return fmt.Errorf("vpc_id not found in resource state")
	}

	fmt.Printf("  Deleting VPC: %s\n", vpcID)

	// Delete VPC using AWS CLI
	cmd := exec.Command("aws", "ec2", "delete-vpc", "--vpc-id", vpcID)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("aws cli error: %w", err)
	}

	fmt.Printf("  ✓ VPC deleted: %s\n", vpcID)
	return nil
}

func (e *Engine) createSubnetWithAWSCLI(resource *state.ResourceState, currentState *state.State) error {
	// Get required parameters
	cidrBlock, ok := resource.Attributes["cidr_block"].(string)
	if !ok {
		return fmt.Errorf("cidr_block not found in Subnet configuration")
	}

	availabilityZone, ok := resource.Attributes["availability_zone"].(string)
	if !ok {
		return fmt.Errorf("availability_zone not found in Subnet configuration")
	}

	// Resolve VPC ID from dependency
	vpcRef, ok := resource.Attributes["vpc_id"].(string)
	if !ok {
		return fmt.Errorf("vpc_id not found in Subnet configuration")
	}

	// Find the VPC resource in current state
	var vpcID string
	for _, res := range currentState.Resources {
		if res.Name == vpcRef && res.Type == "vpc" {
			if id, exists := res.Attributes["vpc_id"].(string); exists {
				vpcID = id
				break
			}
		}
	}

	if vpcID == "" {
		return fmt.Errorf("could not resolve VPC ID for reference: %s", vpcRef)
	}

	infoColor.Printf("  Creating Subnet with CIDR: %s in VPC: %s\n", cidrBlock, vpcID)

	cmd := exec.Command("aws", "ec2", "create-subnet",
		"--vpc-id", vpcID,
		"--cidr-block", cidrBlock,
		"--availability-zone", availabilityZone,
		"--output", "json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("aws cli error: %w, output: %s", err, string(output))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return fmt.Errorf("failed to parse AWS CLI output: %w", err)
	}

	subnet, ok := result["Subnet"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid AWS CLI response format")
	}

	subnetID, ok := subnet["SubnetId"].(string)
	if !ok {
		return fmt.Errorf("Subnet ID not found in response")
	}

	resource.Attributes["subnet_id"] = subnetID
	resource.Attributes["vpc_id"] = vpcID // Store resolved VPC ID
	resource.Attributes["state"] = subnet["State"]

	successColor.Printf("  Subnet created with ID: %s\n", subnetID)

	// Configure public IP mapping if specified
	if mapPublicIP, exists := resource.Attributes["map_public_ip"]; exists {
		if mapPublic, ok := mapPublicIP.(bool); ok && mapPublic {
			if err := e.configureSubnetPublicIP(subnetID, true); err != nil {
				fmt.Printf("  Warning: failed to configure public IP mapping: %v\n", err)
			}
		}
	}

	// Add tags if specified
	if tags, exists := resource.Attributes["tags"]; exists {
		if err := e.tagSubnetWithAWSCLI(subnetID, tags); err != nil {
			fmt.Printf("  Warning: failed to tag Subnet: %v\n", err)
		}
	}

	return nil
}

func (e *Engine) configureSubnetPublicIP(subnetID string, mapPublicIP bool) error {
	cmd := exec.Command("aws", "ec2", "modify-subnet-attribute",
		"--subnet-id", subnetID,
		"--map-public-ip-on-launch")

	if !mapPublicIP {
		cmd = exec.Command("aws", "ec2", "modify-subnet-attribute",
			"--subnet-id", subnetID,
			"--no-map-public-ip-on-launch")
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure public IP mapping: %w", err)
	}

	successColor.Printf("  Subnet public IP mapping configured\n")
	return nil
}

func (e *Engine) tagSubnetWithAWSCLI(subnetID string, tags interface{}) error {
	tagsMap, ok := tags.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid tags format")
	}

	for key, value := range tagsMap {
		tagSpec := fmt.Sprintf("Key=%s,Value=%s", key, value)
		cmd := exec.Command("aws", "ec2", "create-tags", "--resources", subnetID, "--tags", tagSpec)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to tag Subnet with %s=%s: %w", key, value, err)
		}
	}

	successColor.Printf("  Subnet tagged successfully\n")
	return nil
}

func (e *Engine) deleteSubnetWithAWSCLI(resource *state.ResourceState) error {
	subnetID, ok := resource.Attributes["subnet_id"].(string)
	if !ok {
		return fmt.Errorf("subnet_id not found in resource state")
	}

	fmt.Printf("  Deleting Subnet: %s\n", subnetID)

	cmd := exec.Command("aws", "ec2", "delete-subnet", "--subnet-id", subnetID)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("aws cli error: %w", err)
	}

	fmt.Printf("  ✓ Subnet deleted: %s\n", subnetID)
	return nil
}
