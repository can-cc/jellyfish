package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("config.yaml")
	assert.Equal(t, nil, err)
	assert.Equal(t, "debug", cfg.Logger.Level)
	assert.Equal(t, "0.0.0.0:8180", cfg.Application.Addr)
	assert.Equal(t, []string{"stdout"}, cfg.Logger.OutputPaths)
}

func TestLoadConfig_WithEnv(t *testing.T) {
	_ = os.Setenv("JFISH_LOGGER_LEVEL", "info")
	_ = os.Setenv("JFISH_APPLICATION_ADDR", "www.jellyfish.com:80")
	cfg, _ := LoadConfig("config.yaml")
	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "www.jellyfish.com:80", cfg.Application.Addr)
}
