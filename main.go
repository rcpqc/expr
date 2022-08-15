package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"log"
	"reflect"
)

func main() {
	variables := map[string]interface{}{"a": 2, "b": 2.51, "c": true}
	expr, err := parser.ParseExpr(`b-a>0.7 || c`)
	if err != nil {
		log.Print(err)
	}
	ast.Print(nil, expr)
	res, err := Eval(expr, variables)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(res)
	}
}

func Eval(expr ast.Expr, variables map[string]interface{}) (interface{}, error) {
	rtexpr := reflect.TypeOf(expr)
	switch rtexpr {
	case reflect.TypeOf((*ast.BinaryExpr)(nil)):
		return evalBinary(expr.(*ast.BinaryExpr), variables)
	case reflect.TypeOf((*ast.Ident)(nil)):
		return evalIdent(expr.(*ast.Ident), variables)
	case reflect.TypeOf((*ast.BasicLit)(nil)):
		return evalBasicLit(expr.(*ast.BasicLit), variables)
	}
	return nil, fmt.Errorf("unsupported exprtype(%v)", rtexpr)
}
