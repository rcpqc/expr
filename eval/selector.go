package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/types"
)

func evalSelectorMap(rv reflect.Value, key string) (interface{}, error) {
	if rv.Type().Key().Kind() != reflect.String {
		return nil, fmt.Errorf("[selector] key of map must be string")
	}
	val := rv.MapIndex(reflect.ValueOf(key))
	if !val.IsValid() || !val.CanInterface() {
		return reflect.Zero(rv.Type().Elem()).Interface(), nil
	}
	return val.Interface(), nil
}

func evalSelectorProfile(rv reflect.Value, key string) (interface{}, error) {
	val := types.NewProfile(rv.Type(), "expr").Select(rv, key)
	if !val.IsValid() || !val.CanInterface() {
		return nil, fmt.Errorf("[selector] field(%s) not found", key)
	}
	return val.Interface(), nil
}

func evalSelector(selector *ast.SelectorExpr, variables Variables) (interface{}, error) {
	x, err := eval(selector.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	if rvx.Kind() == reflect.Invalid {
		return nil, fmt.Errorf("[selector] illegal kind(%s)", rvx.Kind().String())
	}
	if rvx.Kind() == reflect.Map {
		return evalSelectorMap(rvx, selector.Sel.Name)
	}
	return evalSelectorProfile(rvx, selector.Sel.Name)
}
