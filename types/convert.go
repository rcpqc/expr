package types

import (
	"reflect"
)

// ConvertInt convert all integer kinds to int
func ConvertInt(x any) (int, bool) {
	rv := reflect.ValueOf(x)
	if rv.CanInt() {
		return int(rv.Int()), true
	}
	if rv.CanUint() {
		return int(rv.Uint()), true
	}
	return 0, false
}
