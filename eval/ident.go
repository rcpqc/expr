package eval

import (
	"errors"
	"go/ast"

	"github.com/rcpqc/expr/errs"
)

func evalIdent(expr ast.Expr, variables Variables) (any, error) {
	ident := expr.(*ast.Ident)
	if variables == nil {
		return nil, errs.New(ident, errors.New("variables == nil"))
	}
	v, err := variables.Get(ident.Name)
	return v, errs.New(ident, err)
}
