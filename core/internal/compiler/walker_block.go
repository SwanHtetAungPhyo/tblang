package compiler

import (
	"fmt"
	"strings"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

func (w *ASTWalker) EnterBlockDeclaration(ctx *parser.BlockDeclarationContext) {
	blockType := ctx.IDENTIFIER().GetText()
	blockName := strings.Trim(ctx.STRING_LITERAL().GetText(), `"'`)

	if blockType == "cloud_vendor" {

		properties := make(map[string]interface{})

		for _, prop := range ctx.AllProperty() {
			propCtx := prop.(*parser.PropertyContext)
			key := propCtx.IDENTIFIER().GetText()
			value := w.evaluateExpression(propCtx.Expression())
			properties[key] = value
		}

		cloudVendor := &ast.CloudVendor{
			Name:       blockName,
			Properties: properties,
		}

		w.compiler.cloudVendors[blockName] = cloudVendor
		fmt.Printf("Registered cloud vendor: %s\n", blockName)
	}
}
