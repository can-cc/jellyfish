package user

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAppUser_SetAvatar(t *testing.T) {
	var user AppUser
	user.SetAvatar("mock_base64")
	assert.Equal(t, "mock_base64", user.avatar.Code)
}

func TestAppUser_SetPasswordHash(t *testing.T) {
	var user AppUser
	user.SetPasswordHash("***")
	assert.Equal(t, "***", *user.passwordHash)
}
