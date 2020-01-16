package impl

import (
	"database/sql"
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
	return u.updateUser(user)
}

func (u *userRepositoryImpl) Has(username string) (bool, error) {
	var exist bool
	err := u.dataSource.RDS.QueryRow(`SELECT EXISTS(SELECT id FROM app_user WHERE username = $1)`, username).Scan(&exist)
	return exist, err
}

func (u *userRepositoryImpl) FindByID(userID string) (*user.AppUser, error) {
	var username, hash string
	var avatar sql.NullString
	var createdAt, updatedAt time.Time
	err := u.dataSource.RDS.QueryRow(`SELECT TRIM(username), TRIM(hash), avatar, created_at, updated_at FROM app_user WHERE id = $1`, userID).Scan(&username, &hash, &avatar, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return factory.NewUser(userID, username, hash, avatar.String, createdAt, updatedAt), nil
}

func (u *userRepositoryImpl) updateUser(user *user.AppUser) error {
	if user.ID == nil {
		return errors.New("Update user without userID")
	}
	_, err := u.dataSource.RDS.Exec(
		`UPDATE app_user SET username = $1, hash = $2, avatar = $3, updated_at = now()
                WHERE id = $4`,
		user.Username, user.GetPasswordHash(), user.GetAvatar().Code, user.ID,
	)
	return err
}
