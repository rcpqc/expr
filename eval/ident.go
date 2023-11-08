package eval

import (
	"go/ast"

	"github.com/rcpqc/expr/errs"
)

func evalIdent(expr ast.Expr, variables Variables) (any, error) {
	ident := expr.(*ast.Ident)
	val, err := variables.Get(ident.Name)
	if err == nil {
		return val, nil
	}
	return nil, errs.New(ident, err)
}
