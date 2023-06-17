package comp

import (
	"go/ast"
	"go/parser"

	"github.com/rcpqc/expr/eval"
	"github.com/rcpqc/expr/types"
)

func compBasicLit(basic *ast.BasicLit) ast.Expr {
	value, err := eval.Eval(basic, nil)
	if err != nil {
		return basic
	}
	return &types.Constant{BasicLit: *basic, Value: value}
}

func comp(expr ast.Expr) ast.Expr {
	if basiclit, ok := expr.(*ast.BasicLit); ok {
		return compBasicLit(basiclit)
	}
	if _, ok := expr.(*ast.Ident); ok {
		return expr
	}
	if binary, ok := expr.(*ast.BinaryExpr); ok {
		binary.X = comp(binary.X)
		binary.Y = comp(binary.Y)
		return expr
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		selector.X = comp(selector.X)
		return expr
	}
	if paren, ok := expr.(*ast.ParenExpr); ok {
		paren.X = comp(paren.X)
		return expr
	}
	if call, ok := expr.(*ast.CallExpr); ok {
		call.Fun = comp(call.Fun)
		for i := 0; i < len(call.Args); i++ {
			call.Args[i] = comp(call.Args[i])
		}
		return expr
	}
	if unary, ok := expr.(*ast.UnaryExpr); ok {
		unary.X = comp(unary.X)
		return expr
	}
	if index, ok := expr.(*ast.IndexExpr); ok {
		index.X = comp(index.X)
		index.Index = comp(index.Index)
		return expr
	}
	if slice, ok := expr.(*ast.SliceExpr); ok {
		slice.X = comp(slice.X)
		slice.Low = comp(slice.Low)
		slice.High = comp(slice.High)
		return expr
	}
	return expr
}

// Comp compile an expression
func Comp(x string) (ast.Expr, error) {
	expr, err := parser.ParseExpr(x)
	if err != nil {
		return nil, err
	}
	return comp(expr), nil
}