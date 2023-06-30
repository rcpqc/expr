package eval

import "go/ast"

// Constant 常数
type Constant struct {
	ast.BasicLit
	Value interface{}
}
