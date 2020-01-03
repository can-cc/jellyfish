package impl

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/user"
	"github.com/juju/errors"
)

type UserRepositoryImpl struct {
	dataSource database.AppDataSource
}

func (u *UserRepositoryImpl) InsertUser(user *user.AppUser, hash string) error {
	sqlStatement := `INSERT INTO users (username, hash, created_at) VALUES ($1, $2, now()) RETURNING id`
	return u.dataSource.RDS.QueryRow(sqlStatement, user.Username, hash).Scan(&user.ID)
}

func (u *UserRepositoryImpl) HasUser(username string) (bool, error) {
	var exist bool
	err := u.dataSource.RDS.QueryRow(`SELECT EXISTS(SELECT id FROM user WHERE username = $1)`, username).Scan(&exist)
	return exist, err
}

func (u *UserRepositoryImpl) FindUser(username string) (*user.AppUser, error) {
	var user user.AppUser
	err := u.dataSource.RDS.QueryRow(`SELECT id, username, avatar_id, created_at, updated_at FROM user WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.AvatarID, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &user, nil
}

func (u *UserRepositoryImpl) FindUserHash(username string) (string, error) {
	return "", nil
}
