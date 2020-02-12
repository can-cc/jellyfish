package service

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/logger"
	"github.com/juju/errors"
	"github.com/minio/minio-go/v6"
	"io"
)

type StorageService struct {
	configure *configs.StorageConfig
	client    *minio.Client
}

func NewStorageService(config *configs.StorageConfig) *StorageService {
	return &StorageService{
		configure: config,
	}
}

func (ss *StorageService) init() error {
	minioClient, err := minio.New(ss.configure.Endpoint, ss.configure.AccessKeyID, ss.configure.SecretAccessKeyID, ss.configure.UseSSL)
	if err != nil {
		return errors.Trace(err)
	}
	ss.client = minioClient
	return nil
}

func (ss *StorageService) MakeBucket(bucketName string) error {
	err := ss.client.MakeBucket(bucketName, ss.configure.Location)
	if err != nil {
		_, errBucketExists := ss.client.BucketExists(bucketName)
		if errBucketExists != nil {
			return errors.Trace(err)
		}
	} else {
		logger.L.Infow(fmt.Sprintf("Successfully created storage bucket %s\n", bucketName))
	}
	return nil
}

func (ss *StorageService) PutObject(bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) error {
	uploadedBytes, err := ss.client.PutObject(bucketName, objectName, reader, objectSize, opts)
	if err != nil {
		return errors.Trace(err)
	}
	logger.L.Infow(fmt.Sprintf("Successfully uploaded %s of size %d\n", objectName, uploadedBytes))
	return nil
}
