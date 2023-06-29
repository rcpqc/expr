package expr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rcpqc/expr/types"
)

type Int32 int32

type S1 struct {
	X  int
	Y  string `expr:"-"`
	a  float32
	IT I1
}

func (o *S1) Foo(a, b, c int) int {
	return a + b - o.bar(c)
}
func (o *S1) Sum(f float32, elems ...int) float32 {
	sum := 0
	for _, e := range elems {
		sum += e
	}
	return f * float32(sum)
}
func (o *S1) bar(x int) int {
	return x
}
func (o *S1) XYZ() (int, error) {
	return 12, nil
}

type I1 interface {
	Sub(a, b int) int
}
type S2 int

func (o S2) Sub(a, b int) int {
	return int(o) * (a - b)
}
func TestEval(t *testing.T) {
	tests := []struct {
		expr      string
		variables Vars
		nocomp    bool
		want      interface{}
		err       error
	}{
		{
			expr:      `s == ""`,
			variables: Vars{"s": ""},
			want:      true,
		},
		{
			expr:      `a+b`,
			variables: Vars{"a": Int32(1231), "b": 565},
			err:       fmt.Errorf("binary(1:3) illegal expr(expr.Int32+int64)"),
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
			expr: `(kkk.abc*2-1)/xyz.abc`,
			variables: Vars{"xyz": map[string]float64{"abc": 3.4}, "kkk": struct {
				ABC float64 `expr:"abc"`
			}{3.4}},
			want: 1.7058823529411764,
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
			err:       fmt.Errorf("binary(1:5) integer divide by zero"),
		},
		{
			expr:      `len(a) + len(b) + len(c) - cap(d) + len(a[0]) + cap(a)`,
			variables: Vars{"a": "abcde", "b": []int{1, 2, 3}, "c": map[string]int{"xx": 1, "yy": 3}, "d": make([]int, 0, 9)},
			want:      int64(1),
		},
		{
			expr:      `float32(a.b)`,
			variables: Vars{"a": map[string]int32{"c": 1234}},
			want:      float32(0),
		},
		{
			expr:      `a/-float64(b)`,
			variables: Vars{"a": float32(123), "b": 321},
			want:      123 / -321.0,
		},
		{
			expr:      `a["b"]+2`,
			variables: Vars{"a": nil},
			err:       fmt.Errorf("index(1:6) illegal kind(invalid)"),
		},
		{
			expr:      `a+2`,
			variables: Vars{"a": nil},
			err:       fmt.Errorf("binary(1:3) illegal expr(<nil>+int64)"),
		},
		{
			expr:      `!a`,
			variables: Vars{"a": nil},
			err:       fmt.Errorf("unary(1:2) illegal expr(!<nil>)"),
		},
		{
			expr:      `o.foo(a,2,6)+o.x+o.sum(1.2,a,b,-1)`,
			variables: Vars{"o": &S1{X: 8, a: 2.0}, "a": 1, "b": 4},
			want:      9.800000190734863,
		},
		{
			expr:      `int(a)+int8(a)+int16(a)+int32(a)+int64(a)+uint(b)+uint8(b)+uint16(b)+uint32(b)+uint64(b)`,
			variables: Vars{"a": -124234, "b": 4232},
			want:      int64(-348874),
		},
		{
			expr:      `a+b`,
			variables: Vars{"a": 3},
			err:       fmt.Errorf("ident(3:3) unknown name(b)"),
		},
		{
			expr:      `sfmt("%s_%s_%s",hex(md5(a)),hex(sha1(b)),hex(sha256(c))[-10:-1])`,
			variables: Vars{"a": "hello", "b": "world", "c": "!"},
			want:      "5d41402abc4b2a76b9719d911017c592_7c211433f02071597741e6ff5a8ea34789abbf43_2c3ba43b6",
		},
		{
			expr:      `itos(a)+utos(b)`,
			variables: Vars{"a": 123, "b": 4567},
			want:      "1234567",
		},
		{
			expr:      `a[4]+a[-2]+b["a"]+c[int32(3)]`,
			variables: Vars{"a": []string{"1", "2", "3", "4", "5"}, "b": map[string]string{"d": "xx"}, "c": map[int32]string{2: "fsd", 3: "ggg"}},
			want:      "54ggg",
		},
		{
			expr:      `c[3]`,
			variables: Vars{"c": map[int32]string{2: "fsd", 3: "ggg"}},
			err:       fmt.Errorf("index(1:4) map[int32]string can't index by key(int64)"),
		},
		{
			expr:      `a[-10]`,
			variables: Vars{"a": []string{"1", "2", "3", "4", "5"}},
			err:       fmt.Errorf("index(1:6) out of range index(-5) for len(5)"),
		},
		{
			expr:      `(a[b])`,
			variables: Vars{"a": []string{"1", "2", "3", "4", "5"}, "b": "1"},
			err:       fmt.Errorf("index(2:5) index must be an integer"),
		},
		{
			expr:      `a[b]+a[b.x]`,
			variables: Vars{"a": map[string]int{"x": 1, "y": 2, "z": 3}, "b": nil},
			err:       fmt.Errorf("selector(8:10) illegal kind(invalid)"),
		},
		{
			expr:      `tprs(tfmt(tnow(),layout),layout)==time(tnow()).unix()`,
			variables: Vars{"layout": time.RFC3339},
			want:      true,
		},
		{
			expr:      `max(a,b)+min(a,c)+sin(a)+cos(b)+tan(b*c)+exp(a-b)+log2(abs(b*c))`,
			variables: Vars{"a": 1.23, "b": 4.26, "c": -2.55},
			want:      -1.7936578069369178,
		},
		{
			expr:      `(a>0)*2.3+(a<=0)*ceil(b)+log(sigmoid(c))+sqrt(abs(tanh(c)))`,
			variables: Vars{"a": 1.23, "b": 4.26, "c": -2.55},
			want:      0.6687384992239993,
		},
		{
			expr:      `floor(stof(a))+round(stof(b))+stoi(c)`,
			variables: Vars{"a": "34.3", "b": "4.76", "c": "-2"},
			want:      37.0,
		},
		{
			expr:      `sfind(ports,split(ipport,":")[1])==12`,
			variables: Vars{"ipport": "192.168.1.1:4536", "ports": "3389,445,21,4536,22,5543"},
			want:      true,
		},
		{
			expr:      `stou(str(round(stof(sjoin(a,".")))))`,
			variables: Vars{"a": []string{"12", "34"}},
			want:      uint64(12),
		},
		{
			expr:      `pow(a,n)+b-c`,
			variables: Vars{"a": 3, "n": 2.0, "c": false, "b": true},
			want:      float64(10),
		},
		{
			expr:      `d/(log10(a)*b*c)`,
			variables: Vars{"a": 10, "b": true, "c": 5, "d": true},
			want:      float64(0.2),
		},
		{
			expr:      `a || 1/0`,
			variables: Vars{"a": true},
			want:      true,
		},
		{
			expr:      `a && 1/0`,
			variables: Vars{"a": false},
			want:      false,
		},
		{
			expr:      `b+i+b-f+(b+f)`,
			variables: Vars{"b": true, "f": 2.3, "i": 8},
			want:      float64(11),
		},
		{
			expr:      `(b-(f<i))+(b-i)+(i-b)`,
			variables: Vars{"b": true, "f": 2.3, "i": 8},
			want:      int64(0),
		},
		{
			expr:      `(b*(i>f))-(i*b)/(i*i)-(i*f)`,
			variables: Vars{"b": true, "f": 2.3, "i": 8},
			want:      -17.4,
		},
		{
			expr:      `a[b%c]`,
			variables: Vars{"a": []int{1, 2, 4}, "b": 4, "c": 0},
			err:       fmt.Errorf("binary(3:5) integer divide by zero"),
		},
		{
			expr:      `(a.b)[:4]`,
			variables: Vars{"a": 0},
			err:       fmt.Errorf("selector(2:4) sel(b) not found"),
		},
		{
			expr:      `a[:4]`,
			variables: Vars{"a": 0},
			err:       fmt.Errorf("slice(1:5) illegal kind(int)"),
		},
		{
			expr:      `a[b/c:3]`,
			variables: Vars{"a": []int{}, "b": 1, "c": 0},
			err:       fmt.Errorf(`binary(3:5) integer divide by zero`),
		},
		{
			expr:      `a["2":3]`,
			variables: Vars{"a": []int{}},
			err:       fmt.Errorf(`slice(1:8) low index must be an integer`),
		},
		{
			expr:      `a[2:"3"]`,
			variables: Vars{"a": []int{}},
			err:       fmt.Errorf(`slice(1:8) high index must be an integer`),
		},
		{
			expr:      `a[uint32(2):df]`,
			variables: Vars{"a": []int{}},
			err:       fmt.Errorf("ident(13:14) unknown name(df)"),
		},
		{
			expr:      `a[b:6]`,
			variables: Vars{"a": []int{1, 2, 3, 4}, "b": 2},
			err:       fmt.Errorf("slice(1:6) out of range index(2:6) for len(4)"),
		},
		{
			expr:      `a+123.45i`,
			variables: Vars{"a": 123},
			err:       fmt.Errorf("basic_lit(3:9) illegal kind(IMAG)"),
		},
		{
			expr:      `fn(1,2,3)`,
			variables: Vars{"fn": 123},
			err:       fmt.Errorf("call(1:9) not a func"),
		},
		{
			expr:      `s.xyz()`,
			variables: Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:7) output parameters want(1) got(2)"),
		},
		{
			expr:      `s.foo(1,2)`,
			variables: Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:10) input parameters want(3) got(2)"),
		},
		{
			expr:      `s.foo(1,2,a)`,
			variables: Vars{"s": &S1{}, "a": nil},
			err:       fmt.Errorf("call(1:12) arg2 is invalid"),
		},
		{
			expr:      `s.foo(1,2,a)`,
			variables: Vars{"s": &S1{}, "a": "3"},
			err:       fmt.Errorf("call(1:12) arg2 want int got string"),
		},
		{
			expr:      `s.sum()`,
			variables: Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:7) input parameters want(>=1) got(0)"),
		},
		{
			expr:      `s.sum(a,1,2,3)`,
			variables: Vars{"s": &S1{}, "a": nil},
			err:       fmt.Errorf("call(1:14) arg0 is invalid"),
		},
		{
			expr:      `s.sum(a,1,2,3)`,
			variables: Vars{"s": &S1{}, "a": "3"},
			err:       fmt.Errorf("call(1:14) arg0 want float32 got string"),
		},
		{
			expr:      `s.sum(a,1,b)`,
			variables: Vars{"s": &S1{}, "a": 1.2, "b": nil},
			err:       fmt.Errorf("call(1:12) arg2 is invalid"),
		},
		{
			expr:      `s.sum(a,1,"32")`,
			variables: Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("call(1:15) arg2 want int got string"),
		},
		{
			expr:      `s.sum(a,1,b[0])`,
			variables: Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("ident(11:11) unknown name(b)"),
		},
		{
			expr:      `o.bar(a,1,b[0])`,
			variables: Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("ident(1:1) unknown name(o)"),
		},
		{
			expr:      `s.a`,
			variables: Vars{"s": &S1{}},
			err:       fmt.Errorf("selector(1:3) sel(a) not found"),
		},
		{
			expr:      `m.(a)`,
			variables: Vars{"m": map[int]string{1: "fsd", 2: "fdsf"}},
			err:       fmt.Errorf("unsupported expression type(*ast.TypeAssertExpr)"),
		},
		{
			expr:      `m.a`,
			variables: Vars{"m": map[int]string{1: "fsd", 2: "fdsf"}},
			err:       fmt.Errorf("selector(1:3) key of map must be string"),
		},
		{
			expr:      `(b==b) + (b==i) + (b==f) + (i==b) + (i==i) + (i==f) + (f==b) + (f==i) + (f==f) + (s=="123")`,
			variables: Vars{"b": true, "f": 1.0, "i": 1, "s": "123"},
			want:      int64(10),
		},
		{
			expr:      `(b!=b) + (b!=i) + (b!=f) + (i!=b) + (i!=i) + (i!=f) + (f!=b) + (f!=i) + (f!=f) + (s!="123")`,
			variables: Vars{"b": true, "f": 1.5, "i": 1, "s": "456"},
			want:      int64(5),
		},
		{
			expr:      `(i<i) + (i<f) +(f<i) + (f<f) + (s<"123")`,
			variables: Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(2),
		},
		{
			expr:      `(i<=i) + (i<=f) +(f<=i) + (f<=f) + (s<="0123")`,
			variables: Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(4),
		},
		{
			expr:      `(i>i) + (i>f) +(f>i) + (f>f) + (s>"123")`,
			variables: Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(1),
		},
		{
			expr:      `(i>=i) + (i>=f) +(f>=i) + (f>=f) + (s>="-23")`,
			variables: Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(4),
		},
		{
			expr:      `s.it.sub(6,3)`,
			variables: Vars{"s": &S1{IT: S2(5)}},
			want:      15,
		},
		{
			expr:      `s.it.sub(6,3)`,
			variables: Vars{"s": (*S1)(nil)},
			err:       fmt.Errorf("selector(1:4) nil value"),
		},
		{
			expr:      `!false && true`,
			variables: Vars{},
			want:      true,
		},
		{
			expr:      `x+`,
			variables: Vars{},
			err:       fmt.Errorf(`1:3: expected operand, found 'EOF'`),
		},
		{
			expr:      `s.a`,
			variables: Vars{"s": map[string]int32{"a": 43}},
			nocomp:    true,
			want:      int32(43),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var expr ast.Expr
			var err error
			if tt.nocomp {
				expr, err = parser.ParseExpr(tt.expr)
			} else {
				expr, err = Comp(tt.expr)
			}
			var got interface{}
			if err == nil {
				got, err = Eval(expr, tt.variables)
			}
			if (err == nil && tt.err != nil) ||
				(err != nil && tt.err == nil) ||
				(err != nil && tt.err != nil && err.Error() != tt.err.Error()) ||
				!reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n[EXPR  ] %s\n[RESULT] [%v]%v, [err]%v\n[EXPECT] [%v]%v, [err]%v",
					tt.expr, reflect.TypeOf(got), got, err, reflect.TypeOf(tt.want), tt.want, tt.err)
			}
		})
	}
}

