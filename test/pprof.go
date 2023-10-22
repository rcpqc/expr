package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/rcpqc/expr"
)

func main() {
	f, _ := os.Create("cpu.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// ex, _ := expr.Comp(`a==3`)
	// ex, _ := expr.Comp(`a==3 || g=="123"`)
	// ex, _ := expr.Comp(`!((a*b+c)/d>2)`)
	// ex, _ := expr.Comp(`e.x-f[2]<3.4`)
	ex, _ := expr.Comp(`!((a*b+c)/d>2 && e.x-f[2]<3.4) || g=="123"`) // 720
	vars := expr.Vars{"a": 2, "b": 3, "c": -1, "d": 1.5,
		"e": map[string]float32{"x": 4.2, "y": 3.6}, "f": []float64{1.2, 1.8, 2.4}, "g": "1243"}
	log.Print(expr.Eval(ex, vars))
	n := 10000000
	st := time.Now()
	for i := 0; i < n; i++ {
		expr.Eval(ex, vars)
	}
	log.Printf("%vns/op", float64(time.Since(st).Nanoseconds())/float64(n))
}
