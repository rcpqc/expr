package eval

import (
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

var (
	Eval     = eval
	EvalType = evaltype
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
	if compositelit, ok := expr.(*ast.CompositeLit); ok {
		return evalCompositeLit(compositelit, variables)
	}
	return nil, errs.Newf(expr, "unsupported expression type(%v)", reflect.TypeOf(expr))
}

func evaltype(expr ast.Expr, variables Variables, t reflect.Type) (any, error) {
	if t == nil {
		t = types.Any
	}
	val, err := eval(expr, variables)
	if err != nil {
		return nil, err
	}
	rv := reflect.ValueOf(val)
	if !rv.IsValid() {
		return reflect.Zero(t).Interface(), nil
	}
	if rv.Type() == t {
		return val, nil
	}
	if rv.CanConvert(t) {
		return rv.Convert(t).Interface(), nil
	}
	return val, errs.Newf(expr, "%v(%v) can't convert to type(%v)", rv.Type(), rv, t)
}
