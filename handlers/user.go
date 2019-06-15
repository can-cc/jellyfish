package handlers

import (
	"crypto/md5"
	sql2 "database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/fwchen/jellyfish/database"
	"gopkg.in/guregu/null.v3"

	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
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

// PostAvatarByBase64 :
func PostAvatarByBase64() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := database.GetDB()

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		userID := claims.ID

		request := new(struct {
			AvatarData string `json:"avatarData"`
		})

		c.Bind(&request)

		avatardir := viper.GetString("avatar-dir")
		fileNameHash := GetMD5Hash(request.AvatarData)

		fm, err := saveImageToDisk(avatardir+fileNameHash, request.AvatarData)
		if err != nil {
			panic(err)
		}

		_, err2 := db.Exec(`UPDATE users set avatar = $1 where id = $2`, fileNameHash+"."+fm, userID)

		if err2 != nil {
			panic(err2)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"avatarUrl": avatardir + fileNameHash + "." + fm,
		})

	}
}

// GetUserInfo :
func GetUserInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := database.GetDB()

		userID := c.Param("userId")

		sql := `SELECT username, avatar FROM users where id = $1`
		row := db.QueryRow(sql, userID)

		userInfo := new(struct {
			Username string
			Avatar   sql2.NullString
		})
		avatarDir := viper.GetString("avatar-dir")

		err := row.Scan(&userInfo.Username, &userInfo.Avatar)
		if err != nil {
			panic(err)
		}

		var avatarURL null.String
		if userInfo.Avatar.Valid {
			avatarURL.Valid = true
			avatarURL.String = avatarDir + userInfo.Avatar.String
		}

		return c.JSON(http.StatusOK, struct {
			Username  string      `json:"username"`
			avatarURL null.String `json:"avatarUrl"`
		}{
			Username:  userInfo.Username,
			avatarURL: avatarURL,
		})
	}
}

// PostAvatar :
func PostAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := database.GetDB()

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		userID := claims.ID

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
			_, err3 := stmt.Exec(fileNameHash, userID)
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
