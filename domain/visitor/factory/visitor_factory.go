package factory

import "github.com/fwchen/jellyfish/domain/visitor"

func NewVisitor(name, password string) *visitor.Visitor {
	return &visitor.Visitor{
		Name:        name,
		Password:    password,
		IsCertified: false,
	}
}
