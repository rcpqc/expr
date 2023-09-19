package eval

import (
	"go/ast"
)

func evalParen(paren *ast.ParenExpr, variables Variables) (any, error) {
	x, err := eval(paren.X, variables)
	if err != nil {
		return nil, err
	}
	return x, nil
}
