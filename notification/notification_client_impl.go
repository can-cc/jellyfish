package notification

import (
	"net/url"

	"github.com/juju/errors"
	"github.com/labstack/echo/middleware"
)

type ClientImpl struct {
	Endpoint string
}

func (client ClientImpl) GetTargets() ([]*middleware.ProxyTarget, error) {
	url1, err := url.Parse(client.Endpoint)
	if err != nil {
		return nil, errors.Trace(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}
	return targets, nil
}
