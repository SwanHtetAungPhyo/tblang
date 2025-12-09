package compiler

import (
	"fmt"

	"github.com/tblang/core/parser"
)

// Helper methods

func (w *ASTWalker) isResourceType(funcName string) bool {
	resourceTypes := []string{"vpc", "subnet", "security_group", "ec2", "internet_gateway", "route_table", "eip", "nat_gateway"}
	for _, rt := range resourceTypes {
		if rt == funcName {
			return true
		}
	}
	return false
}

func (w *ASTWalker) isDataSourceType(funcName string) bool {
	dataSourceTypes := []string{"data_ami", "data_vpc", "data_subnet", "data_availability_zones", "data_caller_identity"}
	for _, dt := range dataSourceTypes {
		if dt == funcName {
			return true
		}
	}
	return false
}

func (w *ASTWalker) extractArguments(argList parser.IArgumentListContext) []interface{} {
	if argList == nil {
		return []interface{}{}
	}

	var args []interface{}
	for _, expr := range argList.(*parser.ArgumentListContext).AllExpression() {
		val := w.evaluateExpression(expr)
		args = append(args, val)
	}
	return args
}

func (w *ASTWalker) extractStringValue(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", value)
}

func (w *ASTWalker) convertToMap(value interface{}) map[string]interface{} {
	if m, ok := value.(map[string]interface{}); ok {
		return m
	}
	return make(map[string]interface{})
}
