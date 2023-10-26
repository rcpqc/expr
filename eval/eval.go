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

func evalUnknown(expr ast.Expr, variables Variables) (any, error) {
	return nil, errs.Newf(expr, "unsupported expression type(%v)", reflect.TypeOf(expr))
}

// evaluator replace eval function to reduce stack depth
func evaluator(expr ast.Expr) func(expr ast.Expr, variables Variables) (any, error) {
	switch expr.(type) {
	case *Constant:
		return evalConstant
	case *ast.Ident:
		return evalIdent
	case *ast.BinaryExpr:
		return evalBinary
	case *ast.SelectorExpr:
		return evalSelector
	case *ast.ParenExpr:
		return evalParen
	case *ast.CallExpr:
		return evalCall
	case *ast.UnaryExpr:
		return evalUnary
	case *ast.IndexExpr:
		return evalIndex
	case *ast.SliceExpr:
		return evalSlice
	case *ast.BasicLit:
		return evalBasicLit
	case *ast.CompositeLit:
		return evalCompositeLit
	default:
		return evalUnknown
	}
}

func eval(expr ast.Expr, variables Variables) (any, error) {
	return evaluator(expr)(expr, variables)
}

func evaltype(expr ast.Expr, variables Variables, t reflect.Type) (any, error) {
	if t == nil {
		t = types.Any
	}
	val, err := evaluator(expr)(expr, variables)
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
