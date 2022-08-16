package expr

import (
	"fmt"
	"go/ast"
	"reflect"
)

func Eval(expr ast.Expr, variables map[string]interface{}) (interface{}, error) {
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
	}
	return nil, fmt.Errorf("unsupported exprtype(%v)", rtexpr)
}
