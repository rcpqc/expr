package eval

import (
	"fmt"
	"go/ast"

	"github.com/rcpqc/expr/builtin"
)

func evalIdent(ident *ast.Ident, variables map[string]interface{}) (interface{}, error) {
	if v, ok := variables[ident.Name]; ok {
		return v, nil
	}
	if fn, ok := builtin.Functions[ident.Name]; ok {
		return fn, nil
	}
	return nil, fmt.Errorf("[ident] unknown ident(%s)", ident.Name)
}
