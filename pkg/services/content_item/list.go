package content_item

import (
	"fmt"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) List(userID string, page, limit int) (contentItems []*domain.ContentItem, err error) {

	if page < 1 {
		return nil, fmt.Errorf("error listing content items: page must be greater than or equal to 1")
	}

	if limit < 1 || limit > 100 {
		return nil, fmt.Errorf("error listing content items: limit must be between 1 and 100")
	}

	contentItems, err = s.Repo.List(userID, page, limit)
	if err != nil {
		return nil, err
	}

	return contentItems, nil

}
