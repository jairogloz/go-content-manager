package content_item

import (
	"fmt"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) Get(id string) (contentItem *domain.ContentItem, err error) {
	contentItem, err = s.Repo.Find(id)
	if err != nil {
		return nil, fmt.Errorf("error getting content item: %w", err)
	}

	return contentItem, nil
}
