package builtin

import "fmt"

var (
	variables map[string]any = map[string]any{}
)

// Vars a simple implementation for eval.Variables
type Vars map[string]any

// Get implement the Get interface
func (o Vars) Get(name string) (any, error) {
	if val, ok := o[name]; ok {
		return val, nil
	}
	if val, ok := variables[name]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("unknown name(%s)", name)
}
