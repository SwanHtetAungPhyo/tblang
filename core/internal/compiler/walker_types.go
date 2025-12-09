package compiler

import (
	"github.com/tblang/core/parser"
)

// ASTWalker walks the parse tree and builds the internal AST
type ASTWalker struct {
	*parser.BasetblangListener
	compiler          *Compiler
	variables         map[string]interface{}
	processedContexts map[interface{}]bool // Track contexts we've manually processed
	inManualExecution bool                 // Flag to indicate we're manually executing (e.g., in a loop)
}
