package expr

import (
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/builtin"
	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/eval"
	"github.com/rcpqc/expr/types"
)

var (
	// Eval evaluate a compiled expression
	Eval = eval.Eval
	// Comp compile an expression
	Comp = eval.Comp
)

type (
	// Vars a framework's identifier container provides basic built-in functions
	Vars = builtin.Vars
)

// EvalType eval and convert type
func EvalType(expr ast.Expr, variables eval.Variables, t reflect.Type) (any, error) {
	if t == nil {
		t = types.Any
	}
	val, err := Eval(expr, variables)
	if err != nil {
		return val, err
	}
	rv := reflect.ValueOf(val)
	if !rv.IsValid() {
		return reflect.Zero(t).Interface(), nil
	}
	if rv.Type() == t {
		return val, nil
	}
	if rv.CanConvert(t) {
		return rv.Convert(t).Interface(), nil
	}
	return val, errs.Newf(expr, "%v(%v) can't convert to type(%v)", rv.Type(), rv, t)
}

// EvalOr eval otherwise
func EvalOr(expr ast.Expr, variables eval.Variables, defaultValue any) any {
	if cval, err := EvalType(expr, variables, reflect.TypeOf(defaultValue)); err == nil {
		return cval
	}
	return defaultValue
}
