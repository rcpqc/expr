package eval

import (
	"go/ast"
	"go/token"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

type unaryKind func(x types.Value) interface{}
type unaryToken [types.MaxKinds]unaryKind

var unaryTokens [MAX_TOKEN]unaryToken

var unaryTokenNOT unaryToken
var unaryTokenSUB unaryToken

func init() {
	// NOT
	unaryTokenNOT[reflect.Bool] = func(x types.Value) interface{} { return !x.B }
	unaryTokens[token.NOT] = unaryTokenNOT

	// SUB
	unaryTokenSUB[reflect.Int64] = func(x types.Value) interface{} { return -x.I }
	unaryTokenSUB[reflect.Float64] = func(x types.Value) interface{} { return -x.F }
	unaryTokens[token.SUB] = unaryTokenSUB
}

func evalUnary(unary *ast.UnaryExpr, variables Variables) (interface{}, error) {
	x, err := eval(unary.X, variables)
	if err != nil {
		return nil, err
	}
	var xvalue types.Value
	kx := types.Convert(x, &xvalue)
	handler := unaryTokens[unary.Op][kx]
	if handler == nil {
		return nil, errs.Newf(unary, "illegal expr(%s%v)", unary.Op, reflect.TypeOf(x))
	}
	return handler(xvalue), nil
}
