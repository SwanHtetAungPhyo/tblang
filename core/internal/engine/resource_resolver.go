package engine

func (e *Engine) resolveResourceReferences(attrs map[string]interface{}) map[string]interface{} {
	resolved := make(map[string]interface{})

	currentState, err := e.stateManager.LoadState()
	if err != nil {
		return attrs
	}

	for key, value := range attrs {

		if strValue, ok := value.(string); ok {

			if resource, exists := currentState.Resources[strValue]; exists {

				switch key {
				case "vpc_id":
					if vpcID, ok := resource.Attributes["vpc_id"].(string); ok {
						resolved[key] = vpcID
						continue
					}
				case "subnet_id":
					if subnetID, ok := resource.Attributes["subnet_id"].(string); ok {
						resolved[key] = subnetID
						continue
					}
				case "group_id":
					if groupID, ok := resource.Attributes["group_id"].(string); ok {
						resolved[key] = groupID
						continue
					}
				case "allocation_id":
					if allocID, ok := resource.Attributes["allocation_id"].(string); ok {
						resolved[key] = allocID
						continue
					}
				case "gateway_id":
					if gwID, ok := resource.Attributes["gateway_id"].(string); ok {
						resolved[key] = gwID
						continue
					}
				case "nat_gateway_id":
					if natGwID, ok := resource.Attributes["nat_gateway_id"].(string); ok {
						resolved[key] = natGwID
						continue
					}
				}
			}
		}

		if arrValue, ok := value.([]interface{}); ok {
			resolvedArr := make([]interface{}, len(arrValue))
			for i, item := range arrValue {
				if strItem, ok := item.(string); ok {

					if resource, exists := currentState.Resources[strItem]; exists {

						switch resource.Type {
						case "security_group":
							if groupID, ok := resource.Attributes["group_id"].(string); ok {
								resolvedArr[i] = groupID
								continue
							}
						case "subnet":
							if subnetID, ok := resource.Attributes["subnet_id"].(string); ok {
								resolvedArr[i] = subnetID
								continue
							}
						}
					}
				}
				resolvedArr[i] = item
			}
			resolved[key] = resolvedArr
			continue
		}

		resolved[key] = value
	}

	return resolved
}
