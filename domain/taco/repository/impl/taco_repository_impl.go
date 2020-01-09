package impl

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/juju/errors"
)

type TacoRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func buildListTacosSQL(userID string, filter repository.ListTacoFilter) (sql string, params []interface{}, err error) {
	statuesFilters := []exp.Expression{goqu.C("creator_id").Eq(userID)}
	if filter.Statues != nil {
		statuesFilters = append(statuesFilters, goqu.C("status").In(filter.Statues))
	}

	return goqu.From("test").Select("id", "TRIM(content)", "TRIM(detail)", "type", "deadline", "status", "created_at", "updated_at").Where(
		statuesFilters...,
	).ToSQL()
}

func (t *TacoRepositoryImpl) ListTacos(userID string, filter repository.ListTacoFilter) ([]taco.Taco, error) {
	sql, _, err := buildListTacosSQL(userID, filter)
	if err != nil {
		return nil, errors.Trace(err)
	}
	rows, err := t.dataSource.RDS.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}
	var tacos []taco.Taco
	for rows.Next() {
		var taco taco.Taco
		if err := rows.Scan(&taco.ID, &taco.Content, &taco.Detail, &taco.Type, &taco.Deadline, &taco.Status, &taco.CreatedAt, &taco.UpdateAt); err != nil {
			return nil, errors.Trace(err)
		}
		tacos = append(tacos, taco)
	}
	return tacos, nil
}
