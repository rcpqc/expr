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
	BU = types.MaxKinds*reflect.Bool + reflect.Uint64
	BF = types.MaxKinds*reflect.Bool + reflect.Float64
	BS = types.MaxKinds*reflect.Bool + reflect.String

	IB = types.MaxKinds*reflect.Int64 + reflect.Bool
	II = types.MaxKinds*reflect.Int64 + reflect.Int64
	IU = types.MaxKinds*reflect.Int64 + reflect.Uint64
	IF = types.MaxKinds*reflect.Int64 + reflect.Float64
	IS = types.MaxKinds*reflect.Int64 + reflect.String

	UB = types.MaxKinds*reflect.Uint64 + reflect.Bool
	UI = types.MaxKinds*reflect.Uint64 + reflect.Int64
	UU = types.MaxKinds*reflect.Uint64 + reflect.Uint64
	UF = types.MaxKinds*reflect.Uint64 + reflect.Float64
	US = types.MaxKinds*reflect.Uint64 + reflect.String

	FB = types.MaxKinds*reflect.Float64 + reflect.Bool
	FI = types.MaxKinds*reflect.Float64 + reflect.Int64
	FU = types.MaxKinds*reflect.Float64 + reflect.Uint64
	FF = types.MaxKinds*reflect.Float64 + reflect.Float64
	FS = types.MaxKinds*reflect.Float64 + reflect.String

	SB = types.MaxKinds*reflect.String + reflect.Bool
	SI = types.MaxKinds*reflect.String + reflect.Int64
	SU = types.MaxKinds*reflect.String + reflect.Uint64
	SF = types.MaxKinds*reflect.String + reflect.Float64
	SS = types.MaxKinds*reflect.String + reflect.String

	MAX_TOKEN = 96
)

var kinds = [types.MaxKinds]reflect.Kind{
	reflect.Invalid, reflect.Bool,
	reflect.Int64, reflect.Int64, reflect.Int64, reflect.Int64, reflect.Int64,
	reflect.Uint64, reflect.Uint64, reflect.Uint64, reflect.Uint64, reflect.Uint64, reflect.Uint64,
	reflect.Float64, reflect.Float64, reflect.Complex64, reflect.Complex128,
	reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice,
	reflect.String, reflect.Struct, reflect.UnsafePointer,
}

