package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

func evalSelectorMap(rv reflect.Value, key any) (any, error) {
	if rv.Type().Key().Kind() != reflect.String {
		return nil, fmt.Errorf("key of map must be string")
	}
	val := rv.MapIndex(reflect.ValueOf(key))
	if !val.IsValid() || !val.CanInterface() {
		return reflect.Zero(rv.Type().Elem()).Interface(), nil
	}
	return val.Interface(), nil
}

func evalSelectorProfile(rv reflect.Value, key string) (any, error) {
	val, ok := types.NewProfile(rv.Type(), "expr").Select(rv, key)
	if !ok {
		return nil, fmt.Errorf("sel(%s) not found", key)
	}
	if !val.IsValid() || !val.CanInterface() {
		return nil, fmt.Errorf("nil value")
	}
	return val.Interface(), nil
}

func evalSelector(expr ast.Expr, variables Variables) (any, error) {
	selector := expr.(*ast.SelectorExpr)
	x, err := evaluator(selector.X)(selector.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	if rvx.Kind() == reflect.Invalid {
		return nil, errs.Newf(selector, "illegal kind(%s)", rvx.Kind().String())
	}
	if rvx.Kind() == reflect.Map {
		if selector.Sel.Obj != nil {
			val, err := evalSelectorMap(rvx, selector.Sel.Obj.Data)
			return val, errs.New(selector, err)
		}
		val, err := evalSelectorMap(rvx, selector.Sel.Name)
		return val, errs.New(selector, err)
	}
	val, err := evalSelectorProfile(rvx, selector.Sel.Name)
	return val, errs.New(selector, err)
}
