package eval

import (
	"go/ast"
	"go/token"
	"strconv"

	"github.com/rcpqc/expr/errs"
)

func evalBasicLit(expr ast.Expr, variables Variables) (any, error) {
	basic := expr.(*ast.BasicLit)
	switch basic.Kind {
	case token.INT:
		return strconv.ParseInt(basic.Value, 10, 64)
	case token.FLOAT:
		return strconv.ParseFloat(basic.Value, 64)
	case token.STRING:
		return basic.Value[1 : len(basic.Value)-1], nil
	case token.CHAR:
		return int64(basic.Value[1]), nil
	default:
		return nil, errs.Newf(basic, "illegal kind(%s)", basic.Kind.String())
	}
}
