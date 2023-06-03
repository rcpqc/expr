package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

func sliceRange(low ast.Expr, high ast.Expr, len int, variables Variables) (int, int, error) {
	idxl, idxh := 0, len
	if low != nil {
		val, err := eval(low, variables)
		if err != nil {
			return 0, 0, err
		}
		idx, ok := types.ConvertInt(val)
		if !ok {
			return 0, 0, fmt.Errorf("low index must be an integer")
		}
		idxl = idx
	}
	if high != nil {
		val, err := eval(high, variables)
		if err != nil {
			return 0, 0, err
		}
		idx, ok := types.ConvertInt(val)
		if !ok {
			return 0, 0, fmt.Errorf("high index must be an integer")
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

func evalSlice(slice *ast.SliceExpr, variables Variables) (interface{}, error) {
	x, err := eval(slice.X, variables)
	if err != nil {
		return nil, err
	}
	rvx := reflect.ValueOf(x)
	if rvx.Kind() != reflect.Slice && rvx.Kind() != reflect.Array && rvx.Kind() != reflect.String {
		return nil, errs.Newf(slice, "illegal kind(%s)", rvx.Kind())
	}
	i, j, err := sliceRange(slice.Low, slice.High, rvx.Len(), variables)
	if err != nil {
		return nil, errs.New(slice, err)
	}
	return rvx.Slice(i, j).Interface(), nil
}
