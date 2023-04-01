package eval

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

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

type binaryKind func(x, y interface{}) (interface{}, error)
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
	binaryADD[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) + b2i(y.(bool)), nil }
	binaryADD[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) + y.(int64), nil }
	binaryADD[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) + y.(float64), nil }
	binaryADD[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) + b2i(y.(bool)), nil }
	binaryADD[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) + y.(int64), nil }
	binaryADD[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) + y.(float64), nil }
	binaryADD[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) + b2f(y.(bool)), nil }
	binaryADD[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) + float64(y.(int64)), nil }
	binaryADD[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) + y.(float64), nil }
	binaryADD[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) + y.(string), nil }
	binaryTokens[token.ADD] = binaryADD

	// SUB
	binarySUB[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) - b2i(y.(bool)), nil }
	binarySUB[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) - y.(int64), nil }
	binarySUB[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) - y.(float64), nil }
	binarySUB[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) - b2i(y.(bool)), nil }
	binarySUB[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) - y.(int64), nil }
	binarySUB[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) - y.(float64), nil }
	binarySUB[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) - b2f(y.(bool)), nil }
	binarySUB[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) - float64(y.(int64)), nil }
	binarySUB[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) - y.(float64), nil }
	binaryTokens[token.SUB] = binarySUB

	// MUL
	binaryMUL[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) * b2i(y.(bool)), nil }
	binaryMUL[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) * y.(int64), nil }
	binaryMUL[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) * y.(float64), nil }
	binaryMUL[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) * b2i(y.(bool)), nil }
	binaryMUL[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) * y.(int64), nil }
	binaryMUL[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) * y.(float64), nil }
	binaryMUL[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) * b2f(y.(bool)), nil }
	binaryMUL[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) * float64(y.(int64)), nil }
	binaryMUL[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) * y.(float64), nil }
	binaryTokens[token.MUL] = binaryMUL

	// QUO
	binaryQUO[BI] = func(x, y interface{}) (interface{}, error) {
		if y.(int64) == 0 {
			return nil, fmt.Errorf("integer divide by zero")
		}
		return b2i(x.(bool)) / y.(int64), nil
	}
	binaryQUO[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) / y.(float64), nil }
	binaryQUO[II] = func(x, y interface{}) (interface{}, error) {
		if y.(int64) == 0 {
			return nil, fmt.Errorf("integer divide by zero")
		}
		return x.(int64) / y.(int64), nil
	}
	binaryQUO[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) / y.(float64), nil }
	binaryQUO[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) / float64(y.(int64)), nil }
	binaryQUO[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) / y.(float64), nil }
	binaryTokens[token.QUO] = binaryQUO

	// REM
	binaryREM[II] = func(x, y interface{}) (interface{}, error) {
		if y.(int64) == 0 {
			return nil, fmt.Errorf("integer divide by zero")
		}
		return x.(int64) % y.(int64), nil
	}
	binaryTokens[token.REM] = binaryREM

	// AND
	binaryAND[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) & b2i(y.(bool)), nil }
	binaryAND[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) & y.(int64), nil }
	binaryAND[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) & b2i(y.(bool)), nil }
	binaryAND[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) & y.(int64), nil }
	binaryTokens[token.AND] = binaryAND

	// OR
	binaryOR[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) | b2i(y.(bool)), nil }
	binaryOR[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) | y.(int64), nil }
	binaryOR[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) | b2i(y.(bool)), nil }
	binaryOR[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) | y.(int64), nil }
	binaryTokens[token.OR] = binaryOR

	// XOR
	binaryXOR[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) ^ b2i(y.(bool)), nil }
	binaryXOR[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) ^ y.(int64), nil }
	binaryXOR[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) ^ b2i(y.(bool)), nil }
	binaryXOR[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) ^ y.(int64), nil }
	binaryTokens[token.XOR] = binaryXOR

	// LAND
	binaryLAND[BB] = func(x, y interface{}) (interface{}, error) { return x.(bool) && y.(bool), nil }
	binaryLAND[BI] = func(x, y interface{}) (interface{}, error) { return x.(bool) && y.(int64) != 0, nil }
	binaryLAND[BF] = func(x, y interface{}) (interface{}, error) { return x.(bool) && y.(float64) != 0, nil }
	binaryLAND[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 && y.(bool), nil }
	binaryLAND[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 && y.(int64) != 0, nil }
	binaryLAND[IF] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 && y.(float64) != 0, nil }
	binaryLAND[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 && y.(bool), nil }
	binaryLAND[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 && y.(int64) != 0, nil }
	binaryLAND[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 && y.(float64) != 0, nil }
	binaryTokens[token.LAND] = binaryLAND

	// LOR
	binaryLOR[BB] = func(x, y interface{}) (interface{}, error) { return x.(bool) || y.(bool), nil }
	binaryLOR[BI] = func(x, y interface{}) (interface{}, error) { return x.(bool) || y.(int64) != 0, nil }
	binaryLOR[BF] = func(x, y interface{}) (interface{}, error) { return x.(bool) || y.(float64) != 0, nil }
	binaryLOR[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 || y.(bool), nil }
	binaryLOR[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 || y.(int64) != 0, nil }
	binaryLOR[IF] = func(x, y interface{}) (interface{}, error) { return x.(int64) != 0 || y.(float64) != 0, nil }
	binaryLOR[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 || y.(bool), nil }
	binaryLOR[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 || y.(int64) != 0, nil }
	binaryLOR[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) != 0 || y.(float64) != 0, nil }
	binaryTokens[token.LOR] = binaryLOR

	// EQL
	binaryEQL[BB] = func(x, y interface{}) (interface{}, error) { return x.(bool) == y.(bool), nil }
	binaryEQL[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) == y.(int64), nil }
	binaryEQL[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) == y.(float64), nil }
	binaryEQL[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) == b2i(y.(bool)), nil }
	binaryEQL[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) == y.(int64), nil }
	binaryEQL[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) == y.(float64), nil }
	binaryEQL[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) == b2f(y.(bool)), nil }
	binaryEQL[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) == float64(y.(int64)), nil }
	binaryEQL[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) == y.(float64), nil }
	binaryEQL[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) == y.(string), nil }
	binaryTokens[token.EQL] = binaryEQL

	// NEQ
	binaryNEQ[BB] = func(x, y interface{}) (interface{}, error) { return x.(bool) != y.(bool), nil }
	binaryNEQ[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) != y.(int64), nil }
	binaryNEQ[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) != y.(float64), nil }
	binaryNEQ[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) != b2i(y.(bool)), nil }
	binaryNEQ[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) != y.(int64), nil }
	binaryNEQ[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) != y.(float64), nil }
	binaryNEQ[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) != b2f(y.(bool)), nil }
	binaryNEQ[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) != float64(y.(int64)), nil }
	binaryNEQ[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) != y.(float64), nil }
	binaryNEQ[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) != y.(string), nil }
	binaryTokens[token.NEQ] = binaryNEQ

	// LSS
	binaryLSS[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) < b2i(y.(bool)), nil }
	binaryLSS[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) < y.(int64), nil }
	binaryLSS[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) < y.(float64), nil }
	binaryLSS[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) < b2i(y.(bool)), nil }
	binaryLSS[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) < y.(int64), nil }
	binaryLSS[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) < y.(float64), nil }
	binaryLSS[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) < b2f(y.(bool)), nil }
	binaryLSS[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) < float64(y.(int64)), nil }
	binaryLSS[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) < y.(float64), nil }
	binaryLSS[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) < y.(string), nil }
	binaryTokens[token.LSS] = binaryLSS

	// GTR
	binaryGTR[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) > b2i(y.(bool)), nil }
	binaryGTR[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) > y.(int64), nil }
	binaryGTR[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) > y.(float64), nil }
	binaryGTR[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) > b2i(y.(bool)), nil }
	binaryGTR[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) > y.(int64), nil }
	binaryGTR[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) > y.(float64), nil }
	binaryGTR[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) > b2f(y.(bool)), nil }
	binaryGTR[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) > float64(y.(int64)), nil }
	binaryGTR[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) > y.(float64), nil }
	binaryGTR[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) > y.(string), nil }
	binaryTokens[token.GTR] = binaryGTR

	// LEQ
	binaryLEQ[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) <= b2i(y.(bool)), nil }
	binaryLEQ[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) <= y.(int64), nil }
	binaryLEQ[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) <= y.(float64), nil }
	binaryLEQ[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) <= b2i(y.(bool)), nil }
	binaryLEQ[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) <= y.(int64), nil }
	binaryLEQ[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) <= y.(float64), nil }
	binaryLEQ[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) <= b2f(y.(bool)), nil }
	binaryLEQ[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) <= float64(y.(int64)), nil }
	binaryLEQ[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) <= y.(float64), nil }
	binaryLEQ[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) <= y.(string), nil }
	binaryTokens[token.LEQ] = binaryLEQ

	// GEQ
	binaryGEQ[BB] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) >= b2i(y.(bool)), nil }
	binaryGEQ[BI] = func(x, y interface{}) (interface{}, error) { return b2i(x.(bool)) >= y.(int64), nil }
	binaryGEQ[BF] = func(x, y interface{}) (interface{}, error) { return b2f(x.(bool)) >= y.(float64), nil }
	binaryGEQ[IB] = func(x, y interface{}) (interface{}, error) { return x.(int64) >= b2i(y.(bool)), nil }
	binaryGEQ[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) >= y.(int64), nil }
	binaryGEQ[IF] = func(x, y interface{}) (interface{}, error) { return float64(x.(int64)) >= y.(float64), nil }
	binaryGEQ[FB] = func(x, y interface{}) (interface{}, error) { return x.(float64) >= b2f(y.(bool)), nil }
	binaryGEQ[FI] = func(x, y interface{}) (interface{}, error) { return x.(float64) >= float64(y.(int64)), nil }
	binaryGEQ[FF] = func(x, y interface{}) (interface{}, error) { return x.(float64) >= y.(float64), nil }
	binaryGEQ[SS] = func(x, y interface{}) (interface{}, error) { return x.(string) >= y.(string), nil }
	binaryTokens[token.GEQ] = binaryGEQ

	// SHL
	binarySHL[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) << y.(int64), nil }
	binaryTokens[token.SHL] = binarySHL

	// SHR
	binarySHR[II] = func(x, y interface{}) (interface{}, error) { return x.(int64) >> y.(int64), nil }
	binaryTokens[token.SHR] = binarySHR
}

func evalBinary(binary *ast.BinaryExpr, variables Variables) (interface{}, error) {
	x, err := eval(binary.X, variables)
	if err != nil {
		return nil, err
	}
	y, err := eval(binary.Y, variables)
	if err != nil {
		return nil, err
	}
	x, kx := types.Convert(x)
	y, ky := types.Convert(y)
	handler := binaryTokens[binary.Op][kx*types.MaxKinds+ky]
	if handler == nil {
		return nil, fmt.Errorf("[binary] illegal expr (%v %s %v)", kx, binary.Op.String(), ky)
	}
	return handler(x, y)
}
