package eval

import (
	"fmt"
	"go/ast"
	"reflect"
)

func sliceRange(low ast.Expr, high ast.Expr, len int, variables map[string]interface{}) (int, int, error) {
	idxl, idxh := 0, len
	if low != nil {
		idx, err := evalInt(low, variables)
		if err != nil {
			return 0, 0, fmt.Errorf("low err: %v", err)
		}
		idxl = idx
	}
	if high != nil {
		idx, err := evalInt(high, variables)
		if err != nil {
			return 0, 0, fmt.Errorf("high err: %v", err)
		}
		idxh = idx
	}
	if idxl < 0 {
		idxl += len
	}
	if idxh < 0 {
		idxh += len
	}
	if idxl < 0 || idxl > len || idxh < 0 || idxh > len {
		return 0, 0, fmt.Errorf("out of range index(%d:%d) for len(%d)", idxl, idxh, len)
	}
	return idxl, idxh, nil
}

func evalSlice(slice *ast.SliceExpr, variables map[string]interface{}) (interface{}, error) {
	x, err := eval(slice.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	if rvx.Kind() != reflect.Slice && rvx.Kind() != reflect.Array && rvx.Kind() != reflect.String {
		return nil, fmt.Errorf("[slice] illegal kind(%s)", rvx.Kind())
	}
	i, j, err := sliceRange(slice.Low, slice.High, rvx.Len(), variables)
	if err != nil {
		return nil, fmt.Errorf("[slice] range err: %v", err)
	}
	return rvx.Slice(i, j).Interface(), nil
}
