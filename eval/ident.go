package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/builtin"
)

func evalIdent(ident *ast.Ident, variables map[string]interface{}) (interface{}, error) {
	if v, ok := variables[ident.Name]; ok {
		switch reflect.ValueOf(v).Kind() {
		case reflect.Int:
			return int64(v.(int)), nil
		case reflect.Int8:
			return int64(v.(int8)), nil
		case reflect.Int16:
			return int64(v.(int16)), nil
		case reflect.Int32:
			return int64(v.(int32)), nil
		case reflect.Uint:
			return int64(v.(uint)), nil
		case reflect.Uint8:
			return int64(v.(uint8)), nil
		case reflect.Uint16:
			return int64(v.(uint16)), nil
		case reflect.Uint32:
			return int64(v.(uint32)), nil
		case reflect.Uint64:
			return int64(v.(uint64)), nil
		case reflect.Float32:
			return float64(v.(float32)), nil
		default:
			return v, nil
		}
	}
	if fn, ok := builtin.Functions[ident.Name]; ok {
		return fn, nil
	}
	return nil, fmt.Errorf("[ident] unknown ident(%s)", ident.Name)
}
