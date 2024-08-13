package ports

import "github.com/jairogloz/go-content-manager/pkg/domain"

type ContentItemRepository interface {
	Insert(contentItem *domain.ContentItem) (insertedID string, err error)
	Find(id string) (contentItem *domain.ContentItem, err error)
}

type ContentItemService interface {
	Create(contentItemCreateParams domain.ContentItemCreateParams) (contentItem *domain.ContentItem, err error)
	Get(id string) (contentItem *domain.ContentItem, err error)
}
