package eval

import (
	"go/ast"
	"go/token"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

type unaryKind func(x reflect.Value) any
type unaryToken [types.MaxKinds]unaryKind

var unaryTokens [MAX_TOKEN]unaryToken

var unaryTokenNOT unaryToken
var unaryTokenSUB unaryToken

func init() {
	// NOT
	unaryTokenNOT[reflect.Bool] = func(x reflect.Value) any { return !x.Bool() }
	unaryTokens[token.NOT] = unaryTokenNOT

	// SUB
	unaryTokenSUB[reflect.Int64] = func(x reflect.Value) any { return -x.Int() }
	unaryTokenSUB[reflect.Uint64] = func(x reflect.Value) any { return -int64(x.Uint()) }
	unaryTokenSUB[reflect.Float64] = func(x reflect.Value) any { return -x.Float() }
	unaryTokens[token.SUB] = unaryTokenSUB
}

func evalUnary(expr ast.Expr, variables Variables) (any, error) {
	unary := expr.(*ast.UnaryExpr)
	x, err := evaluator(unary.X)(unary.X, variables)
	if err != nil {
		return nil, err
	}
	xvalue := reflect.ValueOf(x)
	kx := kinds[xvalue.Kind()]
	handler := unaryTokens[unary.Op][kx]
	if handler == nil {
		return nil, errs.Newf(unary, "illegal expr(%s%v)", unary.Op, reflect.TypeOf(x))
	}
	return handler(xvalue), nil
}
