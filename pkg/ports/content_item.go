package ports

import "github.com/jairogloz/go-content-manager/pkg/domain"

type ContentItemRepository interface {
	Insert(contentItem *domain.ContentItem) (insertedID string, err error)
}

type ContentItemService interface {
	Create(contentItemCreateParams domain.ContentItemCreateParams) (contentItem *domain.ContentItem, err error)
}
