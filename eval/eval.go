package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/types"
)

var (
	Eval        = evalExpr
	EvalInt     = evalInt
	EvalInt64   = evalInt64
	EvalFloat64 = evalFloat64
	EvalString  = evalString
)

func evalExpr(expr ast.Expr, variables map[string]interface{}) (interface{}, error) {
	rtexpr := reflect.TypeOf(expr)
	switch rtexpr {
	case reflect.TypeOf((*ast.BinaryExpr)(nil)):
		return evalBinary(expr.(*ast.BinaryExpr), variables)
	case reflect.TypeOf((*ast.Ident)(nil)):
		return evalIdent(expr.(*ast.Ident), variables)
	case reflect.TypeOf((*ast.BasicLit)(nil)):
		return evalBasicLit(expr.(*ast.BasicLit), variables)
	case reflect.TypeOf((*ast.UnaryExpr)(nil)):
		return evalUnary(expr.(*ast.UnaryExpr), variables)
	case reflect.TypeOf((*ast.CallExpr)(nil)):
		return evalCall(expr.(*ast.CallExpr), variables)
	case reflect.TypeOf((*ast.ParenExpr)(nil)):
		return evalParen(expr.(*ast.ParenExpr), variables)
	case reflect.TypeOf((*ast.SelectorExpr)(nil)):
		return evalSelector(expr.(*ast.SelectorExpr), variables)
	case reflect.TypeOf((*ast.SliceExpr)(nil)):
		return evalSlice(expr.(*ast.SliceExpr), variables)
	case reflect.TypeOf((*ast.IndexExpr)(nil)):
		return evalIndex(expr.(*ast.IndexExpr), variables)
	}
	return nil, fmt.Errorf("unsupported exprtype(%v)", rtexpr)
}

func evalInt(expr ast.Expr, variables map[string]interface{}) (int, error) {
	val, err := evalExpr(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.IntType) {
		return 0, fmt.Errorf("type(%v) cannot convert to int", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.IntType).Interface().(int), nil
}

func evalInt64(expr ast.Expr, variables map[string]interface{}) (int64, error) {
	val, err := evalExpr(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Int64Type) {
		return 0, fmt.Errorf("type(%v) cannot convert to int64", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Int64Type).Interface().(int64), nil
}

func evalFloat64(expr ast.Expr, variables map[string]interface{}) (float64, error) {
	val, err := evalExpr(expr, variables)
	if err != nil {
		return 0, err
	}
	if !reflect.ValueOf(val).CanConvert(types.Float64Type) {
		return 0, fmt.Errorf("type(%v) cannot convert to float64", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.Float64Type).Interface().(float64), nil
}

func evalString(expr ast.Expr, variables map[string]interface{}) (string, error) {
	val, err := evalExpr(expr, variables)
	if err != nil {
		return "", err
	}
	if !reflect.ValueOf(val).CanConvert(types.StringType) {
		return "", fmt.Errorf("type(%v) cannot convert to string", reflect.TypeOf(val))
	}
	return reflect.ValueOf(val).Convert(types.StringType).Interface().(string), nil
}
