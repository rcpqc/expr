package errs

import (
	"fmt"
	"go/ast"
	"reflect"
)

var names = map[reflect.Type]string{}

func init() {
	names[reflect.TypeOf((*ast.BadExpr)(nil))] = "bad"
	names[reflect.TypeOf((*ast.BinaryExpr)(nil))] = "binary"
	names[reflect.TypeOf((*ast.Ident)(nil))] = "ident"
	names[reflect.TypeOf((*ast.BasicLit)(nil))] = "basic_lit"
	names[reflect.TypeOf((*ast.UnaryExpr)(nil))] = "unary"
	names[reflect.TypeOf((*ast.CallExpr)(nil))] = "call"
	names[reflect.TypeOf((*ast.ParenExpr)(nil))] = "paren"
	names[reflect.TypeOf((*ast.SelectorExpr)(nil))] = "selector"
	names[reflect.TypeOf((*ast.SliceExpr)(nil))] = "slice"
	names[reflect.TypeOf((*ast.IndexExpr)(nil))] = "index"
}

// New new an error with expression information
func New(expr ast.Expr, err error) error {
	if err == nil {
		return nil
	}
	if ierr, ok := err.(*Error); ok && ierr.Expr != "" {
		return err
	}
	if name, ok := names[reflect.TypeOf(expr)]; ok {
		return &Error{Message: err.Error(), Pos: int(expr.Pos()), End: int(expr.End()), Expr: name}
	}
	return &Error{Message: err.Error()}
}

// Newf new error
func Newf(expr ast.Expr, format string, args ...any) error {
	return New(expr, fmt.Errorf(format, args...))
}

// Error error information
type Error struct {
	Expr    string
	Pos     int
	End     int
	Message string
}

// Error error interface
func (o *Error) Error() string {
	if o.Expr == "" {
		return o.Message
	}
	return fmt.Sprintf("%s(%d:%d) %s", o.Expr, o.Pos, o.End-1, o.Message)
}
