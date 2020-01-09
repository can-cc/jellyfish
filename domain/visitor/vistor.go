package visitor

import (
	"errors"
	"github.com/fwchen/jellyfish/domain/user"
)

type Visitor struct {
	Name        string
	Password    string
	IsCertified bool
}

// Transform : Passed identity authentication or join our application, turned into user
func (g *Visitor) TransformAppUser() (*user.AppUser, error) {
	if !g.IsCertified {
		return nil, errors.New("transform app user from uot certified visitor")
	}
	panic("implement me")
}
