package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"reflect"
	"strconv"
	"testing"
	"time"

	re "github.com/rcpqc/expr"
	"github.com/rcpqc/expr/types"
)

type Int32 int32

var m map[string]int64

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
		variables re.Vars
		nocomp    bool
		want      any
		err       error
	}{
		{
			expr:      `s == ""`,
			variables: re.Vars{"s": ""},
			want:      true,
		},
		{
			expr:      `a+b`,
			variables: re.Vars{"a": Int32(1231), "b": 565},
			want:      int64(1796),
		},
		{
			expr:      `uint32(a)`,
			variables: re.Vars{"a": 246},
			want:      uint32(246),
		},
		{
			expr:      `a[:-2]`,
			variables: re.Vars{"a": []int{1, 2, 3, 4}},
			want:      []int{1, 2},
		},
		{
			expr:      `a[2]`,
			variables: re.Vars{"a": []int{1, 2, 3, 4}},
			want:      3,
		},
		{
			expr:      `sfmt("%v_%v_%v",a+d,b,c)`,
			variables: re.Vars{"a": 123, "b": "fdf", "c": "5.6", "d": 434},
			want:      "557_fdf_5.6",
		},
		{
			expr: `(kkk.abc*2-1)/xyz.abc`,
			variables: re.Vars{"xyz": map[string]float64{"abc": 3.4}, "kkk": struct {
				ABC float64 `expr:"abc"`
			}{3.4}},
			want: 1.7058823529411764,
		},
		{
			expr:      `x == slen(y)`,
			variables: re.Vars{"x": 3, "y": "ggg"},
			want:      true,
		},
		{
			expr:      `b-a>0.7 || !c && a<2<<2 || xxx(a,b)<=(5-2)`,
			variables: re.Vars{"a": 2, "b": 2.51, "c": true, "xxx": func(a float64, b float64) float64 { return a + b - 2 }},
			want:      true,
		},
		{
			expr:      `a=='1' && b=="xyz"`,
			variables: re.Vars{"a": 49, "b": "xyz"},
			want:      true,
		},
		{
			expr:      `a / b + c / b`,
			variables: re.Vars{"a": 1, "b": 0, "c": true},
			err:       fmt.Errorf("binary(1:5) integer divide by zero"),
		},
		{
			expr:      `len(a) + len(b) + len(c) - cap(d) + len(a[0]) + cap(a)`,
			variables: re.Vars{"a": "abcde", "b": []int{1, 2, 3}, "c": map[string]int{"xx": 1, "yy": 3}, "d": make([]int, 0, 9)},
			want:      int64(1),
		},
		{
			expr:      `float32(a.b)`,
			variables: re.Vars{"a": map[string]int32{"c": 1234}},
			want:      float32(0),
		},
		{
			expr:      `a/-float64(b)`,
			variables: re.Vars{"a": float32(123), "b": 321},
			want:      123 / -321.0,
		},
		{
			expr:      `a["b"]+2`,
			variables: re.Vars{"a": nil},
			err:       fmt.Errorf("index(1:6) illegal kind(invalid)"),
		},
		{
			expr:      `a+2`,
			variables: re.Vars{"a": nil},
			err:       fmt.Errorf("binary(1:3) illegal expr(<nil>+int64)"),
		},
		{
			expr:      `!a`,
			variables: re.Vars{"a": nil},
			err:       fmt.Errorf("unary(1:2) illegal expr(!<nil>)"),
		},
		{
			expr:      `o.foo(a,2,6)+o.x+o.sum(1.2,a,b,-1)`,
			variables: re.Vars{"o": &S1{X: 8, a: 2.0}, "a": 1, "b": 4},
			want:      9.800000190734863,
		},
		{
			expr:      `int(a)+int8(a)+int16(a)+int32(a)+int64(a)+uint(b)+uint8(b)+uint16(b)+uint32(b)+uint64(b)`,
			variables: re.Vars{"a": -124234, "b": 4232},
			want:      int64(-348874),
		},
		{
			expr:      `a+b`,
			variables: re.Vars{"a": 3},
			err:       fmt.Errorf("ident(3:3) unknown name(b)"),
		},
		{
			expr:      `sfmt("%s_%s_%s",hex(md5(a)),hex(sha1(b)),hex(sha256(c))[-10:-1])`,
			variables: re.Vars{"a": "hello", "b": "world", "c": "!"},
			want:      "5d41402abc4b2a76b9719d911017c592_7c211433f02071597741e6ff5a8ea34789abbf43_2c3ba43b6",
		},
		{
			expr:      `itos(a)+utos(b)`,
			variables: re.Vars{"a": 123, "b": 4567},
			want:      "1234567",
		},
		{
			expr:      `a[4]+a[-2]+b["a"]+c[int32(3)]`,
			variables: re.Vars{"a": []string{"1", "2", "3", "4", "5"}, "b": map[string]string{"d": "xx"}, "c": map[int32]string{2: "fsd", 3: "ggg"}},
			want:      "54ggg",
		},
		{
			expr:      `c[3]`,
			variables: re.Vars{"c": map[int32]string{2: "fsd", 3: "ggg"}},
			err:       fmt.Errorf("index(1:4) map[int32]string can't index by key(int64)"),
		},
		{
			expr:      `a[-10]`,
			variables: re.Vars{"a": []string{"1", "2", "3", "4", "5"}},
			err:       fmt.Errorf("index(1:6) out of range index(-5) for len(5)"),
		},
		{
			expr:      `(a[b])`,
			variables: re.Vars{"a": []string{"1", "2", "3", "4", "5"}, "b": "1"},
			err:       fmt.Errorf("index(2:5) index must be an integer"),
		},
		{
			expr:      `a[b]+a[b.x]`,
			variables: re.Vars{"a": map[string]int{"x": 1, "y": 2, "z": 3}, "b": nil},
			err:       fmt.Errorf("selector(8:10) illegal kind(invalid)"),
		},
		{
			expr:      `tprs(tfmt(tnow(),layout),layout)==time(tnow()).unix()`,
			variables: re.Vars{"layout": time.RFC3339},
			want:      true,
		},
		{
			expr:      `max(a,b)+min(a,c)+sin(a)+cos(b)+tan(b*c)+exp(a-b)+log2(abs(b*c))`,
			variables: re.Vars{"a": 1.23, "b": 4.26, "c": -2.55},
			want:      -1.7936578069369178,
		},
		{
			expr:      `(a>0)*2.3+(a<=0)*ceil(b)+log(sigmoid(c))+sqrt(abs(tanh(c)))`,
			variables: re.Vars{"a": 1.23, "b": 4.26, "c": -2.55},
			want:      0.6687384992239993,
		},
		{
			expr:      `floor(stof(a))+round(stof(b))+stoi(c)`,
			variables: re.Vars{"a": "34.3", "b": "4.76", "c": "-2"},
			want:      37.0,
		},
		{
			expr:      `sfind(ports,split(ipport,":")[1])==12`,
			variables: re.Vars{"ipport": "192.168.1.1:4536", "ports": "3389,445,21,4536,22,5543"},
			want:      true,
		},
		{
			expr:      `stou(str(round(stof(sjoin(a,".")))))`,
			variables: re.Vars{"a": []string{"12", "34"}},
			want:      uint64(12),
		},
		{
			expr:      `pow(a,n)+b-c`,
			variables: re.Vars{"a": 3, "n": 2.0, "c": false, "b": true},
			want:      float64(10),
		},
		{
			expr:      `d/(log10(a)*b*c)`,
			variables: re.Vars{"a": 10, "b": true, "c": 5, "d": true},
			nocomp:    true,
			want:      float64(0.2),
		},
		{
			expr:      `a || 1/0`,
			variables: re.Vars{"a": true},
			want:      true,
		},
		{
			expr:      `a && 1/0`,
			variables: re.Vars{"a": false},
			want:      false,
		},
		{
			expr:      `b+i+b-f+(b+f)`,
			variables: re.Vars{"b": true, "f": 2.3, "i": 8},
			want:      float64(11),
		},
		{
			expr:      `(b-(f<i))+(b-i)+(i-b)`,
			variables: re.Vars{"b": true, "f": 2.3, "i": 8},
			want:      int64(0),
		},
		{
			expr:      `(b*(i>f))-(i*b)/(i*i)-(i*f)`,
			variables: re.Vars{"b": true, "f": 2.3, "i": 8},
			want:      -17.4,
		},
		{
			expr:      `a[b]`,
			variables: re.Vars{"a": []int{1, 2, 4, 3}, "b": 4},
			err:       fmt.Errorf("index(1:4) out of range index(4) for len(4)"),
		},
		{
			expr:      `a[b%c]`,
			variables: re.Vars{"a": []int{1, 2, 4}, "b": 4, "c": 0},
			err:       fmt.Errorf("binary(3:5) integer divide by zero"),
		},
		{
			expr:      `(a.b)[:4]`,
			variables: re.Vars{"a": 0},
			nocomp:    true,
			err:       fmt.Errorf("selector(2:4) sel(b) not found"),
		},
		{
			expr:      `a[:4]`,
			variables: re.Vars{"a": 0},
			err:       fmt.Errorf("slice(1:5) illegal kind(int)"),
		},
		{
			expr:      `a[b/c:3]`,
			variables: re.Vars{"a": []int{}, "b": 1, "c": 0},
			err:       fmt.Errorf(`binary(3:5) integer divide by zero`),
		},
		{
			expr:      `a["2":3]`,
			variables: re.Vars{"a": []int{}},
			err:       fmt.Errorf(`slice(1:8) low index must be an integer`),
		},
		{
			expr:      `a[2:"3"]`,
			variables: re.Vars{"a": []int{}},
			err:       fmt.Errorf(`slice(1:8) high index must be an integer`),
		},
		{
			expr:      `a[uint32(2):df]`,
			variables: re.Vars{"a": []int{}},
			err:       fmt.Errorf("ident(13:14) unknown name(df)"),
		},
		{
			expr:      `a[b:6]`,
			variables: re.Vars{"a": []int{1, 2, 3, 4}, "b": 2},
			err:       fmt.Errorf("slice(1:6) out of range index(2:6) for len(4)"),
		},
		{
			expr:      `a+123.45i`,
			variables: re.Vars{"a": 123},
			err:       fmt.Errorf("basic_lit(3:9) illegal kind(IMAG)"),
		},
		{
			expr:      `fn(1,2,3)`,
			variables: re.Vars{"fn": 123},
			err:       fmt.Errorf("call(1:9) not a func"),
		},
		{
			expr:      `s.xyz()`,
			variables: re.Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:7) output parameters want(1) got(2)"),
		},
		{
			expr:      `s.foo(1,2)`,
			variables: re.Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:10) input parameters want(3) got(2)"),
		},
		{
			expr:      `s.foo(1,2,a)`,
			variables: re.Vars{"s": &S1{}, "a": nil},
			want:      3,
		},
		{
			expr:      `s.foo(1,2,a)`,
			variables: re.Vars{"s": &S1{}, "a": "3"},
			err:       fmt.Errorf("call(1:12) arg2 want int got string"),
		},
		{
			expr:      `s.sum()`,
			variables: re.Vars{"s": &S1{}},
			err:       fmt.Errorf("call(1:7) input parameters want(>=1) got(0)"),
		},
		{
			expr:      `s.sum(a,1,2,3)`,
			variables: re.Vars{"s": &S1{}, "a": nil},
			want:      float32(0.0),
		},
		{
			expr:      `s.sum(a,1,2,3)`,
			variables: re.Vars{"s": &S1{}, "a": "3"},
			err:       fmt.Errorf("call(1:14) arg0 want float32 got string"),
		},
		{
			expr:      `s.sum(a,1,b)`,
			variables: re.Vars{"s": &S1{}, "a": 1.2, "b": nil},
			want:      float32(1.2),
		},
		{
			expr:      `s.sum(a,1,"32")`,
			variables: re.Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("call(1:15) arg2 want int got string"),
		},
		{
			expr:      `s.sum(a,1,b[0])`,
			variables: re.Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("ident(11:11) unknown name(b)"),
		},
		{
			expr:      `o.bar(a,1,b[0])`,
			variables: re.Vars{"s": &S1{}, "a": 1.2},
			err:       fmt.Errorf("ident(1:1) unknown name(o)"),
		},
		{
			expr:      `s.a`,
			variables: re.Vars{"s": &S1{}},
			err:       fmt.Errorf("selector(1:3) sel(a) not found"),
		},
		{
			expr:      `m.(a)`,
			variables: re.Vars{"m": map[int]string{1: "fsd", 2: "fdsf"}},
			err:       fmt.Errorf("(1:5) unsupported expression type(*ast.TypeAssertExpr)"),
		},
		{
			expr:      `m.a`,
			variables: re.Vars{"m": map[int]string{1: "fsd", 2: "fdsf"}},
			err:       fmt.Errorf("selector(1:3) key of map must be string"),
		},
		{
			expr:      `(b==b) + (b==i) + (b==f) + (i==b) + (i==i) + (i==f) + (f==b) + (f==i) + (f==f) + (s=="123")`,
			variables: re.Vars{"b": true, "f": 1.0, "i": 1, "s": "123"},
			want:      int64(10),
		},
		{
			expr:      `(b!=b) + (b!=i) + (b!=f) + (i!=b) + (i!=i) + (i!=f) + (f!=b) + (f!=i) + (f!=f) + (s!="123")`,
			variables: re.Vars{"b": true, "f": 1.5, "i": 1, "s": "456"},
			want:      int64(5),
		},
		{
			expr:      `(i<i) + (i<f) +(f<i) + (f<f) + (s<"123")`,
			variables: re.Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(2),
		},
		{
			expr:      `(i<=i) + (i<=f) +(f<=i) + (f<=f) + (s<="0123")`,
			variables: re.Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(4),
		},
		{
			expr:      `(i>i) + (i>f) +(f>i) + (f>f) + (s>"123")`,
			variables: re.Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(1),
		},
		{
			expr:      `(i>=i) + (i>=f) +(f>=i) + (f>=f) + (s>="-23")`,
			variables: re.Vars{"f": 1.5, "i": 1, "s": "0123"},
			want:      int64(4),
		},
		{
			expr:      `s.it.sub(6,3)`,
			variables: re.Vars{"s": &S1{IT: S2(5)}},
			want:      15,
		},
		{
			expr:      `s.it.sub(6,3)`,
			variables: re.Vars{"s": (*S1)(nil)},
			err:       fmt.Errorf("selector(1:4) nil value"),
		},
		{
			expr:      `!false && true`,
			variables: re.Vars{},
			want:      true,
		},
		{
			expr:      `x+`,
			variables: re.Vars{},
			err:       fmt.Errorf(`1:3: expected operand, found 'EOF'`),
		},
		{
			expr:      `s.a`,
			variables: re.Vars{"s": map[string]int32{"a": 43}},
			nocomp:    true,
			want:      int32(43),
		},
		{
			expr:      `slower(a)+supper(b)`,
			variables: re.Vars{"a": "XY123gg", "b": "65yGdz#4"},
			want:      "xy123gg65YGDZ#4",
		},
		{
			expr:      `clamp(a, 0, 1)+clamp(b, 0, 1)+clamp(c, 0, 1)`,
			variables: re.Vars{"a": 0.65, "b": -0.2, "c": 1.3},
			want:      1.65,
		},
		{
			expr:      `[]string{"a","b","c"}`,
			variables: re.Vars{},
			want:      []string{"a", "b", "c"},
		},
		{
			expr:      `[4]int16{-3,4,5}`,
			variables: re.Vars{},
			want:      [4]int16{-3, 4, 5},
		},
		{
			expr:      `map[string]int{"a":0,"b":435,"c":12.0}`,
			variables: re.Vars{},
			want:      map[string]int{"a": 0, "b": 435, "c": 12},
		},
		{
			expr:      `map[string]bool{"true": a>0, "false": a<0}`,
			variables: re.Vars{"a": 3},
			want:      map[string]bool{"true": true, "false": false},
		},
		{
			expr:      `[][]float64{{1,2,3},{4,5,6}}`,
			variables: re.Vars{},
			want:      [][]float64{{1, 2, 3}, {4, 5, 6}},
		},
		{
			expr:      `[2]map[string]map[string]int{{"c":{"a":1,"b":2}}}`,
			variables: re.Vars{},
			want:      [2]map[string]map[string]int{{"c": {"a": 1, "b": 2}}},
		},
		{
			expr:      `[a]int{1,2,3}`,
			variables: re.Vars{},
			err:       fmt.Errorf("ident(2:2) illegal expression for array's length"),
		},
		{
			expr:      `[0]int{1,2,3}`,
			variables: re.Vars{},
			err:       fmt.Errorf("basic_lit(8:8) out of bounds(>=0) for array"),
		},
		{
			expr:      `[2]map[any]chan int{}`,
			variables: re.Vars{},
			err:       fmt.Errorf("(12:19) illegal composite type"),
		},
		{
			expr:      `map[uintptr]string{}`,
			variables: re.Vars{},
			err:       fmt.Errorf("ident(5:11) unsupported type(uintptr)"),
		},
		{
			expr:      `map[string]int{"a":1,"b"}`,
			variables: re.Vars{},
			err:       fmt.Errorf("basic_lit(22:24) expect key:value as an element of map"),
		},
		{
			expr:      `[2][]float64{{1,2,3},2}`,
			variables: re.Vars{},
			err:       fmt.Errorf("basic_lit(22:22) int64(2) can't convert to type([]float64)"),
		},
		{
			expr:      `[]map[int]any{{2:"fds"},"fsjklf"}`,
			variables: re.Vars{},
			err:       fmt.Errorf("basic_lit(25:32) string(fsjklf) can't convert to type(map[int]interface {})"),
		},
		{
			expr:      `map[[2]string]int{{"a","b"}:1,x:2}`,
			variables: re.Vars{"x": [2]string{"c", "d"}},
			want:      map[[2]string]int{{"a", "b"}: 1, {"c", "d"}: 2},
		},
		{
			expr:      `map[[2]string]int{{"a","b"}:1,x:2}`,
			variables: re.Vars{"x": 432},
			err:       fmt.Errorf("ident(31:31) int(432) can't convert to type([2]string)"),
		},
		{
			expr:      `(b+u)+(u+b)+(u+i)+(i+u)+(u+f)+(f+u)+(u+u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": -3, "f": 3.64},
			want:      59.28,
		},
		{
			expr:      `(b-u)+(u-b)+(u-i)+(i-u)+(u-f)+(f-u)+(u-u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": -3, "f": 3.64},
			want:      0.0,
		},
		{
			expr:      `(b*u)+(u*b)+(u*i)+(i*u)+(u*f)+(f*u)+(u*u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": -3, "f": 3.5},
			want:      70.0,
		},
		{
			expr:      `(b/u)+(u/i)+(i/u)+(u/f)+(f/u)+(u/u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": -3, "f": 3.64},
			want:      1.443076923076923,
		},
		{
			expr:      `(i%u)+(u%i)+(u%u)+(i%i)`,
			variables: re.Vars{"u": uint(7), "i": -3},
			want:      int64(-2),
		},
		{
			expr:      `(i&u)+(u&i)+(u&u)+(i|u)+(u|i)+(u|u)+(i^u)+(u^i)+(u^u)`,
			variables: re.Vars{"u": uint(5), "i": 3},
			want:      int64(38),
		},
		{
			expr:      `(b==u)+(u==b)+(u==i)+(i==u)+(u==f)+(f==u)+(u==u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": 1, "f": 7.0},
			want:      int64(3),
		},
		{
			expr:      `(b!=u)+(u!=b)+(u!=i)+(i!=u)+(u!=f)+(f!=u)+(u!=u)`,
			variables: re.Vars{"b": true, "u": uint(7), "i": 1, "f": 7.0},
			want:      int64(4),
		},
		{
			expr:      `(u>i)+(i>=u)+(u<f)+(f<=u)+(u>u)`,
			variables: re.Vars{"u": uint(7), "i": 1, "f": 7.0},
			want:      int64(2),
		},
		{
			expr:      `(u<=i)+(i>u)+(u>=f)+(f<u)+(u<=u)`,
			variables: re.Vars{"u": uint(7), "i": 1, "f": 7.0},
			want:      int64(2),
		},
		{
			expr:      `(u<i)+(i<=u)+(u>f)+(f>=u)+(u<u)`,
			variables: re.Vars{"u": uint(7), "i": 1, "f": 7.0},
			want:      int64(2),
		},
		{
			expr:      `(u>=i)+(i<u)+(u<=f)+(f>u)+(u>=u)`,
			variables: re.Vars{"u": uint(7), "i": 1, "f": 7.0},
			want:      int64(4),
		},
		{
			expr:      `(i<<u)+(u<<i)+(u<<u)+(i>>u)+(u>>i)+(u>>u)+(-u)`,
			variables: re.Vars{"u": uint(5), "i": 3},
			want:      int64(291),
		},
		{
			expr:      `p==nil && q!=nil && !(nil==q) && nil!=fn`,
			variables: re.Vars{"p": (*int)(nil), "q": &m, "fn": func() {}},
			want:      true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var expr ast.Expr
			var err error
			if tt.nocomp {
				expr, err = parser.ParseExpr(tt.expr)
			} else {
				expr, err = re.Comp(tt.expr)
			}
			var got any
			if err == nil {
				got, err = re.Eval(expr, tt.variables)
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
		variables re.Vars
		typ       reflect.Type
		want      any
		err       error
	}{
		{
			expr:      `a & b + b | c - (b >> 2) + (a << 1)`,
			variables: re.Vars{"a": 4234, "b": 12222, "c": 983},
			typ:       types.Byte,
			want:      byte(4),
		},
		{
			expr:      `a - a/b + a*c`,
			variables: re.Vars{"a": true, "b": 2.0, "c": 3},
			typ:       types.Float32,
			want:      float32(3.5),
		},
		{
			expr:      `a != (b % c)`,
			variables: re.Vars{"a": 3, "b": 9, "c": 6},
			typ:       types.Bool,
			want:      false,
		},
		{
			expr:      `-(a/b)`,
			variables: re.Vars{"a": true, "b": 0},
			typ:       types.Int,
			err:       fmt.Errorf("binary(3:5) integer divide by zero"),
		},
		{
			expr:      `a+b`,
			variables: re.Vars{"a": true, "b": false},
			typ:       nil,
			want:      int64(1),
		},
		{
			expr:      `a+b`,
			variables: re.Vars{"a": true, "b": false},
			typ:       reflect.TypeOf([]int32{}),
			want:      int64(1),
			err:       fmt.Errorf("binary(1:3) int64(1) can't convert to type([]int32)"),
		},
		{
			expr:      `m["c"]`,
			variables: re.Vars{"m": map[string]any{"a": 1, "b": "2"}},
			typ:       types.Float32,
			err:       fmt.Errorf("index(1:6) nil can't convert to type(float32)"),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expr, err := re.Comp(tt.expr)
			if err != nil {
				t.Errorf("expr(%s) err: %v", tt.expr, err)
				return
			}
			got, err := re.EvalType(expr, tt.variables, tt.typ)

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
		variables    re.Vars
		defaultValue any
		want         any
	}{
		{
			expr:         `a^b>0 == c>=0`,
			variables:    re.Vars{"a": 12, "b": 23, "c": 3},
			defaultValue: int8(0),
			want:         int8(0),
		},
		{
			expr:         `a/b+b/a`,
			variables:    re.Vars{"a": 12, "b": 3.0},
			defaultValue: float32(0),
			want:         float32(4.25),
		},
		{
			expr:         `a/b+b/c+b%c`,
			variables:    re.Vars{"a": true, "b": 15, "c": 6},
			defaultValue: float32(0),
			want:         float32(5),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expr, err := re.Comp(tt.expr)
			if err != nil {
				t.Errorf("expr(%s) err: %v", tt.expr, err)
				return
			}
			got := re.EvalOr(expr, tt.variables, tt.defaultValue)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n[EXPR  ] %s\n[RESULT] [%v]%v\n[EXPECT] [%v]%v",
					tt.expr, reflect.TypeOf(got), got, reflect.TypeOf(tt.want), tt.want)
			}
		})
	}
}
