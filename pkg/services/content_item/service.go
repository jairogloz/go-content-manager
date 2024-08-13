package content_item

import (
	"fmt"

	"github.com/jairogloz/go-content-manager/pkg/ports"
)

var _ ports.ContentItemService = &Service{}

type Service struct {
	Repo ports.ContentItemRepository
}

func NewService(repo ports.ContentItemRepository) (*Service, error) {
	if repo == nil {
		return nil, fmt.Errorf("error creating content item service: repository is nil")
	}

	return &Service{
		Repo: repo,
	}, nil
}
