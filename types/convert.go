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
func Convert(x interface{}, value *Value) reflect.Kind {
	if b, ok := x.(bool); ok {
		value.B = b
		return reflect.Bool
	}
	if i, ok := x.(int); ok {
		value.I = int64(i)
		return reflect.Int64
	}
	if i, ok := x.(int64); ok {
		value.I = i
		return reflect.Int64
	}
	if s, ok := x.(string); ok {
		value.S = s
		return reflect.String
	}
	if f, ok := x.(float64); ok {
		value.F = f
		return reflect.Float64
	}
	if i, ok := x.(int32); ok {
		value.I = int64(i)
		return reflect.Int64
	}
	if f, ok := x.(float32); ok {
		value.F = float64(f)
		return reflect.Float64
	}
	if u, ok := x.(uint); ok {
		value.I = int64(u)
		return reflect.Int64
	}
	if u, ok := x.(uint64); ok {
		value.I = int64(u)
		return reflect.Int64
	}
	if u, ok := x.(uint32); ok {
		value.I = int64(u)
		return reflect.Int64
	}
	if i, ok := x.(int16); ok {
		value.I = int64(i)
		return reflect.Int64
	}
	if u, ok := x.(uint16); ok {
		value.I = int64(u)
		return reflect.Int64
	}
	if i, ok := x.(int8); ok {
		value.I = int64(i)
		return reflect.Int64
	}
	if u, ok := x.(uint8); ok {
		value.I = int64(u)
		return reflect.Int64
	}
	return reflect.Invalid
}

// ConvertInt convert all integer kinds to int
func ConvertInt(x interface{}) (int, bool) {
	rv := reflect.ValueOf(x)
	if rv.CanInt() {
		return int(rv.Int()), true
	}
	if rv.CanUint() {
		return int(rv.Uint()), true
	}
	return 0, false
}