type binaryKind func(x, y reflect.Value) any
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
	binaryADD[BB] = func(x, y reflect.Value) any { return b2i(x.Bool()) + b2i(y.Bool()) }
	binaryADD[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) + y.Int() }
	binaryADD[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) + int64(y.Uint()) }
	binaryADD[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) + y.Float() }
	binaryADD[IB] = func(x, y reflect.Value) any { return x.Int() + b2i(y.Bool()) }
	binaryADD[II] = func(x, y reflect.Value) any { return x.Int() + y.Int() }
	binaryADD[IU] = func(x, y reflect.Value) any { return x.Int() + int64(y.Uint()) }
	binaryADD[IF] = func(x, y reflect.Value) any { return float64(x.Int()) + y.Float() }
	binaryADD[UB] = func(x, y reflect.Value) any { return int64(x.Uint()) + b2i(y.Bool()) }
	binaryADD[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) + y.Int() }
	binaryADD[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) + int64(y.Uint()) }
	binaryADD[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) + y.Float() }
	binaryADD[FB] = func(x, y reflect.Value) any { return x.Float() + b2f(y.Bool()) }
	binaryADD[FI] = func(x, y reflect.Value) any { return x.Float() + float64(y.Int()) }
	binaryADD[FU] = func(x, y reflect.Value) any { return x.Float() + float64(y.Uint()) }
	binaryADD[FF] = func(x, y reflect.Value) any { return x.Float() + y.Float() }
	binaryADD[SS] = func(x, y reflect.Value) any { return x.String() + y.String() }
	binaryTokens[token.ADD] = binaryADD

	// SUB
	binarySUB[BB] = func(x, y reflect.Value) any { return b2i(x.Bool()) - b2i(y.Bool()) }
	binarySUB[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) - y.Int() }
	binarySUB[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) - int64(y.Uint()) }
	binarySUB[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) - y.Float() }
	binarySUB[IB] = func(x, y reflect.Value) any { return x.Int() - b2i(y.Bool()) }
	binarySUB[II] = func(x, y reflect.Value) any { return x.Int() - y.Int() }
	binarySUB[IU] = func(x, y reflect.Value) any { return x.Int() - int64(y.Uint()) }
	binarySUB[IF] = func(x, y reflect.Value) any { return float64(x.Int()) - y.Float() }
	binarySUB[UB] = func(x, y reflect.Value) any { return int64(x.Uint()) - b2i(y.Bool()) }
	binarySUB[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) - y.Int() }
	binarySUB[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) - int64(y.Uint()) }
	binarySUB[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) - y.Float() }
	binarySUB[FB] = func(x, y reflect.Value) any { return x.Float() - b2f(y.Bool()) }
	binarySUB[FI] = func(x, y reflect.Value) any { return x.Float() - float64(y.Int()) }
	binarySUB[FU] = func(x, y reflect.Value) any { return x.Float() - float64(y.Uint()) }
	binarySUB[FF] = func(x, y reflect.Value) any { return x.Float() - y.Float() }
	binaryTokens[token.SUB] = binarySUB

	// MUL
	binaryMUL[BB] = func(x, y reflect.Value) any { return b2i(x.Bool()) * b2i(y.Bool()) }
	binaryMUL[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) * y.Int() }
	binaryMUL[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) * int64(y.Uint()) }
	binaryMUL[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) * y.Float() }
	binaryMUL[IB] = func(x, y reflect.Value) any { return x.Int() * b2i(y.Bool()) }
	binaryMUL[II] = func(x, y reflect.Value) any { return x.Int() * y.Int() }
	binaryMUL[IU] = func(x, y reflect.Value) any { return x.Int() * int64(y.Uint()) }
	binaryMUL[IF] = func(x, y reflect.Value) any { return float64(x.Int()) * y.Float() }
	binaryMUL[UB] = func(x, y reflect.Value) any { return int64(x.Uint()) * b2i(y.Bool()) }
	binaryMUL[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) * y.Int() }
	binaryMUL[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) * int64(y.Uint()) }
	binaryMUL[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) * y.Float() }
	binaryMUL[FB] = func(x, y reflect.Value) any { return x.Float() * b2f(y.Bool()) }
	binaryMUL[FI] = func(x, y reflect.Value) any { return x.Float() * float64(y.Int()) }
	binaryMUL[FU] = func(x, y reflect.Value) any { return x.Float() * float64(y.Uint()) }
	binaryMUL[FF] = func(x, y reflect.Value) any { return x.Float() * y.Float() }
	binaryTokens[token.MUL] = binaryMUL

	// QUO
	binaryQUO[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) / y.Int() }
	binaryQUO[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) / int64(y.Uint()) }
	binaryQUO[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) / y.Float() }
	binaryQUO[II] = func(x, y reflect.Value) any { return x.Int() / y.Int() }
	binaryQUO[IU] = func(x, y reflect.Value) any { return x.Int() / int64(y.Uint()) }
	binaryQUO[IF] = func(x, y reflect.Value) any { return float64(x.Int()) / y.Float() }
	binaryQUO[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) / y.Int() }
	binaryQUO[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) / int64(y.Uint()) }
	binaryQUO[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) / y.Float() }
	binaryQUO[FI] = func(x, y reflect.Value) any { return x.Float() / float64(y.Int()) }
	binaryQUO[FU] = func(x, y reflect.Value) any { return x.Float() / float64(y.Uint()) }
	binaryQUO[FF] = func(x, y reflect.Value) any { return x.Float() / y.Float() }
	binaryTokens[token.QUO] = binaryQUO

	// REM
	binaryREM[II] = func(x, y reflect.Value) any { return x.Int() % y.Int() }
	binaryREM[IU] = func(x, y reflect.Value) any { return x.Int() % int64(y.Uint()) }
	binaryREM[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) % y.Int() }
	binaryREM[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) % int64(y.Uint()) }
	binaryTokens[token.REM] = binaryREM

	// AND
	binaryAND[II] = func(x, y reflect.Value) any { return x.Int() & y.Int() }
	binaryAND[IU] = func(x, y reflect.Value) any { return x.Int() & int64(y.Uint()) }
	binaryAND[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) & y.Int() }
	binaryAND[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) & int64(y.Uint()) }
	binaryTokens[token.AND] = binaryAND

	// OR
	binaryOR[II] = func(x, y reflect.Value) any { return x.Int() | y.Int() }
	binaryOR[IU] = func(x, y reflect.Value) any { return x.Int() | int64(y.Uint()) }
	binaryOR[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) | y.Int() }
	binaryOR[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) | int64(y.Uint()) }
	binaryTokens[token.OR] = binaryOR

	// XOR
	binaryXOR[II] = func(x, y reflect.Value) any { return x.Int() ^ y.Int() }
	binaryXOR[IU] = func(x, y reflect.Value) any { return x.Int() ^ int64(y.Uint()) }
	binaryXOR[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) ^ y.Int() }
	binaryXOR[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) ^ int64(y.Uint()) }
	binaryTokens[token.XOR] = binaryXOR

	// LAND
	binaryLAND[BB] = func(x, y reflect.Value) any { return x.Bool() && y.Bool() }
	binaryTokens[token.LAND] = binaryLAND

	// LOR
	binaryLOR[BB] = func(x, y reflect.Value) any { return x.Bool() || y.Bool() }
	binaryTokens[token.LOR] = binaryLOR

	// EQL
	binaryEQL[BB] = func(x, y reflect.Value) any { return b2i(x.Bool()) == b2i(y.Bool()) }
	binaryEQL[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) == y.Int() }
	binaryEQL[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) == int64(y.Uint()) }
	binaryEQL[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) == y.Float() }
	binaryEQL[IB] = func(x, y reflect.Value) any { return x.Int() == b2i(y.Bool()) }
	binaryEQL[II] = func(x, y reflect.Value) any { return x.Int() == y.Int() }
	binaryEQL[IU] = func(x, y reflect.Value) any { return x.Int() == int64(y.Uint()) }
	binaryEQL[IF] = func(x, y reflect.Value) any { return float64(x.Int()) == y.Float() }
	binaryEQL[UB] = func(x, y reflect.Value) any { return int64(x.Uint()) == b2i(y.Bool()) }
	binaryEQL[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) == y.Int() }
	binaryEQL[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) == int64(y.Uint()) }
	binaryEQL[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) == y.Float() }
	binaryEQL[FB] = func(x, y reflect.Value) any { return x.Float() == b2f(y.Bool()) }
	binaryEQL[FI] = func(x, y reflect.Value) any { return x.Float() == float64(y.Int()) }
	binaryEQL[FU] = func(x, y reflect.Value) any { return x.Float() == float64(y.Uint()) }
	binaryEQL[FF] = func(x, y reflect.Value) any { return x.Float() == y.Float() }
	binaryEQL[SS] = func(x, y reflect.Value) any { return x.String() == y.String() }
	binaryTokens[token.EQL] = binaryEQL

	// NEQ
	binaryNEQ[BB] = func(x, y reflect.Value) any { return b2i(x.Bool()) != b2i(y.Bool()) }
	binaryNEQ[BI] = func(x, y reflect.Value) any { return b2i(x.Bool()) != y.Int() }
	binaryNEQ[BU] = func(x, y reflect.Value) any { return b2i(x.Bool()) != int64(y.Uint()) }
	binaryNEQ[BF] = func(x, y reflect.Value) any { return b2f(x.Bool()) != y.Float() }
	binaryNEQ[IB] = func(x, y reflect.Value) any { return x.Int() != b2i(y.Bool()) }
	binaryNEQ[II] = func(x, y reflect.Value) any { return x.Int() != y.Int() }
	binaryNEQ[IU] = func(x, y reflect.Value) any { return x.Int() != int64(y.Uint()) }
	binaryNEQ[IF] = func(x, y reflect.Value) any { return float64(x.Int()) != y.Float() }
	binaryNEQ[UB] = func(x, y reflect.Value) any { return int64(x.Uint()) != b2i(y.Bool()) }
	binaryNEQ[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) != y.Int() }
	binaryNEQ[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) != int64(y.Uint()) }
	binaryNEQ[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) != y.Float() }
	binaryNEQ[FB] = func(x, y reflect.Value) any { return x.Float() != b2f(y.Bool()) }
	binaryNEQ[FI] = func(x, y reflect.Value) any { return x.Float() != float64(y.Int()) }
	binaryNEQ[FU] = func(x, y reflect.Value) any { return x.Float() != float64(y.Uint()) }
	binaryNEQ[FF] = func(x, y reflect.Value) any { return x.Float() != y.Float() }
	binaryNEQ[SS] = func(x, y reflect.Value) any { return x.String() != y.String() }
	binaryTokens[token.NEQ] = binaryNEQ

	// LSS
	binaryLSS[II] = func(x, y reflect.Value) any { return x.Int() < y.Int() }
	binaryLSS[IU] = func(x, y reflect.Value) any { return x.Int() < int64(y.Uint()) }
	binaryLSS[IF] = func(x, y reflect.Value) any { return float64(x.Int()) < y.Float() }
	binaryLSS[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) < y.Int() }
	binaryLSS[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) < int64(y.Uint()) }
	binaryLSS[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) < y.Float() }
	binaryLSS[FI] = func(x, y reflect.Value) any { return x.Float() < float64(y.Int()) }
	binaryLSS[FU] = func(x, y reflect.Value) any { return x.Float() < float64(y.Uint()) }
	binaryLSS[FF] = func(x, y reflect.Value) any { return x.Float() < y.Float() }
	binaryLSS[SS] = func(x, y reflect.Value) any { return x.String() < y.String() }
	binaryTokens[token.LSS] = binaryLSS

	// GTR
	binaryGTR[II] = func(x, y reflect.Value) any { return x.Int() > y.Int() }
	binaryGTR[IU] = func(x, y reflect.Value) any { return x.Int() > int64(y.Uint()) }
	binaryGTR[IF] = func(x, y reflect.Value) any { return float64(x.Int()) > y.Float() }
	binaryGTR[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) > y.Int() }
	binaryGTR[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) > int64(y.Uint()) }
	binaryGTR[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) > y.Float() }
	binaryGTR[FI] = func(x, y reflect.Value) any { return x.Float() > float64(y.Int()) }
	binaryGTR[FU] = func(x, y reflect.Value) any { return x.Float() > float64(y.Uint()) }
	binaryGTR[FF] = func(x, y reflect.Value) any { return x.Float() > y.Float() }
	binaryGTR[SS] = func(x, y reflect.Value) any { return x.String() > y.String() }
	binaryTokens[token.GTR] = binaryGTR

	// LEQ
	binaryLEQ[II] = func(x, y reflect.Value) any { return x.Int() <= y.Int() }
	binaryLEQ[IU] = func(x, y reflect.Value) any { return x.Int() <= int64(y.Uint()) }
	binaryLEQ[IF] = func(x, y reflect.Value) any { return float64(x.Int()) <= y.Float() }
	binaryLEQ[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) <= y.Int() }
	binaryLEQ[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) <= int64(y.Uint()) }
	binaryLEQ[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) <= y.Float() }
	binaryLEQ[FI] = func(x, y reflect.Value) any { return x.Float() <= float64(y.Int()) }
	binaryLEQ[FU] = func(x, y reflect.Value) any { return x.Float() <= float64(y.Uint()) }
	binaryLEQ[FF] = func(x, y reflect.Value) any { return x.Float() <= y.Float() }
	binaryLEQ[SS] = func(x, y reflect.Value) any { return x.String() <= y.String() }
	binaryTokens[token.LEQ] = binaryLEQ

	// GEQ
	binaryGEQ[II] = func(x, y reflect.Value) any { return x.Int() >= y.Int() }
	binaryGEQ[IU] = func(x, y reflect.Value) any { return x.Int() >= int64(y.Uint()) }
	binaryGEQ[IF] = func(x, y reflect.Value) any { return float64(x.Int()) >= y.Float() }
	binaryGEQ[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) >= y.Int() }
	binaryGEQ[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) >= int64(y.Uint()) }
	binaryGEQ[UF] = func(x, y reflect.Value) any { return float64(x.Uint()) >= y.Float() }
	binaryGEQ[FI] = func(x, y reflect.Value) any { return x.Float() >= float64(y.Int()) }
	binaryGEQ[FU] = func(x, y reflect.Value) any { return x.Float() >= float64(y.Uint()) }
	binaryGEQ[FF] = func(x, y reflect.Value) any { return x.Float() >= y.Float() }
	binaryGEQ[SS] = func(x, y reflect.Value) any { return x.String() >= y.String() }
	binaryTokens[token.GEQ] = binaryGEQ

	// SHL
	binarySHL[II] = func(x, y reflect.Value) any { return x.Int() << y.Int() }
	binarySHL[IU] = func(x, y reflect.Value) any { return x.Int() << int64(y.Uint()) }
	binarySHL[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) << y.Int() }
	binarySHL[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) << int64(y.Uint()) }
	binaryTokens[token.SHL] = binarySHL

	// SHR
	binarySHR[II] = func(x, y reflect.Value) any { return x.Int() >> y.Int() }
	binarySHR[IU] = func(x, y reflect.Value) any { return x.Int() >> int64(y.Uint()) }
	binarySHR[UI] = func(x, y reflect.Value) any { return int64(x.Uint()) >> y.Int() }
	binarySHR[UU] = func(x, y reflect.Value) any { return int64(x.Uint()) >> int64(y.Uint()) }
	binaryTokens[token.SHR] = binarySHR
}

