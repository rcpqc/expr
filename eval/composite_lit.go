package eval

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"

	"github.com/rcpqc/expr/errs"
	"github.com/rcpqc/expr/types"
)

func compositeArray(arrayType *ast.ArrayType) (reflect.Type, error) {
	elemType, err := compositeType(arrayType.Elt)
	if err != nil {
		return nil, err
	}
	if arrayType.Len == nil {
		return reflect.SliceOf(elemType), nil
	}
	basic, ok := arrayType.Len.(*ast.BasicLit)
	if !ok || basic.Kind != token.INT {
		return nil, errs.Newf(arrayType.Len, "illegal expression for array's length")
	}
	length, _ := strconv.ParseInt(basic.Value, 10, 64)
	return reflect.ArrayOf(int(length), elemType), nil
}

func compositeMap(mapType *ast.MapType) (reflect.Type, error) {
	keyType, err := compositeType(mapType.Key)
	if err != nil {
		return nil, err
	}
	valType, err := compositeType(mapType.Value)
	if err != nil {
		return nil, err
	}
	return reflect.MapOf(keyType, valType), nil
}

func compositeIdent(ident *ast.Ident) (reflect.Type, error) {
	t, ok := types.ByName(ident.Name)
	if !ok {
		return nil, errs.Newf(ident, "unsupported type(%v)", ident.Name)
	}
	return t, nil
}

func compositeType(expr ast.Expr) (reflect.Type, error) {
	if arrayType, ok := expr.(*ast.ArrayType); ok {
		return compositeArray(arrayType)
	} else if mapType, ok := expr.(*ast.MapType); ok {
		return compositeMap(mapType)
	} else if ident, ok := expr.(*ast.Ident); ok {
		return compositeIdent(ident)
	}
	return nil, errs.Newf(expr, "illegal composite type")
}

func evalCompositeSlice(expr ast.Expr, variables Variables, elemType reflect.Type) (any, error) {
	composite, ok := expr.(*ast.CompositeLit)
	if !ok {
		return evaltype(expr, variables, elemType)
	}
	slice := reflect.MakeSlice(elemType, len(composite.Elts), len(composite.Elts))
	for i, elem := range composite.Elts {
		val, err := evalCompositeElement(elem, variables, elemType.Elem())
		if err != nil {
			return nil, err
		}
		slice.Index(i).Set(reflect.ValueOf(val))
	}
	return slice.Interface(), nil

}

func evalCompositeArray(expr ast.Expr, variables Variables, elemType reflect.Type) (any, error) {
	composite, ok := expr.(*ast.CompositeLit)
	if !ok {
		return evaltype(expr, variables, elemType)
	}
	array := reflect.New(elemType).Elem()
	arrayLen := array.Len()
	for i, elem := range composite.Elts {
		if i >= arrayLen {
			return nil, errs.Newf(elem, "out of bounds(>=%d) for array", arrayLen)
		}
		val, err := evalCompositeElement(elem, variables, elemType.Elem())
		if err != nil {
			return nil, err
		}
		array.Index(i).Set(reflect.ValueOf(val))
	}
	return array.Interface(), nil
}

func evalCompositeMap(expr ast.Expr, variables Variables, elemType reflect.Type) (any, error) {
	composite, ok := expr.(*ast.CompositeLit)
	if !ok {
		return evaltype(expr, variables, elemType)
	}
	m := reflect.MakeMap(elemType)
	for _, elem := range composite.Elts {
		kv, ok := elem.(*ast.KeyValueExpr)
		if !ok {
			return nil, errs.Newf(elem, "expect key:value as an element of map")
		}
		key, err := evalCompositeElement(kv.Key, variables, elemType.Key())
		if err != nil {
			return nil, err
		}
		val, err := evalCompositeElement(kv.Value, variables, elemType.Elem())
		if err != nil {
			return nil, err
		}
		m.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	return m.Interface(), nil
}

func evalCompositeElement(expr ast.Expr, variables Variables, elemType reflect.Type) (any, error) {
	switch elemType.Kind() {
	case reflect.Slice:
		return evalCompositeSlice(expr, variables, elemType)
	case reflect.Map:
		return evalCompositeMap(expr, variables, elemType)
	case reflect.Array:
		return evalCompositeArray(expr, variables, elemType)
	default:
		return evaltype(expr, variables, elemType)
	}
}

func evalCompositeLit(expr ast.Expr, variables Variables) (any, error) {
	composite := expr.(*ast.CompositeLit)
	t, err := compositeType(composite.Type)
	if err != nil {
		return nil, err
	}
	return evalCompositeElement(composite, variables, t)
}
