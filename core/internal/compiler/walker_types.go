package compiler

import (
	"github.com/tblang/core/parser"
)

type ASTWalker struct {
	*parser.BasetblangListener
	compiler          *Compiler
	variables         map[string]interface{}
	processedContexts map[interface{}]bool
	inManualExecution bool
}
