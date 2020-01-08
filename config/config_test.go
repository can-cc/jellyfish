package configs

import (
	"fmt"
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
	os.Setenv("JFISH_LOGGER.LEVEL", "info")
	os.Setenv("JFISH_APPLICATION.ADDR", "www.jellyfish.com:80")
	fmt.Print(os.Getenv("jfish_logger.level"))
	cfg, _ := LoadConfig("config.yaml")
	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "www.jellyfish.com:80", cfg.Application.Addr)
}
