package builtin

func init() {
	Functions["float32"] = func(v float32) float32 { return v }
	Functions["float64"] = func(v float64) float64 { return v }
}
