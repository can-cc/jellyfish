package database

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func newIdentifierFunc(name string, col interface{}) exp.SQLFunctionExpression {
	if s, ok := col.(string); ok {
		col = goqu.I(s)
	}
	return goqu.Func(name, col)
}

func TRIM(col interface{}) exp.SQLFunctionExpression { return newIdentifierFunc("TRIM", col) }
