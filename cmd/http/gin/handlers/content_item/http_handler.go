package content_item_hdlr

import (
	"fmt"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"github.com/jairogloz/go-content-manager/pkg/ports"
)

type HttpHandler struct {
	Service ports.ContentItemService
	Config  domain.EnvVars
}

func NewHttpHandler(service ports.ContentItemService, config domain.EnvVars) (*HttpHandler, error) {
	if service == nil {
		return nil, fmt.Errorf("error creating content item http handler: service is nil")
	}

	return &HttpHandler{
		Service: service,
		Config:  config,
	}, nil
}
