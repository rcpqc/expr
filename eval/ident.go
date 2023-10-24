package eval

import (
	"errors"
	"go/ast"

	"github.com/rcpqc/expr/errs"
)

func evalIdent(ident *ast.Ident, variables Variables) (any, error) {
	if variables == nil {
		return nil, errs.New(ident, errors.New("variables == nil"))
	}
	v, err := variables.Get(ident.Name)
	return v, errs.New(ident, err)
}
