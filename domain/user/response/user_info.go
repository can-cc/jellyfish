package response

import "github.com/fwchen/jellyfish/domain/user"

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func TransformToUserInfo(user *user.AppUser) *UserInfo {
	return &UserInfo{
		ID:       *user.ID,
		Username: *user.Username,
	}
}
