package builtin

import "reflect"

var Functions = map[string]interface{}{}

func init() {
	Functions["len"] = func(v interface{}) int { return reflect.ValueOf(v).Len() }
	Functions["cap"] = func(v interface{}) int { return reflect.ValueOf(v).Cap() }
}
