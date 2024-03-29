package eval

import (
	"go/ast"
	"go/parser"
)

func compBasicLit(basic *ast.BasicLit) ast.Expr {
	value, err := eval(basic, nil)
	if err != nil {
		return basic
	}
	return &Constant{BasicLit: *basic, Value: value}
}

func compCompositeLit(composite *ast.CompositeLit) ast.Expr {
	value, err := eval(composite, nil)
	if err != nil {
		return composite
	}
	return &Constant{Value: value}
}

func compIdent(ident *ast.Ident) ast.Expr {
	if ident.Name == "true" {
		return &Constant{Value: true}
	}
	if ident.Name == "false" {
		return &Constant{Value: false}
	}
	if ident.Name == "nil" {
		return &Constant{Value: nil}
	}
	return ident
}

func comp(expr ast.Expr) ast.Expr {
	if basiclit, ok := expr.(*ast.BasicLit); ok {
		return compBasicLit(basiclit)
	}
	if compositelit, ok := expr.(*ast.CompositeLit); ok {
		return compCompositeLit(compositelit)
	}
	if ident, ok := expr.(*ast.Ident); ok {
		return compIdent(ident)
	}
	if binary, ok := expr.(*ast.BinaryExpr); ok {
		binary.X = comp(binary.X)
		binary.Y = comp(binary.Y)
		return expr
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		selector.X = comp(selector.X)
		selector.Sel.Obj = &ast.Object{Data: selector.Sel.Name}
		return expr
	}
	if paren, ok := expr.(*ast.ParenExpr); ok {
		return comp(paren.X)
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
