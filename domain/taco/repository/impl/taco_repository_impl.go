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

func NewTacoRepository(dataSource *database.AppDataSource) repository.Repository {
	return &TacoRepositoryImpl{dataSource: dataSource}
}

type TacoRepositoryImpl struct {
	dataSource *database.AppDataSource
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
		var t taco.Taco
		if err := rows.Scan(&t.ID, &t.Content, &t.Detail, &t.Type, &t.Deadline, &t.Status, &t.CreatedAt, &t.UpdateAt); err != nil {
			return nil, errors.Trace(err)
		}
		tacos = append(tacos, t)
	}
	return tacos, nil
}

func (t *TacoRepositoryImpl) InsertTaco(taco *taco.Taco) (*string, error) {
	sql, _, err := goqu.Insert("todo").Rows(
		goqu.Record{
			"content":    taco.Content,
			"creator_id": taco.CreatorID,
			"detail":     taco.Detail,
			"status":     taco.Status,
			"type":       taco.Type,
			"deadline":   taco.Deadline,
		},
	).Returning("id").ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var id string
	err = t.dataSource.RDS.QueryRow(sql).Scan(&id)
	return &id, err
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
