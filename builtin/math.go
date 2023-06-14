package builtin

import (
	"math"
)

func init() {
	Variables["abs"] = func(x float64) float64 { return math.Abs(x) }

	Variables["ceil"] = func(x float64) float64 { return math.Ceil(x) }
	Variables["round"] = func(x float64) float64 { return math.Round(x) }
	Variables["floor"] = func(x float64) float64 { return math.Floor(x) }

	Variables["sqrt"] = func(x float64) float64 { return math.Sqrt(x) }
	Variables["pow"] = func(x float64, y float64) float64 { return math.Pow(x, y) }
	Variables["exp"] = func(x float64) float64 { return math.Exp(x) }
	Variables["log"] = func(x float64) float64 { return math.Log(x) }
	Variables["log10"] = func(x float64) float64 { return math.Log10(x) }
	Variables["log2"] = func(x float64) float64 { return math.Log2(x) }

	Variables["sin"] = func(x float64) float64 { return math.Sin(x) }
	Variables["cos"] = func(x float64) float64 { return math.Cos(x) }
	Variables["tan"] = func(x float64) float64 { return math.Tan(x) }

	Variables["max"] = func(x float64, y float64) float64 { return math.Max(x, y) }
	Variables["min"] = func(x float64, y float64) float64 { return math.Min(x, y) }

	Variables["sigmoid"] = func(x float64) float64 { return 1 / (1 + math.Exp(-x)) }
	Variables["tanh"] = func(x float64) float64 { return math.Tanh(x) }
}
