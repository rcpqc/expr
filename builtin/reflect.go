package builtin

import "reflect"

func init() {
	Variables["len"] = rlen
	Variables["cap"] = rcap
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
