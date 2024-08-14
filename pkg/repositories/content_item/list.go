package content_item

import (
	"context"
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) List(userID string) (contentItems []*domain.ContentItem, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.coll.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("error listing content items: %w", err)
	}

	if err = cursor.All(ctx, &contentItems); err != nil {
		return nil, fmt.Errorf("error decoding content items: %w", err)
	}

	return contentItems, nil
}
