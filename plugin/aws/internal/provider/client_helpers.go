package provider

import (
"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func (c *AWSClient) buildTags(resourceName string, additionalTags map[string]string) []types.Tag {

	tagMap := make(map[string]string)
	for key, value := range additionalTags {
		tagMap[key] = value
	}

	if _, exists := tagMap["ManagedBy"]; !exists {
		tagMap["ManagedBy"] = "TBLang"
	}

	var tags []types.Tag
	for key, value := range tagMap {
		tags = append(tags, types.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		})
	}

	return tags
}
