package builtin

func init() {
	variables["float32"] = func(v float32) float32 { return v }
	variables["float64"] = func(v float64) float64 { return v }
}
