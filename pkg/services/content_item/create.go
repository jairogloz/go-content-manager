package content_item

import (
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) Create(contentItemCreateParams domain.ContentItemCreateParams) (contentItem *domain.ContentItem, err error) {

	now := time.Now().UTC()
	contentItem = &domain.ContentItem{
		Category:    contentItemCreateParams.Category,
		Description: contentItemCreateParams.Description,
		Title:       contentItemCreateParams.Title,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	contentItemID, err := s.Repo.Insert(contentItem)
	if err != nil {
		return nil, fmt.Errorf("error inserting content item: %w", err)
	}
	log.Println("Content item inserted with ID: ", contentItemID)

	return contentItem, nil
}
