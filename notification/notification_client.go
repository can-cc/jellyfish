package notification

import "github.com/labstack/echo/middleware"

type Client interface {
	GetTargets() ([]*middleware.ProxyTarget, error)
}
