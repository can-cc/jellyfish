package impl

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/juju/errors"
	"time"
)

func NewTacoRepository(dataSource *database.AppDataSource) repository.Repository {
	return &TacoRepositoryImpl{dataSource: dataSource}
}

type TacoRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func (t *TacoRepositoryImpl) List(userID string, filter taco.ListTacoFilter) ([]taco.Taco, error) {
	sql, _, err := buildListTacosSQL(userID, filter)
	if err != nil {
		return nil, errors.Trace(err)
	}
	rows, err := t.dataSource.RDS.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}
	tacos := make([]taco.Taco, 0)
	for rows.Next() {
		var t taco.Taco
		if err := rows.Scan(&t.Id, &t.CreatorId, &t.Content, &t.Detail, &t.Type, &t.Deadline, &t.Status, &t.BoxId, &t.CreatedAt, &t.UpdateAt); err != nil {
			return nil, errors.Trace(err)
		}
		tacos = append(tacos, t)
	}
	return tacos, nil
}

func (t *TacoRepositoryImpl) Save(taco *taco.Taco) (*string, error) {
	if taco.IsNew() {
		return t.insert(taco)
	}
	err := t.updateTaco(taco)
	return &taco.Id, err
}

func (t *TacoRepositoryImpl) FindById(tacoID string) (*taco.Taco, error) {
	sql, _, err := getGoquTacoSelection().Where(goqu.C("id").Eq(tacoID)).ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var ta taco.Taco
	err = t.dataSource.RDS.QueryRow(sql).Scan(&ta.Id, &ta.CreatorId, &ta.Content, &ta.Detail, &ta.Type, &ta.Deadline, &ta.Status, &ta.BoxId, &ta.CreatedAt, &ta.UpdateAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &ta, nil
}

func (t *TacoRepositoryImpl) insert(taco *taco.Taco) (*string, error) {
	sql, _, err := goqu.Insert("taco").Rows(
		goqu.Record{
			"content":    taco.Content,
			"creator_id": taco.CreatorId,
			"box_id":     taco.BoxId,
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

func (t *TacoRepositoryImpl) updateTaco(taco *taco.Taco) error {
	sql, _, err := goqu.Update("taco").Set(
		goqu.Record{
			"content":    taco.Content,
			"detail":     taco.Detail,
			"status":     taco.Status,
			"box_id":     taco.BoxId,
			"type":       taco.Type,
			"deadline":   taco.Deadline,
			"updated_at": time.Now(),
		},
	).Where(goqu.C("id").Eq(taco.Id)).ToSQL()
	if err != nil {
		return errors.Trace(err)
	}
	_, err = t.dataSource.RDS.Exec(sql)
	return err
}

func (t *TacoRepositoryImpl) Delete(tacoId string) error {
	sql, _, err := goqu.Update("taco").Set(
		goqu.Record{
			"deleted": 1,
		}).Where(goqu.C("id").Eq(tacoId)).ToSQL()
	if err != nil {
		return errors.Trace(err)
	}
	_, err = t.dataSource.RDS.Exec(sql)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func buildListTacosSQL(userID string, filter taco.ListTacoFilter) (sql string, params []interface{}, err error) {
	statuesFilters := []exp.Expression{goqu.C("creator_id").Eq(userID)}
	statuesFilters = append(statuesFilters, goqu.C("deleted").IsNull())
	if filter.Statues != nil {
		statuesFilters = append(statuesFilters, goqu.C("status").In(filter.Statues))
	}
	if filter.BoxId != nil {
		statuesFilters = append(statuesFilters, goqu.C("box_id").Eq(filter.BoxId))
	}
	if filter.Type != nil {
		statuesFilters = append(statuesFilters, goqu.C("type").Eq(filter.Type))
	}

	return getGoquTacoSelection().Where(
		statuesFilters...,
	).ToSQL()
}

func getGoquTacoSelection() *goqu.SelectDataset {
	return goqu.From("taco").Select(
		"id",
		"creator_id",
		database.TRIM("content"),
		database.TRIM("detail"),
		database.TRIM("type"),
		"deadline",
		database.TRIM("status"),
		database.TRIM("box_id"),
		"created_at",
		"updated_at",
	)
}
