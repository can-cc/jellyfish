package repository

import "github.com/fwchen/jellyfish/domain/visitor"

type Repository interface {
	Save(visitor *visitor.Visitor, hash string) (string, error)
	FindUserIDAndHash(name string) (string, string, error)
}
