package builtin

import (
	"math"
)

func init() {
	variables["abs"] = math.Abs

	variables["ceil"] = math.Ceil
	variables["round"] = math.Round
	variables["floor"] = math.Floor

	variables["sqrt"] = math.Sqrt
	variables["pow"] = math.Pow
	variables["exp"] = math.Exp
	variables["log"] = math.Log
	variables["log10"] = math.Log10
	variables["log2"] = math.Log2

	variables["sin"] = math.Sin
	variables["cos"] = math.Cos
	variables["tan"] = math.Tan

	variables["max"] = math.Max
	variables["min"] = math.Min
	variables["clamp"] = func(x float64, min float64, max float64) float64 { return math.Min(math.Max(x, min), max) }

	variables["sigmoid"] = func(x float64) float64 { return 1 / (1 + math.Exp(-x)) }
	variables["tanh"] = math.Tanh

	variables["isnan"] = math.IsNaN
	variables["isinf"] = math.IsInf
}