func evalBinary(expr ast.Expr, variables Variables) (any, error) {
	binary := expr.(*ast.BinaryExpr)
	x, err := evaluator(binary.X)(binary.X, variables)
	if err != nil {
		return nil, err
	}
	xvalue := reflect.ValueOf(x)
	kx := kinds[xvalue.Kind()]
	if binary.Op == token.LAND && kx == reflect.Bool && !xvalue.Bool() {
		return false, nil
	}
	if binary.Op == token.LOR && kx == reflect.Bool && xvalue.Bool() {
		return true, nil
	}
	y, err := evaluator(binary.Y)(binary.Y, variables)
	if err != nil {
		return nil, err
	}
	yvalue := reflect.ValueOf(y)
	ky := kinds[yvalue.Kind()]
	if binary.Op == token.QUO &&
		(ky == reflect.Int64 && yvalue.Int() == 0) || (ky == reflect.Uint64 && yvalue.Uint() == 0) {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	if binary.Op == token.REM &&
		(ky == reflect.Int64 && yvalue.Int() == 0) || (ky == reflect.Uint64 && yvalue.Uint() == 0) {
		return nil, errs.Newf(binary, "integer divide by zero")
	}
	handler := binaryTokens[binary.Op][kx*types.MaxKinds+ky]
	if handler == nil {
		return nil, errs.Newf(binary, "illegal expr(%v%s%v)", reflect.TypeOf(x), binary.Op, reflect.TypeOf(y))
	}
	return handler(xvalue, yvalue), nil
}
