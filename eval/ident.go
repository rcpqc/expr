package eval

import (
	"fmt"
	"go/ast"

	"github.com/rcpqc/expr/errs"
)

func evalIdent(ident *ast.Ident, variables Variables) (any, error) {
	if variables == nil {
		return nil, fmt.Errorf("variables == nil")
	}
	v, err := variables.Get(ident.Name)
	return v, errs.New(ident, err)
}
