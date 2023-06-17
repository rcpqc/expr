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

type binaryKind func(x, y interface{}) interface{}
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
	binaryADD[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) + b2i(y.(bool)) }
	binaryADD[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) + y.(int64) }
	binaryADD[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) + y.(float64) }
	binaryADD[IB] = func(x, y interface{}) interface{} { return x.(int64) + b2i(y.(bool)) }
	binaryADD[II] = func(x, y interface{}) interface{} { return x.(int64) + y.(int64) }
	binaryADD[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) + y.(float64) }
	binaryADD[FB] = func(x, y interface{}) interface{} { return x.(float64) + b2f(y.(bool)) }
	binaryADD[FI] = func(x, y interface{}) interface{} { return x.(float64) + float64(y.(int64)) }
	binaryADD[FF] = func(x, y interface{}) interface{} { return x.(float64) + y.(float64) }
	binaryADD[SS] = func(x, y interface{}) interface{} { return x.(string) + y.(string) }
	binaryTokens[token.ADD] = binaryADD

	// SUB
	binarySUB[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) - b2i(y.(bool)) }
	binarySUB[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) - y.(int64) }
	binarySUB[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) - y.(float64) }
	binarySUB[IB] = func(x, y interface{}) interface{} { return x.(int64) - b2i(y.(bool)) }
	binarySUB[II] = func(x, y interface{}) interface{} { return x.(int64) - y.(int64) }
	binarySUB[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) - y.(float64) }
	binarySUB[FB] = func(x, y interface{}) interface{} { return x.(float64) - b2f(y.(bool)) }
	binarySUB[FI] = func(x, y interface{}) interface{} { return x.(float64) - float64(y.(int64)) }
	binarySUB[FF] = func(x, y interface{}) interface{} { return x.(float64) - y.(float64) }
	binaryTokens[token.SUB] = binarySUB

	// MUL
	binaryMUL[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) * b2i(y.(bool)) }
	binaryMUL[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) * y.(int64) }
	binaryMUL[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) * y.(float64) }
	binaryMUL[IB] = func(x, y interface{}) interface{} { return x.(int64) * b2i(y.(bool)) }
	binaryMUL[II] = func(x, y interface{}) interface{} { return x.(int64) * y.(int64) }
	binaryMUL[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) * y.(float64) }
	binaryMUL[FB] = func(x, y interface{}) interface{} { return x.(float64) * b2f(y.(bool)) }
	binaryMUL[FI] = func(x, y interface{}) interface{} { return x.(float64) * float64(y.(int64)) }
	binaryMUL[FF] = func(x, y interface{}) interface{} { return x.(float64) * y.(float64) }
	binaryTokens[token.MUL] = binaryMUL

	// QUO
	binaryQUO[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) / y.(int64) }
	binaryQUO[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) / y.(float64) }
	binaryQUO[II] = func(x, y interface{}) interface{} { return x.(int64) / y.(int64) }
	binaryQUO[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) / y.(float64) }
	binaryQUO[FI] = func(x, y interface{}) interface{} { return x.(float64) / float64(y.(int64)) }
	binaryQUO[FF] = func(x, y interface{}) interface{} { return x.(float64) / y.(float64) }
	binaryTokens[token.QUO] = binaryQUO

	// REM
	binaryREM[II] = func(x, y interface{}) interface{} { return x.(int64) % y.(int64) }
	binaryTokens[token.REM] = binaryREM

	// AND
	binaryAND[II] = func(x, y interface{}) interface{} { return x.(int64) & y.(int64) }
	binaryTokens[token.AND] = binaryAND

	// OR
	binaryOR[II] = func(x, y interface{}) interface{} { return x.(int64) | y.(int64) }
	binaryTokens[token.OR] = binaryOR

	// XOR
	binaryXOR[II] = func(x, y interface{}) interface{} { return x.(int64) ^ y.(int64) }
	binaryTokens[token.XOR] = binaryXOR

	// LAND
	binaryLAND[BB] = func(x, y interface{}) interface{} { return x.(bool) && y.(bool) }
	binaryTokens[token.LAND] = binaryLAND

	// LOR
	binaryLOR[BB] = func(x, y interface{}) interface{} { return x.(bool) || y.(bool) }
	binaryTokens[token.LOR] = binaryLOR

	// EQL
	binaryEQL[BB] = func(x, y interface{}) interface{} { return x.(bool) == y.(bool) }
	binaryEQL[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) == y.(int64) }
	binaryEQL[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) == y.(float64) }
	binaryEQL[IB] = func(x, y interface{}) interface{} { return x.(int64) == b2i(y.(bool)) }
	binaryEQL[II] = func(x, y interface{}) interface{} { return x.(int64) == y.(int64) }
	binaryEQL[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) == y.(float64) }
	binaryEQL[FB] = func(x, y interface{}) interface{} { return x.(float64) == b2f(y.(bool)) }
	binaryEQL[FI] = func(x, y interface{}) interface{} { return x.(float64) == float64(y.(int64)) }
	binaryEQL[FF] = func(x, y interface{}) interface{} { return x.(float64) == y.(float64) }
	binaryEQL[SS] = func(x, y interface{}) interface{} { return x.(string) == y.(string) }
	binaryTokens[token.EQL] = binaryEQL

	// NEQ
	binaryNEQ[BB] = func(x, y interface{}) interface{} { return x.(bool) != y.(bool) }
	binaryNEQ[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) != y.(int64) }
	binaryNEQ[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) != y.(float64) }
	binaryNEQ[IB] = func(x, y interface{}) interface{} { return x.(int64) != b2i(y.(bool)) }
	binaryNEQ[II] = func(x, y interface{}) interface{} { return x.(int64) != y.(int64) }
	binaryNEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) != y.(float64) }
	binaryNEQ[FB] = func(x, y interface{}) interface{} { return x.(float64) != b2f(y.(bool)) }
	binaryNEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) != float64(y.(int64)) }
	binaryNEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) != y.(float64) }
	binaryNEQ[SS] = func(x, y interface{}) interface{} { return x.(string) != y.(string) }
	binaryTokens[token.NEQ] = binaryNEQ

	// LSS
	binaryLSS[II] = func(x, y interface{}) interface{} { return x.(int64) < y.(int64) }
	binaryLSS[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) < y.(float64) }
	binaryLSS[FI] = func(x, y interface{}) interface{} { return x.(float64) < float64(y.(int64)) }
	binaryLSS[FF] = func(x, y interface{}) interface{} { return x.(float64) < y.(float64) }
	binaryLSS[SS] = func(x, y interface{}) interface{} { return x.(string) < y.(string) }
	binaryTokens[token.LSS] = binaryLSS

	// GTR
	binaryGTR[II] = func(x, y interface{}) interface{} { return x.(int64) > y.(int64) }
	binaryGTR[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) > y.(float64) }
	binaryGTR[FI] = func(x, y interface{}) interface{} { return x.(float64) > float64(y.(int64)) }
	binaryGTR[FF] = func(x, y interface{}) interface{} { return x.(float64) > y.(float64) }
	binaryGTR[SS] = func(x, y interface{}) interface{} { return x.(string) > y.(string) }
	binaryTokens[token.GTR] = binaryGTR

	// LEQ
	binaryLEQ[II] = func(x, y interface{}) interface{} { return x.(int64) <= y.(int64) }
	binaryLEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) <= y.(float64) }
	binaryLEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) <= float64(y.(int64)) }
	binaryLEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) <= y.(float64) }
	binaryLEQ[SS] = func(x, y interface{}) interface{} { return x.(string) <= y.(string) }
	binaryTokens[token.LEQ] = binaryLEQ

	// GEQ
	binaryGEQ[II] = func(x, y interface{}) interface{} { return x.(int64) >= y.(int64) }
	binaryGEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) >= y.(float64) }
	binaryGEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) >= float64(y.(int64)) }
	binaryGEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) >= y.(float64) }
	binaryGEQ[SS] = func(x, y interface{}) interface{} { return x.(string) >= y.(string) }
	binaryTokens[token.GEQ] = binaryGEQ

	// SHL
	binarySHL[II] = func(x, y interface{}) interface{} { return x.(int64) << y.(int64) }
	binaryTokens[token.SHL] = binarySHL

	// SHR
	binarySHR[II] = func(x, y interface{}) interface{} { return x.(int64) >> y.(int64) }
	binaryTokens[token.SHR] = binarySHR
}

func evalBinary(binary *ast.BinaryExpr, variables Variables) (interface{}, error) {
	x, err := eval(binary.X, variables)
	if err != nil {
		return nil, err
	}
	x, tx, kx := types.Convert(x)
	if binary.Op == token.LAND && kx == reflect.Bool && !x.(bool) {
		return false, nil
	}
	if binary.Op == token.LOR && kx == reflect.Bool && x.(bool) {
		return true, nil
	}
	y, err := eval(binary.Y, variables)
	if err != nil {
		return nil, err
	}
	y, ty, ky := types.Convert(y)
	if binary.Op == token.QUO && ky == reflect.Int64 && y.(int64) == 0 {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	if binary.Op == token.REM && ky == reflect.Int64 && y.(int64) == 0 {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	handler := binaryTokens[binary.Op][kx*types.MaxKinds+ky]
	if handler == nil {
		return nil, errs.Newf(binary, "illegal expr(%v%s%v)", tx, binary.Op, ty)
	}
	return handler(x, y), nil
}
