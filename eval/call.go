package eval

import (
	"fmt"
	"go/ast"
	"reflect"
)

func evalNonVariadicCall(rvfn reflect.Value, rvargs []reflect.Value) (interface{}, error) {
	rtfn := rvfn.Type()
	if rtfn.NumIn() != len(rvargs) {
		return nil, fmt.Errorf("[call] func(%s) input parameters want(%d) got(%d)", rtfn.Name(), rtfn.NumIn(), len(rvargs))
	}
	in := []reflect.Value{}
	for i := 0; i < rtfn.NumIn(); i++ {
		rvarg := rvargs[i]
		if rtfn.In(i) == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(rtfn.In(i)) {
			in = append(in, rvarg.Convert(rtfn.In(i)))
			continue
		}
		return nil, fmt.Errorf("[call] arg%d want %v got %v", i, rtfn.In(i), rvarg.Type())
	}
	out := rvfn.Call(in)
	return out[0].Interface(), nil
}

func evalVariadicCall(rvfn reflect.Value, rvargs []reflect.Value) (interface{}, error) {
	rtfn := rvfn.Type()
	if rtfn.NumIn()-1 > len(rvargs) {
		return nil, fmt.Errorf("[call] func(%s) input parameters want(>=%d) got(%d)", rtfn.Name(), rtfn.NumIn()-1, len(rvargs))
	}
	in := []reflect.Value{}
	for i := 0; i < rtfn.NumIn()-1; i++ {
		rvarg := rvargs[i]
		if rtfn.In(i) == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(rtfn.In(i)) {
			in = append(in, rvarg.Convert(rtfn.In(i)))
			continue
		}
		return nil, fmt.Errorf("[call] arg%d want %v got %v", i, rtfn.In(i), rvarg.Type())
	}
	variadicType := rtfn.In(rtfn.NumIn() - 1).Elem()
	for i := rtfn.NumIn() - 1; i < len(rvargs); i++ {
		rvarg := rvargs[i]
		if variadicType == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(variadicType) {
			in = append(in, rvarg.Convert(variadicType))
			continue
		}
		return nil, fmt.Errorf("[call] arg%d want %v got %v", i, variadicType, rvarg.Type())
	}
	out := rvfn.Call(in)
	return out[0].Interface(), nil
}

func evalCall(expr *ast.CallExpr, variables map[string]interface{}) (interface{}, error) {
	fn, err := evalExpr(expr.Fun, variables)
	if err != nil {
		return nil, err
	}
	rvfn := reflect.ValueOf(fn)
	if rvfn.Type().Kind() != reflect.Func {
		return nil, fmt.Errorf("[call] not a func")
	}
	if rvfn.Type().NumOut() != 1 {
		return nil, fmt.Errorf("[call] func(%s) output parameters want(1) got(%d)", rvfn.Type().Name(), rvfn.Type().NumOut())
	}

	rvargs := []reflect.Value{}
	for _, argexpr := range expr.Args {
		arg, err := evalExpr(argexpr, variables)
		if err != nil {
			return nil, err
		}
		rvargs = append(rvargs, reflect.ValueOf(arg))
	}

	if rvfn.Type().IsVariadic() {
		return evalVariadicCall(rvfn, rvargs)
	} else {
		return evalNonVariadicCall(rvfn, rvargs)
	}
}
