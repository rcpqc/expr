package eval

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/rcpqc/expr/types"
)

type unaryKind func(x interface{}) interface{}
type unaryToken [types.MaxKinds]unaryKind

var unaryTokens [MAX_TOKEN]unaryToken

var unaryTokenNOT unaryToken
var unaryTokenSUB unaryToken

func init() {
	// NOT
	unaryTokenNOT[reflect.Bool] = func(x interface{}) interface{} { return !x.(bool) }
	unaryTokens[token.NOT] = unaryTokenNOT

	// SUB
	unaryTokenSUB[reflect.Int64] = func(x interface{}) interface{} { return -x.(int64) }
	unaryTokenSUB[reflect.Float64] = func(x interface{}) interface{} { return -x.(float64) }
	unaryTokens[token.SUB] = unaryTokenSUB
}

func evalUnary(unary *ast.UnaryExpr, variables Variables) (interface{}, error) {
	x, err := eval(unary.X, variables)
	if err != nil {
		return nil, err
	}
	x, kx := types.Convert(x)
	handler := unaryTokens[unary.Op][kx]
	if handler == nil {
		return nil, fmt.Errorf("[unary] illegal expr (%s %v)", unary.Op.String(), kx)
	}
	return handler(x), nil
}
