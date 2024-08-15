package ports

import (
	"github.com/jairogloz/go-content-manager/pkg/domain"
)

type ContentItemRepository interface {
	Insert(contentItem *domain.ContentItem) (insertedID string, err error)
	Find(id string) (contentItem *domain.ContentItem, err error)
	Delete(id string) error
	Update(id string, contentItem domain.ContentItemUpdateParams) (updatedCount int, err error)
	List(userID string, page, limit int) (contentItems []*domain.ContentItem, err error)
}

type ContentItemService interface {
	Create(userID string, contentItemCreateParams domain.ContentItemCreateParams) (contentItem *domain.ContentItem, err error)
	Get(id string) (contentItem *domain.ContentItem, err error)
	Delete(id string) error
	Update(id string, contentItem domain.ContentItemUpdateParams) (updatedCount int, err error)
	List(userID string, page, limit int) (contentItems []*domain.ContentItem, err error)
}
