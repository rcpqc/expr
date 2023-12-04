package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

func evalIndexArray(rvx reflect.Value, index ast.Expr, variables Variables) (any, error) {
	val, err := evaluator(index)(index, variables)
	if err != nil {
		return nil, err
	}
	idx, ok := types.ConvertInt(val)
	if !ok {
		return nil, fmt.Errorf("index must be an integer")
	}
	if idx < 0 {
		idx += rvx.Len()
	}
	if idx < 0 || idx >= rvx.Len() {
		return nil, fmt.Errorf("out of range index(%d) for len(%d)", idx, rvx.Len())
	}
	return rvx.Index(idx).Interface(), nil
}

func evalIndexMap(rvx reflect.Value, index ast.Expr, variables Variables) (any, error) {
	key, err := evaluator(index)(index, variables)
	if err != nil {
		return nil, err
	}
	rkey := reflect.ValueOf(key)
	if !rkey.IsValid() {
		return reflect.Zero(rvx.Type().Elem()).Interface(), nil
	}
	if !rkey.Type().AssignableTo(rvx.Type().Key()) {
		return nil, fmt.Errorf("%v can't index by key(%v)", rvx.Type(), rkey.Type())
	}
	val := rvx.MapIndex(rkey)
	if !val.IsValid() || !val.CanInterface() {
		return reflect.Zero(rvx.Type().Elem()).Interface(), nil
	}
	return val.Interface(), nil
}

func evalIndex(expr ast.Expr, variables Variables) (any, error) {
	index := expr.(*ast.IndexExpr)
	x, err := evaluator(index.X)(index.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	switch rvx.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		val, err := evalIndexArray(rvx, index.Index, variables)
		return val, errs.New(index, err)
	case reflect.Map:
		val, err := evalIndexMap(rvx, index.Index, variables)
		return val, errs.New(index, err)
	default:
		return nil, errs.Newf(index, "illegal kind(%s)", rvx.Kind())
	}
}
