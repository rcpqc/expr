package types

import (
	"reflect"
)

// Convert convert bool,int,uint,float,string
func Convert(x interface{}) (interface{}, reflect.Type, reflect.Kind) {
	if _, ok := x.(bool); ok {
		return x, Bool, reflect.Bool
	}
	if _, ok := x.(int64); ok {
		return x, Int64, reflect.Int64
	}
	if _, ok := x.(string); ok {
		return x, String, reflect.String
	}
	if _, ok := x.(float64); ok {
		return x, Float64, reflect.Float64
	}
	if i, ok := x.(int); ok {
		return int64(i), Int64, reflect.Int64
	}
	if i32, ok := x.(int32); ok {
		return int64(i32), Int64, reflect.Int64
	}
	if f32, ok := x.(float32); ok {
		return float64(f32), Float64, reflect.Float64
	}
	if u, ok := x.(uint); ok {
		return int64(u), Int64, reflect.Int64
	}
	if u64, ok := x.(uint64); ok {
		return int64(u64), Int64, reflect.Int64
	}
	if u32, ok := x.(uint32); ok {
		return int64(u32), Int64, reflect.Int64
	}
	if i16, ok := x.(int16); ok {
		return int64(i16), Int64, reflect.Int64
	}
	if u16, ok := x.(uint16); ok {
		return int64(u16), Int64, reflect.Int64
	}
	if i8, ok := x.(int8); ok {
		return int64(i8), Int64, reflect.Int64
	}
	if u8, ok := x.(uint8); ok {
		return int64(u8), Int64, reflect.Int64
	}
	if x == nil {
		return x, nil, reflect.Invalid
	}
	rt := reflect.TypeOf(x)
	return x, rt, rt.Kind()
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
