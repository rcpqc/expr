package types

import (
	"reflect"
	"sync"
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
