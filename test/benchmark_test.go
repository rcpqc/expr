package main

import (
	"testing"

	"github.com/rcpqc/expr"
)

var vars = expr.Vars{"a": 2, "b": 3, "c": -1, "d": 1.5,
	"e": struct {
		X float32
		Y float32
	}{4.2, 3.6}, "f": []float64{1.2, 1.8, 2.4}, "g": "1243"}

func BenchmarkExpr0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = vars["a"].(int) == 3
	}
}

func BenchmarkExpr1(b *testing.B) {
	ex, _ := expr.Comp(`a==3`) // 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(ex, vars)
	}
}

func BenchmarkExpr2(b *testing.B) {
	ex, _ := expr.Comp(`a==3 || g=="123"`) // 111
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(ex, vars)
	}
}
func BenchmarkExpr3(b *testing.B) {
	ex, _ := expr.Comp(`!((a*b+c)/d>2)`) // 220
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(ex, vars)
	}
}
func BenchmarkExpr4(b *testing.B) {
	ex, _ := expr.Comp(`e.x-f[2]<3.4`) // 324
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(ex, vars)
	}
}

func BenchmarkExpr5(b *testing.B) {
	ex, _ := expr.Comp(`!((a*b+c)/d>2 && e.x-f[2]<3.4) || g=="123"`) // 688
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(ex, vars)
	}
}
