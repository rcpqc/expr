package expr

import (
	"github.com/rcpqc/expr/eval"
)

var (
	Eval        = eval.Eval
	EvalInt     = eval.EvalInt
	EvalInt8    = eval.EvalInt8
	EvalInt16   = eval.EvalInt16
	EvalInt32   = eval.EvalInt32
	EvalInt64   = eval.EvalInt64
	EvalUint    = eval.EvalUint
	EvalUint8   = eval.EvalUint8
	EvalUint16  = eval.EvalUint16
	EvalUint32  = eval.EvalUint32
	EvalUint64  = eval.EvalUint64
	EvalFloat32 = eval.EvalFloat32
	EvalFloat64 = eval.EvalFloat64
	EvalString  = eval.EvalString
	EvalBytes   = eval.EvalBytes
)
