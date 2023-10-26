package eval

import (
	"go/ast"

	"github.com/rcpqc/expr/builtin"
	"github.com/rcpqc/expr/errs"
)

func evalIdent(expr ast.Expr, variables Variables) (any, error) {
	ident := expr.(*ast.Ident)
	variable, err := variables.Get(ident.Name)
	if err == nil {
		return variable, nil
	}
	if constant, ok := builtin.Get(ident.Name); ok {
		return constant, nil
	}
	return nil, errs.New(ident, err)
}
