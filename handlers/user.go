package handlers

import (
	"database/sql"

	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"

	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

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
	// _, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fm, err
}

func PostAvatarByBase64(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		request := new(struct {
			Avatar string `json:"avatar"`
		})

		c.Bind(&request)

		avatardir := viper.GetString("avatardir")
		fileNameHash := GetMD5Hash(request.Avatar)

		fm, err := saveImageToDisk(avatardir+fileNameHash, request.Avatar)
		if err != nil {
			panic(err)
		}

		sql := "UPDATE users set avatar = ? where id = ?"

		stmt, err2 := db.Prepare(sql)

		if err2 != nil {
			panic(err2)
		}

		defer stmt.Close()
		_, err3 := stmt.Exec(fileNameHash+"."+fm, userId)
		if err3 != nil {
			panic(err3)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"avatar": avatardir + fileNameHash + "." + fm,
		})

	}
}

func GetUserInfo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		sql := "SELECT username, avatar FROM users where id = ?"
		row := db.QueryRow(sql, userId)

		userInfo := new(struct {
			Username string `json:"username"`
			Avatar   string `json:"avatar"`
		})
		avatardir := viper.GetString("avatardir")

		err := row.Scan(&userInfo.Username, &userInfo.Avatar)
		if err != nil {
			panic(err)
		}
		userInfo.Avatar = avatardir + userInfo.Avatar
		return c.JSON(http.StatusOK, userInfo)
	}
}

func PostAvatar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		form, err := c.MultipartForm()

		if err != nil {
			return err
		}
		files := form.File["avatar"]

		if len(files) != 1 {
			return c.NoContent(http.StatusBadRequest)
		}

		for _, file := range files {
			// Source
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			avatardir := viper.GetString("avatardir")

			fileNameHash := GetMD5Hash(file.Filename)

			// Destination
			dst, err := os.Create(avatardir + fileNameHash)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			sql := "UPDATE users set avatar = ? where id = ?"

			stmt, err2 := db.Prepare(sql)

			if err2 != nil {
				panic(err2)
			}

			defer stmt.Close()
			_, err3 := stmt.Exec(fileNameHash, userId)
			if err3 != nil {
				panic(err3)
			}

			return c.JSON(http.StatusOK, map[string]string{
				"avatar": avatardir + fileNameHash,
			})
		}
		return c.NoContent(http.StatusBadRequest)
	}
}
