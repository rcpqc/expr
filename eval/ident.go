package eval

import (
	"go/ast"

	"github.com/rcpqc/expr/errs"
)

func evalIdent(ident *ast.Ident, variables Variables) (any, error) {
	v, err := variables.Get(ident.Name)
	return v, errs.New(ident, err)
}
