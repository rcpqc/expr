package eval

import (
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
)

var (
	Eval = eval
)

// Variables variable getter
type Variables interface {
	Get(string) (any, error)
}

func eval(expr ast.Expr, variables Variables) (any, error) {
	if constant, ok := expr.(*Constant); ok {
		return constant.Value, nil
	}
	if ident, ok := expr.(*ast.Ident); ok {
		return evalIdent(ident, variables)
	}
	if binary, ok := expr.(*ast.BinaryExpr); ok {
		return evalBinary(binary, variables)
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		return evalSelector(selector, variables)
	}
	if paren, ok := expr.(*ast.ParenExpr); ok {
		return evalParen(paren, variables)
	}
	if call, ok := expr.(*ast.CallExpr); ok {
		return evalCall(call, variables)
	}
	if unary, ok := expr.(*ast.UnaryExpr); ok {
		return evalUnary(unary, variables)
	}
	if index, ok := expr.(*ast.IndexExpr); ok {
		return evalIndex(index, variables)
	}
	if slice, ok := expr.(*ast.SliceExpr); ok {
		return evalSlice(slice, variables)
	}
	if basiclit, ok := expr.(*ast.BasicLit); ok {
		return evalBasicLit(basiclit, variables)
	}
	return nil, errs.Newf(expr, "unsupported expression type(%v)", reflect.TypeOf(expr))
}
