package builtin

import "reflect"

func init() {
	Functions["len"] = rlen
	Functions["cap"] = rcap
	Functions["has"] = rhas
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

func rhas(element interface{}, container interface{}) bool {
	rc := reflect.ValueOf(container)
	re := reflect.ValueOf(element)
	if !re.IsValid() {
		return false
	}
	switch rc.Kind() {
	case reflect.Map:
		return re.Type() == rc.Type().Key() && rc.MapIndex(re).IsValid()
	case reflect.Slice, reflect.Array:
		for i := 0; i < rc.Len(); i++ {
			if reflect.DeepEqual(rc.Index(i).Interface(), element) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
