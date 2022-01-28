package response

import (
	"jellyfish/domain/user"
)

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func TransformToUserInfo(user *user.AppUser) *UserInfo {
	avatarFileName := ""
	if user.GetAvatar() != nil {
		avatarFileName = user.GetAvatar().FileName
	}
	return &UserInfo{
		ID:       *user.ID,
		Username: *user.Username,
		Avatar:   avatarFileName,
	}
}
