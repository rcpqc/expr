package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/types"
)

var (
	Eval        = eval
	EvalBool    = evalBool
	EvalInt     = evalInt
	EvalInt8    = evalInt8
	EvalInt16   = evalInt16
	EvalInt32   = evalInt32
	EvalInt64   = evalInt64
	EvalUint    = evalUint
	EvalUint8   = evalUint8
	EvalUint16  = evalUint16
	EvalUint32  = evalUint32
	EvalUint64  = evalUint64
	EvalFloat32 = evalFloat32
	EvalFloat64 = evalFloat64
	EvalString  = evalString
	EvalBytes   = evalBytes
)

// Variables 变量接口
type Variables interface {
	Get(string) (interface{}, bool)
}

var handlers = map[reflect.Type]func(expr ast.Expr, variables Variables) (interface{}, error){}

func init() {
	handlers[reflect.TypeOf((*ast.BinaryExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalBinary(expr.(*ast.BinaryExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.Ident)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalIdent(expr.(*ast.Ident), variables)
	}
	handlers[reflect.TypeOf((*ast.BasicLit)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalBasicLit(expr.(*ast.BasicLit), variables)
	}
	handlers[reflect.TypeOf((*ast.UnaryExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalUnary(expr.(*ast.UnaryExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.CallExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalCall(expr.(*ast.CallExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.ParenExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalParen(expr.(*ast.ParenExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.SelectorExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalSelector(expr.(*ast.SelectorExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.SliceExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalSlice(expr.(*ast.SliceExpr), variables)
	}
	handlers[reflect.TypeOf((*ast.IndexExpr)(nil))] = func(expr ast.Expr, variables Variables) (interface{}, error) {
		return evalIndex(expr.(*ast.IndexExpr), variables)
	}
}

func eval(expr ast.Expr, variables Variables) (interface{}, error) {
	rtexpr := reflect.TypeOf(expr)
	if handler, ok := handlers[rtexpr]; ok {
		return handler(expr, variables)
	}
	return nil, fmt.Errorf("unsupported exprtype(%v)", rtexpr)
}

func evalBool(expr ast.Expr, variables Variables) (bool, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return false, err
	}
	if !reflect.ValueOf(val).CanConvert(types.BoolType) {
		return false, fmt.Errorf("type(%v) can't convert to bool", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.BoolType).Interface().(bool), nil
}

func evalInt(expr ast.Expr, variables Variables) (int, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.IntType) {
		return 0, fmt.Errorf("type(%v) can't convert to int", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.IntType).Interface().(int), nil
}

func evalInt8(expr ast.Expr, variables Variables) (int8, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Int8Type) {
		return 0, fmt.Errorf("type(%v) can't convert to int8", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Int8Type).Interface().(int8), nil
}

func evalInt16(expr ast.Expr, variables Variables) (int16, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Int16Type) {
		return 0, fmt.Errorf("type(%v) can't convert to int16", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Int16Type).Interface().(int16), nil
}

func evalInt32(expr ast.Expr, variables Variables) (int32, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Int32Type) {
		return 0, fmt.Errorf("type(%v) can't convert to int32", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Int32Type).Interface().(int32), nil
}
func evalInt64(expr ast.Expr, variables Variables) (int64, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Int64Type) {
		return 0, fmt.Errorf("type(%v) can't convert to int64", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Int64Type).Interface().(int64), nil
}

func evalUint(expr ast.Expr, variables Variables) (uint, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.UintType) {
		return 0, fmt.Errorf("type(%v) can't convert to uint", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.UintType).Interface().(uint), nil
}

func evalUint8(expr ast.Expr, variables Variables) (uint8, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Uint8Type) {
		return 0, fmt.Errorf("type(%v) can't convert to uint8", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Uint8Type).Interface().(uint8), nil
}

func evalUint16(expr ast.Expr, variables Variables) (uint16, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Uint16Type) {
		return 0, fmt.Errorf("type(%v) can't convert to uint16", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Uint16Type).Interface().(uint16), nil
}
func evalUint32(expr ast.Expr, variables Variables) (uint32, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Uint32Type) {
		return 0, fmt.Errorf("type(%v) can't convert to uint32", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Uint32Type).Interface().(uint32), nil
}
func evalUint64(expr ast.Expr, variables Variables) (uint64, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Uint64Type) {
		return 0, fmt.Errorf("type(%v) can't convert to uint64", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Uint64Type).Interface().(uint64), nil
}

func evalFloat32(expr ast.Expr, variables Variables) (float32, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Float32Type) {
		return 0, fmt.Errorf("type(%v) can't convert to float32", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Float32Type).Interface().(float32), nil
}

func evalFloat64(expr ast.Expr, variables Variables) (float64, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Float64Type) {
		return 0, fmt.Errorf("type(%v) can't convert to float64", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Float64Type).Interface().(float64), nil
}

func evalString(expr ast.Expr, variables Variables) (string, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return "", err
	}
	if !reflect.ValueOf(val).CanConvert(types.StringType) {
		return "", fmt.Errorf("type(%v) can't convert to string", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.StringType).Interface().(string), nil
}

func evalBytes(expr ast.Expr, variables Variables) ([]byte, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return nil, err
	}
	if !reflect.ValueOf(val).CanConvert(types.BytesType) {
		return nil, fmt.Errorf("type(%v) can't convert to bytes", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.BytesType).Interface().([]byte), nil
}
