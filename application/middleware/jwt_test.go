package middleware

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestSignedToken(t *testing.T) {
	token, _ := SignedToken(SignData{ID: "u1"}, "jellyfish_secret")
	ts := strings.Split(token, ".")
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", ts[0])
}
