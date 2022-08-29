package eval

import (
	"reflect"

	"github.com/rcpqc/expr/types"
)

var converters = map[reflect.Type]func(x interface{}) (interface{}, reflect.Kind){
	types.BoolType:    func(x interface{}) (interface{}, reflect.Kind) { return x.(bool), reflect.Bool },
	types.IntType:     func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(int)), reflect.Int64 },
	types.Int8Type:    func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(int8)), reflect.Int64 },
	types.Int16Type:   func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(int16)), reflect.Int64 },
	types.Int32Type:   func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(int32)), reflect.Int64 },
	types.Int64Type:   func(x interface{}) (interface{}, reflect.Kind) { return x.(int64), reflect.Int64 },
	types.UintType:    func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(uint)), reflect.Int64 },
	types.Uint8Type:   func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(uint8)), reflect.Int64 },
	types.Uint16Type:  func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(uint16)), reflect.Int64 },
	types.Uint32Type:  func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(uint32)), reflect.Int64 },
	types.Uint64Type:  func(x interface{}) (interface{}, reflect.Kind) { return int64(x.(uint64)), reflect.Int64 },
	types.Float32Type: func(x interface{}) (interface{}, reflect.Kind) { return float64(x.(float32)), reflect.Float64 },
	types.Float64Type: func(x interface{}) (interface{}, reflect.Kind) { return x.(float64), reflect.Float64 },
	types.StringType:  func(x interface{}) (interface{}, reflect.Kind) { return x.(string), reflect.String },
}
