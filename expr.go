package expr

import (
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/comp"
	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/eval"
	"github.com/rcpqc/expr/types"
)

// Vars 变量
type Vars map[string]interface{}

// Get 获取参数
func (o Vars) Get(name string) (interface{}, bool) {
	val, ok := o[name]
	return val, ok
}

var (
	// Eval evaluate a compiled expression
	Eval = eval.Eval
	// Comp compile an expression
	Comp = comp.Comp
)

// EvalType eval and convert type
func EvalType(expr ast.Expr, variables eval.Variables, t reflect.Type) (interface{}, error) {
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
func EvalOr(expr ast.Expr, variables eval.Variables, defaultValue interface{}) interface{} {
	if cval, err := EvalType(expr, variables, reflect.TypeOf(defaultValue)); err == nil {
		return cval
	}
	return defaultValue
}
