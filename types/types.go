package types

import (
	"reflect"
)

const MaxKinds = 32

var (
	Any = reflect.TypeOf((*any)(nil)).Elem()

	Bool = reflect.TypeOf((*bool)(nil)).Elem()
	Byte = reflect.TypeOf((*byte)(nil)).Elem()

	Int   = reflect.TypeOf((*int)(nil)).Elem()
	Int8  = reflect.TypeOf((*int8)(nil)).Elem()
	Int16 = reflect.TypeOf((*int16)(nil)).Elem()
	Int32 = reflect.TypeOf((*int32)(nil)).Elem()
	Int64 = reflect.TypeOf((*int64)(nil)).Elem()

	Uint   = reflect.TypeOf((*uint)(nil)).Elem()
	Uint8  = reflect.TypeOf((*uint8)(nil)).Elem()
	Uint16 = reflect.TypeOf((*uint16)(nil)).Elem()
	Uint32 = reflect.TypeOf((*uint32)(nil)).Elem()
	Uint64 = reflect.TypeOf((*uint64)(nil)).Elem()

	Float32 = reflect.TypeOf((*float32)(nil)).Elem()
	Float64 = reflect.TypeOf((*float64)(nil)).Elem()

	String = reflect.TypeOf((*string)(nil)).Elem()
	Bytes  = reflect.TypeOf((*[]byte)(nil)).Elem()
)

var byName = map[string]reflect.Type{
	"any": Any,

	"bool": Bool,
	"byte": Byte,

	"int":   Int,
	"int8":  Int8,
	"int16": Int16,
	"int32": Int32,
	"int64": Int64,

	"uint":   Uint,
	"uint8":  Uint8,
	"uint16": Uint16,
	"uint32": Uint32,
	"uint64": Uint64,

	"float32": Float32,
	"float64": Float64,

	"string": String,
}

func ByName(name string) (reflect.Type, bool) {
	t, ok := byName[name]
	return t, ok
}
