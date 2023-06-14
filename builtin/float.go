package builtin

func init() {
	Variables["float32"] = func(v float32) float32 { return v }
	Variables["float64"] = func(v float64) float64 { return v }
}
