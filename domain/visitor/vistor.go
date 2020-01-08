package visitor

import "github.com/fwchen/jellyfish/domain/user"

type Visitor struct {
	Name        string
	Password    string
	IsCertified bool
}

// Transform : Passed identity authentication or join our application, turned into user
func (g *Visitor) TransformAppUser() (*user.AppUser, error) {
	return nil, nil
}
