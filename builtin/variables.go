package builtin

var (
	variables map[string]any = map[string]any{}
)

// Get implement the Get interface
func Get(name string) (any, bool) {
	val, ok := variables[name]
	return val, ok
}
