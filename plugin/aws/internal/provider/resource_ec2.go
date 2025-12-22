package provider

import (
"context"

"github.com/tblang/core/pkg/plugin"
)
func (p *AWSProvider) applyEC2(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
	config, ok := req.Config.(map[string]interface{})
	if !ok {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Invalid EC2 configuration",
					Detail:   "Configuration must be a map",
				},
			},
		}, nil
	}

	ami, _ := config["ami"].(string)
	instanceType, _ := config["instance_type"].(string)
	subnetID, _ := config["subnet_id"].(string)

	if ami == "" || instanceType == "" || subnetID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing required fields",
					Detail:   "ami, instance_type, and subnet_id are required for EC2",
				},
			},
		}, nil
	}

	keyName, _ := config["key_name"].(string)
	userData, _ := config["user_data"].(string)
	associatePublicIP, _ := config["associate_public_ip"].(bool)

	var securityGroups []string
	if sgs, exists := config["security_groups"]; exists {
		if sgList, ok := sgs.([]interface{}); ok {
			for _, sg := range sgList {
				if sgStr, ok := sg.(string); ok {
					securityGroups = append(securityGroups, sgStr)
				}
			}
		}
	}

	var rootVolumeSize int32 = 8
	var rootVolumeType string = "gp3"
	if size, exists := config["root_volume_size"]; exists {
		if sizeFloat, ok := size.(float64); ok {
			rootVolumeSize = int32(sizeFloat)
		}
	}
	if volType, exists := config["root_volume_type"]; exists {
		if volTypeStr, ok := volType.(string); ok {
			rootVolumeType = volTypeStr
		}
	}

	instance, err := p.client.CreateEC2Instance(ctx, &EC2InstanceConfig{
		AMI:               ami,
		InstanceType:      instanceType,
		SubnetID:          subnetID,
		SecurityGroupIDs:    securityGroups,
		KeyName:           keyName,
		UserData:          userData,
		AssociatePublicIP: associatePublicIP,
		RootVolumeSize:    rootVolumeSize,
		RootVolumeType:    rootVolumeType,
		Tags:              extractTags(config),
	})
	if err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to create EC2 instance",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	newState := make(map[string]interface{})
	for k, v := range config {
		newState[k] = v
	}
	newState["instance_id"] = instance.InstanceID
	newState["public_ip"] = instance.PublicIP
	newState["private_ip"] = instance.PrivateIP
	newState["state"] = instance.State

	return &plugin.ApplyResourceChangeResponse{
		NewState: newState,
	}, nil
}

func (p *AWSProvider) destroyEC2(ctx context.Context, req *plugin.ApplyResourceChangeRequest) (*plugin.ApplyResourceChangeResponse, error) {
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

	instanceID, ok := priorState["instance_id"].(string)
	if !ok || instanceID == "" {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Missing Instance ID",
					Detail:   "instance_id is required to terminate EC2 instance",
				},
			},
		}, nil
	}

	if err := p.client.TerminateEC2Instance(ctx, instanceID); err != nil {
		return &plugin.ApplyResourceChangeResponse{
			Diagnostics: []*plugin.Diagnostic{
				{
					Severity: "error",
					Summary:  "Failed to terminate EC2 instance",
					Detail:   err.Error(),
				},
			},
		}, nil
	}

	return &plugin.ApplyResourceChangeResponse{
		NewState: nil,
	}, nil
}
