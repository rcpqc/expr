package eval

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

func evalBasicLit(basic *ast.BasicLit, variables Variables) (interface{}, error) {
	switch basic.Kind {
	case token.INT:
		return strconv.ParseInt(basic.Value, 10, 64)
	case token.FLOAT:
		return strconv.ParseFloat(basic.Value, 64)
	case token.STRING:
		return basic.Value[1 : len(basic.Value)-1], nil
	case token.CHAR:
		return int64(basic.Value[1]), nil
	}
	return nil, fmt.Errorf("[basiclit] illegal kind (%s)", basic.Kind.String())
}
