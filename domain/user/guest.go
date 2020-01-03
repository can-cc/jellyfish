package user

import "os/user"

type Guest struct {
	Name        string
	Password    string
	IsCertified bool
}

// Transform : Passed identity authentication or join our application, turned into user
func (g *Guest) TransformAppUser() (*user.User, error) {
	return nil, nil
}
