package builtin

import (
	"math"
)

func init() {
	Functions["abs"] = func(x float64) float64 { return math.Abs(x) }

	Functions["ceil"] = func(x float64) float64 { return math.Ceil(x) }
	Functions["round"] = func(x float64) float64 { return math.Round(x) }
	Functions["floor"] = func(x float64) float64 { return math.Floor(x) }

	Functions["sqrt"] = func(x float64) float64 { return math.Sqrt(x) }
	Functions["pow"] = func(x float64, y float64) float64 { return math.Pow(x, y) }
	Functions["exp"] = func(x float64) float64 { return math.Exp(x) }
	Functions["log"] = func(x float64) float64 { return math.Log(x) }
	Functions["log10"] = func(x float64) float64 { return math.Log10(x) }
	Functions["log2"] = func(x float64) float64 { return math.Log2(x) }

	Functions["sin"] = func(x float64) float64 { return math.Sin(x) }
	Functions["cos"] = func(x float64) float64 { return math.Cos(x) }
	Functions["tan"] = func(x float64) float64 { return math.Tan(x) }

	Functions["max"] = func(x float64, y float64) float64 { return math.Max(x, y) }
	Functions["min"] = func(x float64, y float64) float64 { return math.Min(x, y) }

	Functions["sigmoid"] = func(x float64) float64 { return 1 / (1 + math.Exp(-x)) }
	Functions["tanh"] = func(x float64) float64 { return math.Tanh(x) }
}
