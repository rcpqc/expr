package eval

import "go/ast"

// Constant constant
type Constant struct {
	ast.BasicLit
	Value any
}
