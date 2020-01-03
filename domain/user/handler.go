package user

import "github.com/fwchen/jellyfish/domain/user/repository"

type Handler struct {
	userRepo repository.UserRepository
}
