package builtin

import "reflect"

func init() {
	Functions["len"] = _len
	Functions["cap"] = _cap
}

func _len(v interface{}) int {
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

func _cap(v interface{}) int {
	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind == reflect.Slice ||
		kind == reflect.Array ||
		kind == reflect.Chan {
		return rv.Cap()
	}
	return 0
}
