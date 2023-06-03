package eval

import (
	"go/ast"

	"github.com/rcpqc/expr/builtin"
	"github.com/rcpqc/expr/errs"
)

func evalIdent(ident *ast.Ident, variables Variables) (interface{}, error) {
	if v, ok := variables.Get(ident.Name); ok {
		return v, nil
	}
	if fn, ok := builtin.Functions[ident.Name]; ok {
		return fn, nil
	}
	return nil, errs.Newf(ident, "unknown name(%s)", ident.Name)
}
