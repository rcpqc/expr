package types

import (
	"reflect"
	"sync"
)

var (
	BoolType  = reflect.TypeOf((*bool)(nil)).Elem()
	IntType   = reflect.TypeOf((*int)(nil)).Elem()
	Int8Type  = reflect.TypeOf((*int8)(nil)).Elem()
	Int16Type = reflect.TypeOf((*int16)(nil)).Elem()
	Int32Type = reflect.TypeOf((*int32)(nil)).Elem()
	Int64Type = reflect.TypeOf((*int64)(nil)).Elem()

	UintType   = reflect.TypeOf((*uint)(nil)).Elem()
	Uint8Type  = reflect.TypeOf((*uint8)(nil)).Elem()
	Uint16Type = reflect.TypeOf((*uint16)(nil)).Elem()
	Uint32Type = reflect.TypeOf((*uint32)(nil)).Elem()
	Uint64Type = reflect.TypeOf((*uint64)(nil)).Elem()

	Float32Type = reflect.TypeOf((*float32)(nil)).Elem()
	Float64Type = reflect.TypeOf((*float64)(nil)).Elem()

	StringType = reflect.TypeOf((*string)(nil)).Elem()
	BytesType  = reflect.TypeOf((*[]byte)(nil)).Elem()
)

var cache sync.Map

func LoadOrCreate(t reflect.Type, constructor func(t reflect.Type) interface{}) (interface{}, bool) {
	if f, ok := cache.Load(t); ok {
		return f.(func() interface{})(), true
	}
	var once sync.Once
	var res interface{}
	f, loaded := cache.LoadOrStore(t, func() interface{} {
		once.Do(func() {
			res = constructor(t)
			cache.Store(t, func() interface{} { return res })
		})
		return res
	})
	return f.(func() interface{})(), loaded
}
