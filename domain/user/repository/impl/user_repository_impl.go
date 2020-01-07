package impl

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/user"
	"github.com/fwchen/jellyfish/domain/user/factory"
	"github.com/fwchen/jellyfish/domain/user/repository"
	"github.com/juju/errors"
	"time"
)

func NewUserRepository(dataSource *database.AppDataSource) repository.Repository {
	return &userRepositoryImpl{dataSource: dataSource}
}

type userRepositoryImpl struct {
	dataSource *database.AppDataSource
}

func (u *userRepositoryImpl) Save(user *user.AppUser) error {
	if user.ID != nil {
		return u.updateUser(user)
	}
	return u.insertUser(user)
}

func (u *userRepositoryImpl) Has(username string) (bool, error) {
	var exist bool
	err := u.dataSource.RDS.QueryRow(`SELECT EXISTS(SELECT id FROM app_user WHERE username = $1)`, username).Scan(&exist)
	return exist, err
}

func (u *userRepositoryImpl) FindByID(userID string) (*user.AppUser, error) {
	var username, hash, avatar *string
	var createdAt, updatedAt *time.Time
	err := u.dataSource.RDS.QueryRow(`SELECT id, username, hash, avatar, created_at, updated_at FROM app_user WHERE id = $1`, userID).Scan(username, hash, createdAt, updatedAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return factory.NewUser(userID, *username, *hash, *avatar, *createdAt, *updatedAt), nil
}

func (u *userRepositoryImpl) insertUser(user *user.AppUser) error {
	sqlStatement := `
		INSERT INTO app_user (username, hash, created_at) 
		VALUES ($1, $2, now()) RETURNING id`
	return u.dataSource.RDS.QueryRow(sqlStatement, user.Username, user.GetPasswordHash()).Scan(&user.ID)
}

func (u *userRepositoryImpl) updateUser(user *user.AppUser) error {
	_, err := u.dataSource.RDS.Exec(
		`UPDATE app_user SET username = $1, hash = $2, avatar = $3, updated_at = now()
                WHERE id = $5`,
		user.Username, user.GetPasswordHash(), user.GetAvatar(),
	)
	return err
}
