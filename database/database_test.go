package database

import (
	configs "jellyfish/config"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetDatabase(t *testing.T) {
	cfg, _ := configs.LoadConfig("../config/config.yaml")
	appDatabase, err := GetDatabase(cfg.DataSource)
	assert.Equal(t, nil, err)
	err = appDatabase.RDS.Close()
	assert.Equal(t, nil, err)
}
