package builtin

import "reflect"

func init() {
	Variables["len"] = rlen
	Variables["cap"] = rcap
	Variables["exist"] = exist
}

func rlen(v interface{}) int {
	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind == reflect.Slice ||
		kind == reflect.String ||
		kind == reflect.Map ||
		kind == reflect.Array ||
		kind == reflect.Chan {
		return rv.Len()
	}
	return 0
}

func rcap(v interface{}) int {
	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind == reflect.Slice ||
		kind == reflect.Array ||
		kind == reflect.Chan {
		return rv.Cap()
	}
	return 0
}

func exist(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func,
		reflect.UnsafePointer:
		return !rv.IsNil()
	}
	return true
}
