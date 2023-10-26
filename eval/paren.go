package eval

import (
	"go/ast"
)

func evalParen(expr ast.Expr, variables Variables) (any, error) {
	paren := expr.(*ast.ParenExpr)
	x, err := evaluator(paren.X)(paren.X, variables)
	if err != nil {
		return nil, err
	}
	return x, nil
}
