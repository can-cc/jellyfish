package logger

import (
	configs "jellyfish/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger(configs.LoggerConfig{
		Level: "info",
	})
	assert.Nil(t, err)
	assert.NotNil(t, sugaredLogger)
}
