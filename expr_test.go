package expr

import (
	"go/parser"
	"reflect"
	"strconv"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr      string
		variables map[string]interface{}
		want      interface{}
		wantErr   bool
	}{
		{
			expr:      `uint32(a)`,
			variables: map[string]interface{}{"a": 246},
			want:      uint32(246),
		},
		{
			expr:      `a[:-2]`,
			variables: map[string]interface{}{"a": []int{1, 2, 3, 4}},
			want:      []int{1, 2},
		},
		{
			expr:      `a[2]`,
			variables: map[string]interface{}{"a": []int{1, 2, 3, 4}},
			want:      3,
		},
		{
			expr:      `sfmt("%v_%v_%v",a+d,b,c)`,
			variables: map[string]interface{}{"a": 123, "b": "fdf", "c": "5.6", "d": 434},
			want:      "557_fdf_5.6",
		},
		{
			expr: `(kkk.abc*2-1)/2==2.9`,
			variables: map[string]interface{}{"xyz": map[string]float64{"abc": 3.4}, "kkk": struct {
				ABC float64 `expr:"abc"`
			}{3.4}},
			want: true,
		},
		{
			expr:      `x == slen(y)`,
			variables: map[string]interface{}{"x": 3, "y": "ggg"},
			want:      true,
		},
		{
			expr:      `b-a>0.7 || !c && a<2<<2 || xxx(a,b)<=(5-2)`,
			variables: map[string]interface{}{"a": 2, "b": 2.51, "c": true, "xxx": func(a float64, b float64) float64 { return a + b - 2 }},
			want:      true,
		},
		{
			expr:      `a=='1' && b=="xyz"`,
			variables: map[string]interface{}{"a": 49, "b": "xyz"},
			want:      true,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
