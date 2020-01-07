package handlers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"

	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
)

// GetMD5Hash :
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func saveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fm, err
}
