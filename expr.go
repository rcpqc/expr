package expr

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/eval"
)

// Vars 变量
type Vars map[string]interface{}

// Get 获取参数
func (o Vars) Get(name string) (interface{}, bool) {
	val, ok := o[name]
	return val, ok
}

var (
	Eval = eval.Eval
)

// EvalType eval and convert type
func EvalType(expr ast.Expr, variables eval.Variables, target interface{}) (interface{}, error) {
	val, err := Eval(expr, variables)
	if err != nil {
		return val, err
	}
	rv := reflect.ValueOf(val)
	rtv := reflect.ValueOf(target)
	if !rtv.IsValid() {
		return nil, nil
	}
	if rv.Type() == rtv.Type() {
		return val, nil
	}
	if rv.CanConvert(rtv.Type()) {
		return rv.Convert(rtv.Type()).Interface(), nil
	}
	return val, fmt.Errorf("%v can't convert to type(%v)", rv, rtv.Type())
}

// EvalOr eval otherwise
func EvalOr(expr ast.Expr, variables eval.Variables, defaultValue interface{}) interface{} {
	if cval, err := EvalType(expr, variables, defaultValue); err == nil {
		return cval
	}
	return defaultValue
}
