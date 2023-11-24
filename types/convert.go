package types

import (
	"reflect"
)

// Value common value
type Value struct {
	B bool
	I int64
	F float64
	S string
}

// Convert convert bool,int64,float64,string
func Convert(x any, value *Value) reflect.Kind {
	switch v := x.(type) {
	case bool:
		value.B = v
		return reflect.Bool
	case int:
		value.I = int64(v)
		return reflect.Int64
	case int64:
		value.I = v
		return reflect.Int64
	case string:
		value.S = v
		return reflect.String
	case float64:
		value.F = v
		return reflect.Float64
	case int32:
		value.I = int64(v)
		return reflect.Int64
	case float32:
		value.F = float64(v)
		return reflect.Float64
	case uint:
		value.I = int64(v)
		return reflect.Int64
	case uint64:
		value.I = int64(v)
		return reflect.Int64
	case uint32:
		value.I = int64(v)
		return reflect.Int64
	case int16:
		value.I = int64(v)
		return reflect.Int64
	case uint16:
		value.I = int64(v)
		return reflect.Int64
	case int8:
		value.I = int64(v)
		return reflect.Int64
	case uint8:
		value.I = int64(v)
		return reflect.Int64
	}
	return reflect.Invalid
}

// ConvertInt convert all integer kinds to int
func ConvertInt(x any) (int, bool) {
	rv := reflect.ValueOf(x)
	if rv.CanInt() {
		return int(rv.Int()), true
	}
	if rv.CanUint() {
		return int(rv.Uint()), true
	}
	return 0, false
}
