package service

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/juju/errors"
	"github.com/minio/minio-go/v6"
	"io"
)

type ImageStorageService struct {
	storageService *StorageService
	bucketName     string
}

func NewImageStorageService(bucketName string, storageService *StorageService) (*ImageStorageService, error) {
	s := &ImageStorageService{bucketName: bucketName, storageService: storageService}
	err := s.storageService.MakeBucket(bucketName)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return s, nil
}

func (i *ImageStorageService) SaveBase64Image(code string) (string, error) {
	h := md5.New()
	io.WriteString(h, code)
	fileName := fmt.Sprintf("%x", h.Sum(nil))
	dec, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", errors.Trace(err)
	}
	reader := bytes.NewReader(dec)
	contentType := "image/png"
	opts := minio.PutObjectOptions{ContentType: contentType}
	err = i.storageService.PutObject(i.bucketName, fileName, reader, reader.Size(), opts)
	if err != nil {
		return "", errors.Trace(err)
	}
	return fileName, nil
}
