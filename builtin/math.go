package builtin

import (
	"math"
)

func init() {
	Variables["abs"] = math.Abs

	Variables["ceil"] = math.Ceil
	Variables["round"] = math.Round
	Variables["floor"] = math.Floor

	Variables["sqrt"] = math.Sqrt
	Variables["pow"] = math.Pow
	Variables["exp"] = math.Exp
	Variables["log"] = math.Log
	Variables["log10"] = math.Log10
	Variables["log2"] = math.Log2

	Variables["sin"] = math.Sin
	Variables["cos"] = math.Cos
	Variables["tan"] = math.Tan

	Variables["max"] = math.Max
	Variables["min"] = math.Min
	Variables["clamp"] = func(x float64, min float64, max float64) float64 { return math.Min(math.Max(x, min), max) }

	Variables["sigmoid"] = func(x float64) float64 { return 1 / (1 + math.Exp(-x)) }
	Variables["tanh"] = math.Tanh

	Variables["isnan"] = math.IsNaN
	Variables["isinf"] = math.IsInf
}
