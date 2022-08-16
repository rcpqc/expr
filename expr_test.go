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
				t.Errorf("evalUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalUnary() = %v, want %v", got, tt.want)
			}
		})
	}
}
