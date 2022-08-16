package expr

import (
	"go/ast"

	"github.com/rcpqc/expr/eval"
)

func Eval(expr ast.Expr, variables map[string]interface{}) (interface{}, error) {
	return eval.EvalExpr(expr, variables)
}