func TestEvalType(t *testing.T) {
	tests := []struct {
		expr      string
		variables Vars
		typ       reflect.Type
		want      interface{}
		err       error
	}{
		{
			expr:      `a & b + b | c - (b >> 2) + (a << 1)`,
			variables: Vars{"a": 4234, "b": 12222, "c": 983},
			typ:       types.Byte,
			want:      byte(4),
		},
		{
			expr:      `a - a/b + a*c`,
			variables: Vars{"a": true, "b": 2.0, "c": 3},
			typ:       types.Float32,
			want:      float32(3.5),
		},
		{
			expr:      `a != (b % c)`,
			variables: Vars{"a": 3, "b": 9, "c": 6},
			typ:       types.Bool,
			want:      false,
		},
		{
			expr:      `-(a/b)`,
			variables: Vars{"a": true, "b": 0},
			typ:       types.Int,
			err:       fmt.Errorf("binary(3:5) integer divide by zero"),
		},
		{
			expr:      `a+b`,
			variables: Vars{"a": true, "b": false},
			typ:       nil,
			want:      int64(1),
		},
		{
			expr:      `a+b`,
			variables: Vars{"a": true, "b": false},
			typ:       reflect.TypeOf([]int32{}),
			want:      int64(1),
			err:       fmt.Errorf("binary(1:3) int64(1) can't convert to type([]int32)"),
		},
		{
			expr:      `a`,
			variables: Vars{"a": nil},
			typ:       types.Float32,
			want:      float32(0.0),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expr, err := Comp(tt.expr)
			if err != nil {
				t.Errorf("expr(%s) err: %v", tt.expr, err)
				return
			}
			got, err := EvalType(expr, tt.variables, tt.typ)

			if (err == nil && tt.err != nil) ||
				(err != nil && tt.err == nil) ||
				(err != nil && tt.err != nil && err.Error() != tt.err.Error()) ||
				!reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n[EXPR  ] %s\n[RESULT] [%v]%v, [err]%v\n[EXPECT] [%v]%v, [err]%v",
					tt.expr, reflect.TypeOf(got), got, err, reflect.TypeOf(tt.want), tt.want, tt.err)
			}
		})
	}
}

func TestEvalOr(t *testing.T) {
	tests := []struct {
		expr         string
		variables    Vars
		defaultValue interface{}
		want         interface{}
	}{
		{
			expr:         `a^b>0 == c>=0`,
			variables:    Vars{"a": 12, "b": 23, "c": 3},
			defaultValue: int8(0),
			want:         int8(0),
		},
		{
			expr:         `a/b+b/a`,
			variables:    Vars{"a": 12, "b": 3.0},
			defaultValue: float32(0),
			want:         float32(4.25),
		},
		{
			expr:         `a/b+b/c+b%c`,
			variables:    Vars{"a": true, "b": 15, "c": 6},
			defaultValue: float32(0),
			want:         float32(5),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expr, err := Comp(tt.expr)
			if err != nil {
				t.Errorf("expr(%s) err: %v", tt.expr, err)
				return
			}
			got := EvalOr(expr, tt.variables, tt.defaultValue)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n[EXPR  ] %s\n[RESULT] [%v]%v\n[EXPECT] [%v]%v",
					tt.expr, reflect.TypeOf(got), got, reflect.TypeOf(tt.want), tt.want)
			}
		})
	}
}
