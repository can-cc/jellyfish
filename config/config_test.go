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
	assert.Equal(t, []string{"stdout"}, cfg.Logger.ErrorOutputPaths)
}

func TestLoadConfig_WithEnv(t *testing.T) {
	os.Setenv("jfish_logger.level", "info")
	os.Setenv("jfish_application.addr", "www.jellyfish.com:80")
	cfg, _ := LoadConfig("config.yaml")
	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "www.jellyfish.com:80", cfg.Application.Addr)
}
