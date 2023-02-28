package expr

import (
	"go/parser"
	"reflect"
	"strconv"
	"testing"
)

type Int32 int32

func TestEval(t *testing.T) {
	tests := []struct {
		expr      string
		variables Vars
		want      interface{}
		err       string
	}{
		{
			expr:      `s == ""`,
			variables: Vars{"s": ""},
			want:      true,
		},
		{
			expr:      `a+b`,
			variables: Vars{"a": Int32(1231), "b": 565},
			err:       "[binary] illegal expr (expr.Int32 + int)",
		},
		{
			expr:      `uint32(a)`,
			variables: Vars{"a": 246},
			want:      uint32(246),
		},
		{
			expr:      `a[:-2]`,
			variables: Vars{"a": []int{1, 2, 3, 4}},
			want:      []int{1, 2},
		},
		{
			expr:      `a[2]`,
			variables: Vars{"a": []int{1, 2, 3, 4}},
			want:      3,
		},
		{
			expr:      `sfmt("%v_%v_%v",a+d,b,c)`,
			variables: Vars{"a": 123, "b": "fdf", "c": "5.6", "d": 434},
			want:      "557_fdf_5.6",
		},
		{
			expr: `(kkk.abc*2-1)/2==2.9`,
			variables: Vars{"xyz": map[string]float64{"abc": 3.4}, "kkk": struct {
				ABC float64 `expr:"abc"`
			}{3.4}},
			want: true,
		},
		{
			expr:      `x == slen(y)`,
			variables: Vars{"x": 3, "y": "ggg"},
			want:      true,
		},
		{
			expr:      `b-a>0.7 || !c && a<2<<2 || xxx(a,b)<=(5-2)`,
			variables: Vars{"a": 2, "b": 2.51, "c": true, "xxx": func(a float64, b float64) float64 { return a + b - 2 }},
			want:      true,
		},
		{
			expr:      `a=='1' && b=="xyz"`,
			variables: Vars{"a": 49, "b": "xyz"},
			want:      true,
		},
		{
			expr:      `a / b + c / b`,
			variables: Vars{"a": 1, "b": 0, "c": true},
			err:       "integer divide by zero",
		},
		{
			expr:      `len(a) + len(b) + len(c) - cap(d) + len(a[0])`,
			variables: Vars{"a": "abcde", "b": []int{1, 2, 3}, "c": map[string]int{"xx": 1, "yy": 3}, "d": make([]int, 0, 9)},
			want:      int64(1),
		},
		{
			expr:      `a.b`,
			variables: Vars{"a": map[string]int32{"c": 1234}},
			want:      int32(0),
		},
		{
			expr:      `has("xxx",a) && has("1231",b) && has(float32(1.23),c) && !has(2,d)`,
			variables: Vars{"a": map[string]int32{"xxx": 1234}, "b": []string{"1231", "fjls", "32e", "bfd"}, "c": [2]float32{1.23, 54.45}, "d": map[int]int{1: 2, 3: 4}},
			want:      true,
		},
		{
			expr:      `has("xxx",a)`,
			variables: Vars{"a": nil},
			err:       "[call] arg1 is invalid",
		},
		{
			expr:      `a["b"]+2`,
			variables: Vars{"a": nil},
			err:       "[index] illegal kind(invalid)",
		},
		{
			expr:      `a+2`,
			variables: Vars{"a": nil},
			err:       "[binary] illegal expr (<nil> + int64)",
		},
		{
			expr:      `!a`,
			variables: Vars{"a": nil},
			err:       "[unary] illegal expr (! <nil>)",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expr, err := parser.ParseExpr(tt.expr)
			if err != nil {
				t.Errorf("expr(%s) err: %v", tt.expr, err)
				return
			}
			got, err := Eval(expr, tt.variables)
			if (err != nil && err.Error() != tt.err) || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GOT(%v, %v) !=  WANT(%v, %v)", got, err, tt.want, tt.err)
			}
		})
	}
}
