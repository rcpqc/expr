package eval

import "go/ast"

// Constant constant
type Constant struct {
	ast.BasicLit
	Value any
}

func evalConstant(expr ast.Expr, variables Variables) (any, error) {
	return expr.(*Constant).Value, nil
}
