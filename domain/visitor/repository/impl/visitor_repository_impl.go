package impl

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/visitor"
	"github.com/fwchen/jellyfish/domain/visitor/repository"
)

func NewVisitorRepository(dataSource *database.AppDataSource) repository.Repository {
	return &visitorRepositoryImpl{dataSource: dataSource}
}

type visitorRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func (v *visitorRepositoryImpl) Save(visitor *visitor.Visitor, hash string) (string, error) {
	var id string
	sqlStatement := `
		INSERT INTO app_user (username, hash, created_at)
		VALUES ($1, $2, now()) RETURNING id`
	err := v.dataSource.RDS.QueryRow(sqlStatement, visitor.Name, hash).Scan(&id)
	return id, err
}

func (v *visitorRepositoryImpl) FindUserPasswordHash(name string) (string, error) {
	panic("implement me")
}
