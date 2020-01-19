package tacoBoxImpl

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/repository"
	"github.com/juju/errors"
	"time"
)

type TacoBoxRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func NewTacoBoxRepositoryImpl(dataSource *database.AppDataSource) repository.TacoBoxRepository {
	return &TacoBoxRepositoryImpl{dataSource: dataSource}
}

func (t *TacoBoxRepositoryImpl) SaveTacoBox(box *taco_box.TacoBox) (*taco_box.TacoBoxID, error) {
	if box.IsNew() {
		return t.insertTacoBox(box)
	}
	err := t.updateTacoBox(box)
	return &box.ID, err
}

func (t *TacoBoxRepositoryImpl) ListTacoBoxes(userID string) ([]taco_box.TacoBox, error) {
	sql, _, err := buildListTacoBoxesSQL(userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	rows, err := t.dataSource.RDS.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}
	tacoBoxes := make([]taco_box.TacoBox, 0)
	for rows.Next() {
		var tb taco_box.TacoBox
		if err := rows.Scan(&tb.ID, &tb.Name, &tb.Icon, &tb.CreatorID, &tb.CreatedAt, &tb.UpdatedAt); err != nil {
			return nil, errors.Trace(err)
		}
		tacoBoxes = append(tacoBoxes, tb)
	}
	return tacoBoxes, nil
}

func (t *TacoBoxRepositoryImpl) FindTacoBox(boxID string) (*taco_box.TacoBox, error) {
	sql, _, err := getGoquTacoSelection().Where(goqu.C("id").Eq(boxID)).ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var tb taco_box.TacoBox
	err = t.dataSource.RDS.QueryRow(sql).Scan(&tb.ID, &tb.Name, &tb.Icon, &tb.CreatorID, &tb.CreatedAt, &tb.UpdatedAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &tb, nil
}

func (t *TacoBoxRepositoryImpl) insertTacoBox(box *taco_box.TacoBox) (*taco_box.TacoBoxID, error) {
	sql, _, err := goqu.Insert("taco_box").Rows(
		goqu.Record{
			"name":       box.Name,
			"icon":       box.Icon,
			"creator_id": box.CreatorID,
		},
	).Returning("id").ToSQL()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var id taco_box.TacoBoxID
	err = t.dataSource.RDS.QueryRow(sql).Scan(&id)
	return &id, err
}

func (t *TacoBoxRepositoryImpl) updateTacoBox(box *taco_box.TacoBox) error {
	sql, _, err := goqu.Update("taco_box").Set(
		goqu.Record{
			"name":      box.Name,
			"icon":      box.Icon,
			"updatedAt": time.Now(),
		},
	).ToSQL()
	if err != nil {
		return errors.Trace(err)
	}
	_, err = t.dataSource.RDS.Exec(sql)
	return err
}

func buildListTacoBoxesSQL(userID string) (sql string, params []interface{}, err error) {
	statuesFilters := []exp.Expression{goqu.C("creator_id").Eq(userID)}
	return getGoquTacoSelection().Where(
		statuesFilters...,
	).ToSQL()
}

func getGoquTacoSelection() *goqu.SelectDataset {
	return goqu.From("taco_box").Select(
		"id",
		"name",
		"icon",
		"creator_id",
		"created_at",
		"updated_at",
	)
}
