package provider

// extractTags extracts tags from configuration
func extractTags(config map[string]interface{}) map[string]string {
	tags := make(map[string]string)

	if tagsInterface, exists := config["tags"]; exists {
		if tagsMap, ok := tagsInterface.(map[string]interface{}); ok {
			for key, value := range tagsMap {
				if strValue, ok := value.(string); ok {
					tags[key] = strValue
				}
			}
		}
	}

	return tags
}
