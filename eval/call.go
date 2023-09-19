package eval

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/rcpqc/expr/errs"
)

func evalNonVariadicCall(rvfn reflect.Value, rvargs []reflect.Value) (any, error) {
	rtfn := rvfn.Type()
	if rtfn.NumIn() != len(rvargs) {
		return nil, fmt.Errorf("input parameters want(%d) got(%d)", rtfn.NumIn(), len(rvargs))
	}
	in := []reflect.Value{}
	for i := 0; i < rtfn.NumIn(); i++ {
		rvarg := rvargs[i]
		if !rvarg.IsValid() {
			return nil, fmt.Errorf("arg%d is invalid", i)
		}
		if rtfn.In(i) == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(rtfn.In(i)) {
			in = append(in, rvarg.Convert(rtfn.In(i)))
			continue
		}
		return nil, fmt.Errorf("arg%d want %v got %v", i, rtfn.In(i), rvarg.Type())
	}
	out := rvfn.Call(in)
	return out[0].Interface(), nil
}

func evalVariadicCall(rvfn reflect.Value, rvargs []reflect.Value) (any, error) {
	rtfn := rvfn.Type()
	if rtfn.NumIn()-1 > len(rvargs) {
		return nil, fmt.Errorf("input parameters want(>=%d) got(%d)", rtfn.NumIn()-1, len(rvargs))
	}
	in := []reflect.Value{}
	for i := 0; i < rtfn.NumIn()-1; i++ {
		rvarg := rvargs[i]
		if !rvarg.IsValid() {
			return nil, fmt.Errorf("arg%d is invalid", i)
		}
		if rtfn.In(i) == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(rtfn.In(i)) {
			in = append(in, rvarg.Convert(rtfn.In(i)))
			continue
		}
		return nil, fmt.Errorf("arg%d want %v got %v", i, rtfn.In(i), rvarg.Type())
	}
	variadicType := rtfn.In(rtfn.NumIn() - 1).Elem()
	for i := rtfn.NumIn() - 1; i < len(rvargs); i++ {
		rvarg := rvargs[i]
		if !rvarg.IsValid() {
			return nil, fmt.Errorf("arg%d is invalid", i)
		}
		if variadicType == rvarg.Type() {
			in = append(in, rvarg)
			continue
		}
		if rvarg.CanConvert(variadicType) {
			in = append(in, rvarg.Convert(variadicType))
			continue
		}
		return nil, fmt.Errorf("arg%d want %v got %v", i, variadicType, rvarg.Type())
	}
	out := rvfn.Call(in)
	return out[0].Interface(), nil
}

func evalCall(call *ast.CallExpr, variables Variables) (any, error) {
	fn, err := eval(call.Fun, variables)
	if err != nil {
		return nil, err
	}
	rvfn := reflect.ValueOf(fn)
	if rvfn.Kind() != reflect.Func {
		return nil, errs.Newf(call, "not a func")
	}
	if rvfn.Type().NumOut() != 1 {
		return nil, errs.Newf(call, "output parameters want(1) got(%d)", rvfn.Type().NumOut())
	}

	rvargs := []reflect.Value{}
	for _, argexpr := range call.Args {
		arg, err := eval(argexpr, variables)
		if err != nil {
			return nil, err
		}
		rvargs = append(rvargs, reflect.ValueOf(arg))
	}

	if rvfn.Type().IsVariadic() {
		val, err := evalVariadicCall(rvfn, rvargs)
		return val, errs.New(call, err)
	} else {
		val, err := evalNonVariadicCall(rvfn, rvargs)
		return val, errs.New(call, err)
	}
}
