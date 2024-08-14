package content_item

import "github.com/jairogloz/go-content-manager/pkg/domain"

func (s *Service) List(userID string) (contentItems []*domain.ContentItem, err error) {
	contentItems, err = s.Repo.List(userID)
	if err != nil {
		return nil, err
	}

	return contentItems, nil

}
