package expr

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

const (
	BB = 32*reflect.Bool + reflect.Bool
	BI = 32*reflect.Bool + reflect.Int64
	BF = 32*reflect.Bool + reflect.Float64
	BS = 32*reflect.Bool + reflect.String

	IB = 32*reflect.Int64 + reflect.Bool
	II = 32*reflect.Int64 + reflect.Int64
	IF = 32*reflect.Int64 + reflect.Float64
	IS = 32*reflect.Int64 + reflect.String

	FB = 32*reflect.Float64 + reflect.Bool
	FI = 32*reflect.Float64 + reflect.Int64
	FF = 32*reflect.Float64 + reflect.Float64
	FS = 32*reflect.Float64 + reflect.String

	SB = 32*reflect.String + reflect.Bool
	SI = 32*reflect.String + reflect.Int64
	SF = 32*reflect.String + reflect.Float64
	SS = 32*reflect.String + reflect.String

	MAX_KIND  = 32
	MAX_TOKEN = 96
)

type binaryKind func(x, y interface{}) interface{}
type binaryToken [MAX_KIND * MAX_KIND]binaryKind

var binaryTokens [MAX_TOKEN]binaryToken

var binaryTokenADD binaryToken
var binaryTokenSUB binaryToken
var binaryTokenMUL binaryToken
var binaryTokenQUO binaryToken
var binaryTokenREM binaryToken
var binaryTokenAND binaryToken
var binaryTokenOR binaryToken
var binaryTokenXOR binaryToken
var binaryTokenLAND binaryToken
var binaryTokenLOR binaryToken
var binaryTokenEQL binaryToken
var binaryTokenNEQ binaryToken
var binaryTokenLSS binaryToken
var binaryTokenGTR binaryToken
var binaryTokenLEQ binaryToken
var binaryTokenGEQ binaryToken
var binaryTokenSHL binaryToken
var binaryTokenSHR binaryToken

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
	binaryTokenADD[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) + b2i(y.(bool)) }
	binaryTokenADD[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) + y.(int64) }
	binaryTokenADD[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) + y.(float64) }
	binaryTokenADD[IB] = func(x, y interface{}) interface{} { return x.(int64) + b2i(y.(bool)) }
	binaryTokenADD[II] = func(x, y interface{}) interface{} { return x.(int64) + y.(int64) }
	binaryTokenADD[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) + y.(float64) }
	binaryTokenADD[FB] = func(x, y interface{}) interface{} { return x.(float64) + b2f(y.(bool)) }
	binaryTokenADD[FI] = func(x, y interface{}) interface{} { return x.(float64) + float64(y.(int64)) }
	binaryTokenADD[FF] = func(x, y interface{}) interface{} { return x.(float64) + y.(float64) }
	binaryTokens[token.ADD] = binaryTokenADD

	// SUB
	binaryTokenSUB[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) - b2i(y.(bool)) }
	binaryTokenSUB[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) - y.(int64) }
	binaryTokenSUB[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) - y.(float64) }
	binaryTokenSUB[IB] = func(x, y interface{}) interface{} { return x.(int64) - b2i(y.(bool)) }
	binaryTokenSUB[II] = func(x, y interface{}) interface{} { return x.(int64) - y.(int64) }
	binaryTokenSUB[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) - y.(float64) }
	binaryTokenSUB[FB] = func(x, y interface{}) interface{} { return x.(float64) - b2f(y.(bool)) }
	binaryTokenSUB[FI] = func(x, y interface{}) interface{} { return x.(float64) - float64(y.(int64)) }
	binaryTokenSUB[FF] = func(x, y interface{}) interface{} { return x.(float64) - y.(float64) }
	binaryTokens[token.SUB] = binaryTokenSUB

	// MUL
	binaryTokenMUL[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) * b2i(y.(bool)) }
	binaryTokenMUL[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) * y.(int64) }
	binaryTokenMUL[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) * y.(float64) }
	binaryTokenMUL[IB] = func(x, y interface{}) interface{} { return x.(int64) * b2i(y.(bool)) }
	binaryTokenMUL[II] = func(x, y interface{}) interface{} { return x.(int64) * y.(int64) }
	binaryTokenMUL[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) * y.(float64) }
	binaryTokenMUL[FB] = func(x, y interface{}) interface{} { return x.(float64) * b2f(y.(bool)) }
	binaryTokenMUL[FI] = func(x, y interface{}) interface{} { return x.(float64) * float64(y.(int64)) }
	binaryTokenMUL[FF] = func(x, y interface{}) interface{} { return x.(float64) * y.(float64) }
	binaryTokens[token.MUL] = binaryTokenMUL

	// QUO
	binaryTokenQUO[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) / b2i(y.(bool)) }
	binaryTokenQUO[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) / y.(int64) }
	binaryTokenQUO[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) / y.(float64) }
	binaryTokenQUO[IB] = func(x, y interface{}) interface{} { return x.(int64) / b2i(y.(bool)) }
	binaryTokenQUO[II] = func(x, y interface{}) interface{} { return x.(int64) / y.(int64) }
	binaryTokenQUO[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) / y.(float64) }
	binaryTokenQUO[FB] = func(x, y interface{}) interface{} { return x.(float64) / b2f(y.(bool)) }
	binaryTokenQUO[FI] = func(x, y interface{}) interface{} { return x.(float64) / float64(y.(int64)) }
	binaryTokenQUO[FF] = func(x, y interface{}) interface{} { return x.(float64) / y.(float64) }
	binaryTokens[token.QUO] = binaryTokenQUO

	// REM
	binaryTokenQUO[II] = func(x, y interface{}) interface{} { return x.(int64) % y.(int64) }
	binaryTokens[token.REM] = binaryTokenREM

	// AND
	binaryTokenAND[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) & b2i(y.(bool)) }
	binaryTokenAND[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) & y.(int64) }
	binaryTokenAND[IB] = func(x, y interface{}) interface{} { return x.(int64) & b2i(y.(bool)) }
	binaryTokenAND[II] = func(x, y interface{}) interface{} { return x.(int64) & y.(int64) }
	binaryTokens[token.AND] = binaryTokenAND

	// OR
	binaryTokenOR[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) | b2i(y.(bool)) }
	binaryTokenOR[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) | y.(int64) }
	binaryTokenOR[IB] = func(x, y interface{}) interface{} { return x.(int64) | b2i(y.(bool)) }
	binaryTokenOR[II] = func(x, y interface{}) interface{} { return x.(int64) | y.(int64) }
	binaryTokens[token.OR] = binaryTokenOR

	// XOR
	binaryTokenXOR[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) ^ b2i(y.(bool)) }
	binaryTokenXOR[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) ^ y.(int64) }
	binaryTokenXOR[IB] = func(x, y interface{}) interface{} { return x.(int64) ^ b2i(y.(bool)) }
	binaryTokenXOR[II] = func(x, y interface{}) interface{} { return x.(int64) ^ y.(int64) }
	binaryTokens[token.XOR] = binaryTokenXOR

	// LAND
	binaryTokenLAND[BB] = func(x, y interface{}) interface{} { return x.(bool) && y.(bool) }
	binaryTokenLAND[BI] = func(x, y interface{}) interface{} { return x.(bool) && y.(int64) != 0 }
	binaryTokenLAND[BF] = func(x, y interface{}) interface{} { return x.(bool) && y.(float64) != 0 }
	binaryTokenLAND[IB] = func(x, y interface{}) interface{} { return x.(int64) != 0 && y.(bool) }
	binaryTokenLAND[II] = func(x, y interface{}) interface{} { return x.(int64) != 0 && y.(int64) != 0 }
	binaryTokenLAND[IF] = func(x, y interface{}) interface{} { return x.(int64) != 0 && y.(float64) != 0 }
	binaryTokenLAND[FB] = func(x, y interface{}) interface{} { return x.(float64) != 0 && y.(bool) }
	binaryTokenLAND[FI] = func(x, y interface{}) interface{} { return x.(float64) != 0 && y.(int64) != 0 }
	binaryTokenLAND[FF] = func(x, y interface{}) interface{} { return x.(float64) != 0 && y.(float64) != 0 }
	binaryTokens[token.LAND] = binaryTokenLAND

	// LOR
	binaryTokenLOR[BB] = func(x, y interface{}) interface{} { return x.(bool) || y.(bool) }
	binaryTokenLOR[BI] = func(x, y interface{}) interface{} { return x.(bool) || y.(int64) != 0 }
	binaryTokenLOR[BF] = func(x, y interface{}) interface{} { return x.(bool) || y.(float64) != 0 }
	binaryTokenLOR[IB] = func(x, y interface{}) interface{} { return x.(int64) != 0 || y.(bool) }
	binaryTokenLOR[II] = func(x, y interface{}) interface{} { return x.(int64) != 0 || y.(int64) != 0 }
	binaryTokenLOR[IF] = func(x, y interface{}) interface{} { return x.(int64) != 0 || y.(float64) != 0 }
	binaryTokenLOR[FB] = func(x, y interface{}) interface{} { return x.(float64) != 0 || y.(bool) }
	binaryTokenLOR[FI] = func(x, y interface{}) interface{} { return x.(float64) != 0 || y.(int64) != 0 }
	binaryTokenLOR[FF] = func(x, y interface{}) interface{} { return x.(float64) != 0 || y.(float64) != 0 }
	binaryTokens[token.LOR] = binaryTokenLOR

	// EQL
	binaryTokenEQL[BB] = func(x, y interface{}) interface{} { return x.(bool) == y.(bool) }
	binaryTokenEQL[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) == y.(int64) }
	binaryTokenEQL[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) == y.(float64) }
	binaryTokenEQL[IB] = func(x, y interface{}) interface{} { return x.(int64) == b2i(y.(bool)) }
	binaryTokenEQL[II] = func(x, y interface{}) interface{} { return x.(int64) == y.(int64) }
	binaryTokenEQL[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) == y.(float64) }
	binaryTokenEQL[FB] = func(x, y interface{}) interface{} { return x.(float64) == b2f(y.(bool)) }
	binaryTokenEQL[FI] = func(x, y interface{}) interface{} { return x.(float64) == float64(y.(int64)) }
	binaryTokenEQL[FF] = func(x, y interface{}) interface{} { return x.(float64) == y.(float64) }
	binaryTokenEQL[SS] = func(x, y interface{}) interface{} { return x.(string) == y.(string) }
	binaryTokens[token.EQL] = binaryTokenEQL

	// NEQ
	binaryTokenNEQ[BB] = func(x, y interface{}) interface{} { return x.(bool) != y.(bool) }
	binaryTokenNEQ[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) != y.(int64) }
	binaryTokenNEQ[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) != y.(float64) }
	binaryTokenNEQ[IB] = func(x, y interface{}) interface{} { return x.(int64) != b2i(y.(bool)) }
	binaryTokenNEQ[II] = func(x, y interface{}) interface{} { return x.(int64) != y.(int64) }
	binaryTokenNEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) != y.(float64) }
	binaryTokenNEQ[FB] = func(x, y interface{}) interface{} { return x.(float64) != b2f(y.(bool)) }
	binaryTokenNEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) != float64(y.(int64)) }
	binaryTokenNEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) != y.(float64) }
	binaryTokenNEQ[SS] = func(x, y interface{}) interface{} { return x.(string) != y.(string) }
	binaryTokens[token.NEQ] = binaryTokenNEQ

	// LSS
	binaryTokenLSS[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) < b2i(y.(bool)) }
	binaryTokenLSS[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) < y.(int64) }
	binaryTokenLSS[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) < y.(float64) }
	binaryTokenLSS[IB] = func(x, y interface{}) interface{} { return x.(int64) < b2i(y.(bool)) }
	binaryTokenLSS[II] = func(x, y interface{}) interface{} { return x.(int64) < y.(int64) }
	binaryTokenLSS[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) < y.(float64) }
	binaryTokenLSS[FB] = func(x, y interface{}) interface{} { return x.(float64) < b2f(y.(bool)) }
	binaryTokenLSS[FI] = func(x, y interface{}) interface{} { return x.(float64) < float64(y.(int64)) }
	binaryTokenLSS[FF] = func(x, y interface{}) interface{} { return x.(float64) < y.(float64) }
	binaryTokenLSS[SS] = func(x, y interface{}) interface{} { return x.(string) < y.(string) }
	binaryTokens[token.LSS] = binaryTokenLSS

	// GTR
	binaryTokenGTR[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) > b2i(y.(bool)) }
	binaryTokenGTR[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) > y.(int64) }
	binaryTokenGTR[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) > y.(float64) }
	binaryTokenGTR[IB] = func(x, y interface{}) interface{} { return x.(int64) > b2i(y.(bool)) }
	binaryTokenGTR[II] = func(x, y interface{}) interface{} { return x.(int64) > y.(int64) }
	binaryTokenGTR[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) > y.(float64) }
	binaryTokenGTR[FB] = func(x, y interface{}) interface{} { return x.(float64) > b2f(y.(bool)) }
	binaryTokenGTR[FI] = func(x, y interface{}) interface{} { return x.(float64) > float64(y.(int64)) }
	binaryTokenGTR[FF] = func(x, y interface{}) interface{} { return x.(float64) > y.(float64) }
	binaryTokenGTR[SS] = func(x, y interface{}) interface{} { return x.(string) > y.(string) }
	binaryTokens[token.GTR] = binaryTokenGTR

	// LEQ
	binaryTokenLEQ[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) <= b2i(y.(bool)) }
	binaryTokenLEQ[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) <= y.(int64) }
	binaryTokenLEQ[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) <= y.(float64) }
	binaryTokenLEQ[IB] = func(x, y interface{}) interface{} { return x.(int64) <= b2i(y.(bool)) }
	binaryTokenLEQ[II] = func(x, y interface{}) interface{} { return x.(int64) <= y.(int64) }
	binaryTokenLEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) <= y.(float64) }
	binaryTokenLEQ[FB] = func(x, y interface{}) interface{} { return x.(float64) <= b2f(y.(bool)) }
	binaryTokenLEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) <= float64(y.(int64)) }
	binaryTokenLEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) <= y.(float64) }
	binaryTokenLEQ[SS] = func(x, y interface{}) interface{} { return x.(string) <= y.(string) }
	binaryTokens[token.LEQ] = binaryTokenLEQ

	// GEQ
	binaryTokenGEQ[BB] = func(x, y interface{}) interface{} { return b2i(x.(bool)) >= b2i(y.(bool)) }
	binaryTokenGEQ[BI] = func(x, y interface{}) interface{} { return b2i(x.(bool)) >= y.(int64) }
	binaryTokenGEQ[BF] = func(x, y interface{}) interface{} { return b2f(x.(bool)) >= y.(float64) }
	binaryTokenGEQ[IB] = func(x, y interface{}) interface{} { return x.(int64) >= b2i(y.(bool)) }
	binaryTokenGEQ[II] = func(x, y interface{}) interface{} { return x.(int64) >= y.(int64) }
	binaryTokenGEQ[IF] = func(x, y interface{}) interface{} { return float64(x.(int64)) >= y.(float64) }
	binaryTokenGEQ[FB] = func(x, y interface{}) interface{} { return x.(float64) >= b2f(y.(bool)) }
	binaryTokenGEQ[FI] = func(x, y interface{}) interface{} { return x.(float64) >= float64(y.(int64)) }
	binaryTokenGEQ[FF] = func(x, y interface{}) interface{} { return x.(float64) >= y.(float64) }
	binaryTokenGEQ[SS] = func(x, y interface{}) interface{} { return x.(string) >= y.(string) }
	binaryTokens[token.GEQ] = binaryTokenGEQ

	// SHL
	binaryTokenSHL[II] = func(x, y interface{}) interface{} { return x.(int64) << y.(int64) }
	binaryTokens[token.SHL] = binaryTokenSHL

	// SHR
	binaryTokenSHR[II] = func(x, y interface{}) interface{} { return x.(int64) >> y.(int64) }
	binaryTokens[token.SHR] = binaryTokenSHR
}

func evalBinary(binary *ast.BinaryExpr, variables map[string]interface{}) (interface{}, error) {
	x, err := Eval(binary.X, variables)
	if err != nil {
		return nil, err
	}
	y, err := Eval(binary.Y, variables)
	if err != nil {
		return nil, err
	}
	kx := reflect.ValueOf(x).Kind()
	ky := reflect.ValueOf(y).Kind()
	handler := binaryTokens[binary.Op][kx*32+ky]
	if handler == nil {
		return nil, fmt.Errorf("[binary] illegal expr (%s %s %s)", kx.String(), binary.Op.String(), ky.String())
	}
	return handler(x, y), nil
}
