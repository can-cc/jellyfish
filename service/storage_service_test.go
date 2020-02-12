package service

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStorageService_MakeBucket(t *testing.T) {
	cfg, _ := configs.LoadConfig("../config/config.yaml")
	storageService := NewStorageService(&cfg.Storage)
	err := storageService.init()
	assert.Nil(t, err)
	bucketName := fmt.Sprintf("jellyfish-storage-test-%d", time.Now().Nanosecond())
	err = storageService.MakeBucket(bucketName)
	assert.Nil(t, err)
	exists, _ := storageService.client.BucketExists(bucketName)
	assert.True(t, exists)
	_ = storageService.client.RemoveBucket(bucketName)
}
