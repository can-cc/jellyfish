package impl

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/util"
	"github.com/juju/errors"
)

func NewTacoRepository(dataSource *database.AppDataSource) repository.Repository {
	return &TacoRepositoryImpl{dataSource: dataSource}
}

type TacoRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func (t *TacoRepositoryImpl) List(userID string, filter taco.TacoFilter) ([]taco.Taco, error) {
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
		if err := rows.Scan(&t.Id, &t.CreatorId, &t.Content, &t.Detail, &t.Type, &t.Deadline, &t.Status, &t.BoxId, &t.Order, &t.CreatedAt, &t.UpdateAt); err != nil {
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

// TODO batch save
func (t *TacoRepositoryImpl) SaveList(tacos []taco.Taco) error {
	for i := 0; i < len(tacos); i++ {
		_, err := t.Save(&tacos[i])
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (t *TacoRepositoryImpl) FindById(tacoId string) (*taco.Taco, error) {
	sql, _, err := getGoquTacoSelection().Where(goqu.C("id").Eq(tacoId)).ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var ta taco.Taco
	err = t.dataSource.RDS.QueryRow(sql).Scan(&ta.Id, &ta.CreatorId, &ta.Content, &ta.Detail, &ta.Type, &ta.Deadline, &ta.Status, &ta.BoxId, &ta.Order, &ta.CreatedAt, &ta.UpdateAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &ta, nil
}

func (t *TacoRepositoryImpl) insert(taco *taco.Taco) (*string, error) {
	sql, _, err := goqu.Insert("taco").Rows(
		goqu.Record{
			"content":     taco.Content,
			"creator_id":  taco.CreatorId,
			"box_id":      taco.BoxId,
			"detail":      taco.Detail,
			"status":      taco.Status,
			"type":        taco.Type,
			"important":   util.ConvertBool2Int(taco.IsImportant),
			"order_index": taco.Order,
			"deadline":    taco.Deadline,
		},
	).Returning("id").ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var id string
	err = t.dataSource.RDS.QueryRow(sql).Scan(&id)
	return &id, err
}

func (t *TacoRepositoryImpl) MaxOrderByCreatorId(creatorId string) (*float64, error) {
	sql, _, err := goqu.From("taco").Select(
		goqu.MAX("order_index").As("order"),
	).Where(goqu.C("creator_id").Eq(creatorId)).ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return t.findOrder(sql)
}

func (t *TacoRepositoryImpl) MaxOrderByBoxId(boxId string) (*float64, error) {
	sql, _, err := goqu.From("taco").Select(
		goqu.MAX("order_index").As("order"),
	).Where(goqu.C("box_id").Eq(boxId)).ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return t.findOrder(sql)
}

func (t *TacoRepositoryImpl) findOrder(sql string) (*float64, error) {
	var order *float64
	err := t.dataSource.RDS.QueryRow(sql).Scan(&order)
	if order == nil {
		order = util.PointerFloat64(float64(0))
	}
	return order, err
}

func (t *TacoRepositoryImpl) updateTaco(taco *taco.Taco) error {
	sql, _, err := goqu.Update("taco").Set(
		goqu.Record{
			"content":     taco.Content,
			"detail":      taco.Detail,
			"status":      taco.Status,
			"box_id":      taco.BoxId,
			"type":        taco.Type,
			"important":   util.BoolToInt(taco.IsImportant),
			"order_index": taco.Order,
			"deadline":    taco.Deadline,
			"updated_at":  time.Now(),
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

func buildListTacosSQL(userID string, filter taco.TacoFilter) (sql string, params []interface{}, err error) {
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
	if filter.Important == true {
		statuesFilters = append(statuesFilters, goqu.C("important").Eq(1))
	}
	if filter.Scheduled == true {
		statuesFilters = append(statuesFilters, goqu.C("deadline").IsNotNull())
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
		"order_index",
		"created_at",
		"updated_at",
	)
}
