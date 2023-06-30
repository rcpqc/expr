package eval

import (
	"go/ast"
	"go/token"
	"reflect"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

const (
	BB = types.MaxKinds*reflect.Bool + reflect.Bool
	BI = types.MaxKinds*reflect.Bool + reflect.Int64
	BF = types.MaxKinds*reflect.Bool + reflect.Float64
	BS = types.MaxKinds*reflect.Bool + reflect.String

	IB = types.MaxKinds*reflect.Int64 + reflect.Bool
	II = types.MaxKinds*reflect.Int64 + reflect.Int64
	IF = types.MaxKinds*reflect.Int64 + reflect.Float64
	IS = types.MaxKinds*reflect.Int64 + reflect.String

	FB = types.MaxKinds*reflect.Float64 + reflect.Bool
	FI = types.MaxKinds*reflect.Float64 + reflect.Int64
	FF = types.MaxKinds*reflect.Float64 + reflect.Float64
	FS = types.MaxKinds*reflect.Float64 + reflect.String

	SB = types.MaxKinds*reflect.String + reflect.Bool
	SI = types.MaxKinds*reflect.String + reflect.Int64
	SF = types.MaxKinds*reflect.String + reflect.Float64
	SS = types.MaxKinds*reflect.String + reflect.String

	MAX_TOKEN = 96
)

type binaryKind func(x, y types.Value) interface{}
type binaryToken [types.MaxKinds * types.MaxKinds]binaryKind

var binaryTokens [MAX_TOKEN]binaryToken

var binaryADD binaryToken
var binarySUB binaryToken
var binaryMUL binaryToken
var binaryQUO binaryToken
var binaryREM binaryToken
var binaryAND binaryToken
var binaryOR binaryToken
var binaryXOR binaryToken
var binaryLAND binaryToken
var binaryLOR binaryToken
var binaryEQL binaryToken
var binaryNEQ binaryToken
var binaryLSS binaryToken
var binaryGTR binaryToken
var binaryLEQ binaryToken
var binaryGEQ binaryToken
var binarySHL binaryToken
var binarySHR binaryToken

func b2i(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func b2f(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func init() {
	// ADD
	binaryADD[BB] = func(x, y types.Value) interface{} { return b2i(x.B) + b2i(y.B) }
	binaryADD[BI] = func(x, y types.Value) interface{} { return b2i(x.B) + y.I }
	binaryADD[BF] = func(x, y types.Value) interface{} { return b2f(x.B) + y.F }
	binaryADD[IB] = func(x, y types.Value) interface{} { return x.I + b2i(y.B) }
	binaryADD[II] = func(x, y types.Value) interface{} { return x.I + y.I }
	binaryADD[IF] = func(x, y types.Value) interface{} { return float64(x.I) + y.F }
	binaryADD[FB] = func(x, y types.Value) interface{} { return x.F + b2f(y.B) }
	binaryADD[FI] = func(x, y types.Value) interface{} { return x.F + float64(y.I) }
	binaryADD[FF] = func(x, y types.Value) interface{} { return x.F + y.F }
	binaryADD[SS] = func(x, y types.Value) interface{} { return x.S + y.S }
	binaryTokens[token.ADD] = binaryADD

	// SUB
	binarySUB[BB] = func(x, y types.Value) interface{} { return b2i(x.B) - b2i(y.B) }
	binarySUB[BI] = func(x, y types.Value) interface{} { return b2i(x.B) - y.I }
	binarySUB[BF] = func(x, y types.Value) interface{} { return b2f(x.B) - y.F }
	binarySUB[IB] = func(x, y types.Value) interface{} { return x.I - b2i(y.B) }
	binarySUB[II] = func(x, y types.Value) interface{} { return x.I - y.I }
	binarySUB[IF] = func(x, y types.Value) interface{} { return float64(x.I) - y.F }
	binarySUB[FB] = func(x, y types.Value) interface{} { return x.F - b2f(y.B) }
	binarySUB[FI] = func(x, y types.Value) interface{} { return x.F - float64(y.I) }
	binarySUB[FF] = func(x, y types.Value) interface{} { return x.F - y.F }
	binaryTokens[token.SUB] = binarySUB

	// MUL
	binaryMUL[BB] = func(x, y types.Value) interface{} { return b2i(x.B) * b2i(y.B) }
	binaryMUL[BI] = func(x, y types.Value) interface{} { return b2i(x.B) * y.I }
	binaryMUL[BF] = func(x, y types.Value) interface{} { return b2f(x.B) * y.F }
	binaryMUL[IB] = func(x, y types.Value) interface{} { return x.I * b2i(y.B) }
	binaryMUL[II] = func(x, y types.Value) interface{} { return x.I * y.I }
	binaryMUL[IF] = func(x, y types.Value) interface{} { return float64(x.I) * y.F }
	binaryMUL[FB] = func(x, y types.Value) interface{} { return x.F * b2f(y.B) }
	binaryMUL[FI] = func(x, y types.Value) interface{} { return x.F * float64(y.I) }
	binaryMUL[FF] = func(x, y types.Value) interface{} { return x.F * y.F }
	binaryTokens[token.MUL] = binaryMUL

	// QUO
	binaryQUO[BI] = func(x, y types.Value) interface{} { return b2i(x.B) / y.I }
	binaryQUO[BF] = func(x, y types.Value) interface{} { return b2f(x.B) / y.F }
	binaryQUO[II] = func(x, y types.Value) interface{} { return x.I / y.I }
	binaryQUO[IF] = func(x, y types.Value) interface{} { return float64(x.I) / y.F }
	binaryQUO[FI] = func(x, y types.Value) interface{} { return x.F / float64(y.I) }
	binaryQUO[FF] = func(x, y types.Value) interface{} { return x.F / y.F }
	binaryTokens[token.QUO] = binaryQUO

	// REM
	binaryREM[II] = func(x, y types.Value) interface{} { return x.I % y.I }
	binaryTokens[token.REM] = binaryREM

	// AND
	binaryAND[II] = func(x, y types.Value) interface{} { return x.I & y.I }
	binaryTokens[token.AND] = binaryAND

	// OR
	binaryOR[II] = func(x, y types.Value) interface{} { return x.I | y.I }
	binaryTokens[token.OR] = binaryOR

	// XOR
	binaryXOR[II] = func(x, y types.Value) interface{} { return x.I ^ y.I }
	binaryTokens[token.XOR] = binaryXOR

	// LAND
	binaryLAND[BB] = func(x, y types.Value) interface{} { return x.B && y.B }
	binaryTokens[token.LAND] = binaryLAND

	// LOR
	binaryLOR[BB] = func(x, y types.Value) interface{} { return x.B || y.B }
	binaryTokens[token.LOR] = binaryLOR

	// EQL
	binaryEQL[BB] = func(x, y types.Value) interface{} { return x.B == y.B }
	binaryEQL[BI] = func(x, y types.Value) interface{} { return b2i(x.B) == y.I }
	binaryEQL[BF] = func(x, y types.Value) interface{} { return b2f(x.B) == y.F }
	binaryEQL[IB] = func(x, y types.Value) interface{} { return x.I == b2i(y.B) }
	binaryEQL[II] = func(x, y types.Value) interface{} { return x.I == y.I }
	binaryEQL[IF] = func(x, y types.Value) interface{} { return float64(x.I) == y.F }
	binaryEQL[FB] = func(x, y types.Value) interface{} { return x.F == b2f(y.B) }
	binaryEQL[FI] = func(x, y types.Value) interface{} { return x.F == float64(y.I) }
	binaryEQL[FF] = func(x, y types.Value) interface{} { return x.F == y.F }
	binaryEQL[SS] = func(x, y types.Value) interface{} { return x.S == y.S }
	binaryTokens[token.EQL] = binaryEQL

	// NEQ
	binaryNEQ[BB] = func(x, y types.Value) interface{} { return x.B != y.B }
	binaryNEQ[BI] = func(x, y types.Value) interface{} { return b2i(x.B) != y.I }
	binaryNEQ[BF] = func(x, y types.Value) interface{} { return b2f(x.B) != y.F }
	binaryNEQ[IB] = func(x, y types.Value) interface{} { return x.I != b2i(y.B) }
	binaryNEQ[II] = func(x, y types.Value) interface{} { return x.I != y.I }
	binaryNEQ[IF] = func(x, y types.Value) interface{} { return float64(x.I) != y.F }
	binaryNEQ[FB] = func(x, y types.Value) interface{} { return x.F != b2f(y.B) }
	binaryNEQ[FI] = func(x, y types.Value) interface{} { return x.F != float64(y.I) }
	binaryNEQ[FF] = func(x, y types.Value) interface{} { return x.F != y.F }
	binaryNEQ[SS] = func(x, y types.Value) interface{} { return x.S != y.S }
	binaryTokens[token.NEQ] = binaryNEQ

	// LSS
	binaryLSS[II] = func(x, y types.Value) interface{} { return x.I < y.I }
	binaryLSS[IF] = func(x, y types.Value) interface{} { return float64(x.I) < y.F }
	binaryLSS[FI] = func(x, y types.Value) interface{} { return x.F < float64(y.I) }
	binaryLSS[FF] = func(x, y types.Value) interface{} { return x.F < y.F }
	binaryLSS[SS] = func(x, y types.Value) interface{} { return x.S < y.S }
	binaryTokens[token.LSS] = binaryLSS

	// GTR
	binaryGTR[II] = func(x, y types.Value) interface{} { return x.I > y.I }
	binaryGTR[IF] = func(x, y types.Value) interface{} { return float64(x.I) > y.F }
	binaryGTR[FI] = func(x, y types.Value) interface{} { return x.F > float64(y.I) }
	binaryGTR[FF] = func(x, y types.Value) interface{} { return x.F > y.F }
	binaryGTR[SS] = func(x, y types.Value) interface{} { return x.S > y.S }
	binaryTokens[token.GTR] = binaryGTR

	// LEQ
	binaryLEQ[II] = func(x, y types.Value) interface{} { return x.I <= y.I }
	binaryLEQ[IF] = func(x, y types.Value) interface{} { return float64(x.I) <= y.F }
	binaryLEQ[FI] = func(x, y types.Value) interface{} { return x.F <= float64(y.I) }
	binaryLEQ[FF] = func(x, y types.Value) interface{} { return x.F <= y.F }
	binaryLEQ[SS] = func(x, y types.Value) interface{} { return x.S <= y.S }
	binaryTokens[token.LEQ] = binaryLEQ

	// GEQ
	binaryGEQ[II] = func(x, y types.Value) interface{} { return x.I >= y.I }
	binaryGEQ[IF] = func(x, y types.Value) interface{} { return float64(x.I) >= y.F }
	binaryGEQ[FI] = func(x, y types.Value) interface{} { return x.F >= float64(y.I) }
	binaryGEQ[FF] = func(x, y types.Value) interface{} { return x.F >= y.F }
	binaryGEQ[SS] = func(x, y types.Value) interface{} { return x.S >= y.S }
	binaryTokens[token.GEQ] = binaryGEQ

	// SHL
	binarySHL[II] = func(x, y types.Value) interface{} { return x.I << y.I }
	binaryTokens[token.SHL] = binarySHL

	// SHR
	binarySHR[II] = func(x, y types.Value) interface{} { return x.I >> y.I }
	binaryTokens[token.SHR] = binarySHR
}

func evalBinary(binary *ast.BinaryExpr, variables Variables) (interface{}, error) {
	x, err := eval(binary.X, variables)
	if err != nil {
		return nil, err
	}
	var xvalue, yvalue types.Value
	kx := types.Convert(x, &xvalue)
	if binary.Op == token.LAND && kx == reflect.Bool && !xvalue.B {
		return false, nil
	}
	if binary.Op == token.LOR && kx == reflect.Bool && xvalue.B {
		return true, nil
	}
	y, err := eval(binary.Y, variables)
	if err != nil {
		return nil, err
	}
	ky := types.Convert(y, &yvalue)
	if binary.Op == token.QUO && ky == reflect.Int64 && yvalue.I == 0 {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	if binary.Op == token.REM && ky == reflect.Int64 && yvalue.I == 0 {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	handler := binaryTokens[binary.Op][kx*types.MaxKinds+ky]
	if handler == nil {
		return nil, errs.Newf(binary, "illegal expr(%v%s%v)", reflect.TypeOf(x), binary.Op, reflect.TypeOf(y))
	}
	return handler(xvalue, yvalue), nil
}
