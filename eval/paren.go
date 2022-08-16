package eval

import (
	"go/ast"
)

func evalParen(paren *ast.ParenExpr, variables map[string]interface{}) (interface{}, error) {
	x, err := evalExpr(paren.X, variables)
	if err != nil {
		return nil, err
	}
	return x, nil
}
