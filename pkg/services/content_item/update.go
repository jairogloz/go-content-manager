package content_item

import (
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) Update(id string, contentItemUpdateParams domain.ContentItemUpdateParams) (updatedCount int, err error) {
	now := time.Now().UTC()
	contentItemUpdateParams.UpdatedAt = &now
	updateCount, err := s.Repo.Update(id, contentItemUpdateParams)
	if err != nil {
		return 0, fmt.Errorf("error getting content item: %w", err)
	}

	return updateCount, nil
}
