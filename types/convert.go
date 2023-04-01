package types

import "reflect"

var cvrs [MaxKinds]func(rv reflect.Value) (interface{}, reflect.Kind)

func init() {
	cvrs[reflect.Bool] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Bool(), reflect.Bool }

	cvrs[reflect.Int] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Int(), reflect.Int64 }
	cvrs[reflect.Int8] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Int(), reflect.Int64 }
	cvrs[reflect.Int16] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Int(), reflect.Int64 }
	cvrs[reflect.Int32] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Int(), reflect.Int64 }
	cvrs[reflect.Int64] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Int(), reflect.Int64 }

	cvrs[reflect.Uint] = func(rv reflect.Value) (interface{}, reflect.Kind) { return int64(rv.Uint()), reflect.Int64 }
	cvrs[reflect.Uint8] = func(rv reflect.Value) (interface{}, reflect.Kind) { return int64(rv.Uint()), reflect.Int64 }
	cvrs[reflect.Uint16] = func(rv reflect.Value) (interface{}, reflect.Kind) { return int64(rv.Uint()), reflect.Int64 }
	cvrs[reflect.Uint32] = func(rv reflect.Value) (interface{}, reflect.Kind) { return int64(rv.Uint()), reflect.Int64 }
	cvrs[reflect.Uint64] = func(rv reflect.Value) (interface{}, reflect.Kind) { return int64(rv.Uint()), reflect.Int64 }

	cvrs[reflect.Float32] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Float(), reflect.Float64 }
	cvrs[reflect.Float64] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.Float(), reflect.Float64 }

	cvrs[reflect.String] = func(rv reflect.Value) (interface{}, reflect.Kind) { return rv.String(), reflect.String }
}

// Convert convert bool,int,uint,float,string
func Convert(x interface{}) (interface{}, reflect.Kind) {
	rv := reflect.ValueOf(x)
	if cvr := cvrs[rv.Kind()]; cvr != nil {
		return cvr(rv)
	}
	return x, rv.Kind()
}
