package compiler

// This file serves as the main entry point for the walker functionality.
// All functionality has been split into separate files for better organization:
//
// - walker_types.go: Type definitions for ASTWalker
// - walker_block.go: Block declaration handling (cloud_vendor)
// - walker_variable.go: Variable declaration handling
// - walker_loop.go: For loop handling and statement execution
// - walker_function.go: Function call handling (resources, data sources)
// - walker_print.go: Print and output function handling
// - walker_helpers.go: Helper methods (type checking, argument extraction)
// - walker_expression.go: Expression evaluation (literals, objects, arrays)
