package content_item

import (
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"github.com/jairogloz/go-content-manager/pkg/repositories/content_item"
)

func CreateContentItem(contentItemCreateParams domain.ContentItemCreateParams, config domain.EnvVars) (contentItem *domain.ContentItem, err error) {

	now := time.Now().UTC()
	contentItem = &domain.ContentItem{
		Category:    contentItemCreateParams.Category,
		Description: contentItemCreateParams.Description,
		Title:       contentItemCreateParams.Title,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	err = content_item.InsertContentItem(contentItem, config)
	if err != nil {
		return nil, fmt.Errorf("error inserting content item: %w", err)
	}

	return contentItem, nil
}
