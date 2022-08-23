package eval

import (
	"fmt"
	"go/ast"
	"reflect"
)

func evalIndex(index *ast.IndexExpr, variables map[string]interface{}) (interface{}, error) {
	x, err := evalExpr(index.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	if rvx.Kind() != reflect.Slice && rvx.Kind() != reflect.Array && rvx.Kind() != reflect.String {
		return nil, fmt.Errorf("[index] illegal kind(%s)", rvx.Kind())
	}
	idx, err := evalInt(index.Index, variables)
	if err != nil {
		return nil, fmt.Errorf("[index] index err: %v", err)
	}
	if idx < 0 {
		idx += rvx.Len()
	}
	if idx < 0 || idx > rvx.Len() {
		return 0, fmt.Errorf("[index] out of range index(%d) for len(%d)", idx, rvx.Len())
	}
	return rvx.Index(idx).Interface(), nil
}
