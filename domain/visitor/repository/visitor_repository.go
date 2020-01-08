package repository

import "github.com/fwchen/jellyfish/domain/visitor"

type Repository interface {
	Save(visitor *visitor.Visitor, hash string) (string, error)
	FindUserPasswordHash(name string) (string, error)
}
