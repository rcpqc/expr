package expr

import (
	"fmt"
	"go/ast"
	"reflect"
)

type CallFn func(args ...interface{}) interface{}

func evalCall(expr *ast.CallExpr, variables map[string]interface{}) (interface{}, error) {
	fn, err := Eval(expr.Fun, variables)
	if err != nil {
		return nil, err
	}
	rvfn := reflect.ValueOf(fn)
	rtfn := rvfn.Type()
	if rtfn.NumIn() != len(expr.Args) {
		return nil, fmt.Errorf("[call] func(%s) input parameters want(%d) got(%d)", rtfn.Name(), rtfn.NumIn(), len(expr.Args))
	}
	if rtfn.NumOut() != 1 {
		return nil, fmt.Errorf("[call] func(%s) output parameters want(1) got(%d)", rtfn.Name(), rtfn.NumOut())
	}

	in := []reflect.Value{}
	for i, argexpr := range expr.Args {
		arg, err := Eval(argexpr, variables)
		if err != nil {
			return nil, err
		}
		rvarg := reflect.ValueOf(arg)
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
