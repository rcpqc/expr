package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/types"
)

func evalIndexArray(rvx reflect.Value, i interface{}) (interface{}, error) {
	ri := reflect.ValueOf(i)
	if !ri.IsValid() || !ri.CanConvert(types.IntType) {
		return nil, fmt.Errorf("[index] index(%v) can't convert to int", ri.Type())
	}
	idx := int(ri.Int())
	if idx < 0 {
		idx += rvx.Len()
	}
	if idx < 0 || idx > rvx.Len() {
		return nil, fmt.Errorf("[index] out of range index(%d) for len(%d)", idx, rvx.Len())
	}
	return rvx.Index(idx).Interface(), nil
}

func evalIndexMap(rvx reflect.Value, key interface{}) (interface{}, error) {
	rkey := reflect.ValueOf(key)
	if rvx.Type().Key() != rkey.Type() {
		return nil, fmt.Errorf("[index] %v can't index by key(%v)", rvx.Type(), rkey.Type())
	}
	val := rvx.MapIndex(reflect.ValueOf(key))
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
	idx, err := eval(index.Index, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	switch rvx.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return evalIndexArray(rvx, idx)
	case reflect.Map:
		return evalIndexMap(rvx, idx)
	default:
		return nil, fmt.Errorf("[index] illegal kind(%s)", rvx.Kind())
	}
}
