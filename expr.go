package expr

import (
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/eval"
)

var (
	// Eval evaluate a compiled expression
	Eval = eval.Eval
	// EvalType eval and convert type
	EvalType = eval.EvalType
	// Comp compile an expression
	Comp = eval.Comp
)

type (
	// Vars a framework's identifier container provides basic built-in functions
	Vars = eval.Vars
)

// EvalOr eval otherwise
func EvalOr(expr ast.Expr, variables eval.Variables, defaultValue any) any {
	if cval, err := EvalType(expr, variables, reflect.TypeOf(defaultValue)); err == nil {
		return cval
	}
	return defaultValue
}
