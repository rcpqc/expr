package eval

import (
	"fmt"
	"go/ast"
	"reflect"
)

var (
	Eval = eval
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
	return nil, fmt.Errorf("unsupported expression type(%v)", rtexpr)
}

func evalint(expr ast.Expr, variables Variables) (int, error) {
	val, err := eval(expr, variables)
	if err != nil {
		return 0, err
	}
	rv := reflect.ValueOf(val)
	if rv.CanInt() {
		return int(rv.Int()), nil
	}
	if rv.CanUint() {
		return int(rv.Uint()), nil
	}
	return 0, fmt.Errorf("%v can't convert to an integer", rv)
}
