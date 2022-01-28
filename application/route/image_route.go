package route

import (
	"jellyfish/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

func GetImageRoute(imageStorageService *service.ImageStorageService) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileName := c.Param("fileName")
		reader, err := imageStorageService.GetImage(fileName)
		if err != nil {
			return errors.Trace(err)
		}
		return c.Stream(200, "image/png", reader)
	}
}
