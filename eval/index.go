package eval

import (
	"fmt"
	"go/ast"
	"reflect"
)

func evalIndexArray(rvx reflect.Value, index ast.Expr, variables Variables) (interface{}, error) {
	idx, err := evalint(index, variables)
	if err != nil {
		return nil, err
	}
	if idx < 0 {
		idx += rvx.Len()
	}
	if idx < 0 || idx > rvx.Len() {
		return nil, fmt.Errorf("[index] out of range index(%d) for len(%d)", idx, rvx.Len())
	}
	return rvx.Index(idx).Interface(), nil
}

func evalIndexMap(rvx reflect.Value, index ast.Expr, variables Variables) (interface{}, error) {
	key, err := eval(index, variables)
	if err != nil {
		return nil, err
	}
	rkey := reflect.ValueOf(key)
	if !rkey.IsValid() {
		return reflect.Zero(rvx.Type().Elem()).Interface(), nil
	}
	if !rkey.Type().AssignableTo(rvx.Type().Key()) {
		return nil, fmt.Errorf("[index] %v can't index by key(%v)", rvx.Type(), rkey.Type())
	}
	val := rvx.MapIndex(rkey)
	if !val.IsValid() || !val.CanInterface() {
		return reflect.Zero(rvx.Type().Elem()).Interface(), nil
	}
	return val.Interface(), nil
}

func evalIndex(index *ast.IndexExpr, variables Variables) (interface{}, error) {
	x, err := eval(index.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	switch rvx.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return evalIndexArray(rvx, index.Index, variables)
	case reflect.Map:
		return evalIndexMap(rvx, index.Index, variables)
	default:
		return nil, fmt.Errorf("[index] illegal kind(%s)", rvx.Kind())
	}
}
