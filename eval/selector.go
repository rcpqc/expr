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

func evalSelectorStruct(rv reflect.Value, key string) (interface{}, error) {
	val := types.NewProfile(rv.Type(), "expr").FieldFromTagName(rv, key)
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
	for rvx.Kind() == reflect.Ptr {
		rvx = rvx.Elem()
	}
	switch rvx.Kind() {
	case reflect.Map:
		return evalSelectorMap(rvx, selector.Sel.Name)
	case reflect.Struct:
		return evalSelectorStruct(rvx, selector.Sel.Name)
	}
	return nil, fmt.Errorf("[selector] illegal kind(%s)", rvx.Kind().String())
}
