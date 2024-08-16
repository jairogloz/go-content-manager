package content_item

import (
	"context"
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) List(userID string, page, limit int, sortByField, sortByOrder, category string) (contentItems []*domain.ContentItem, err error) {

	// Compute skip out of page and limit
	skip := (page - 1) * limit

	sortByOrderInt := -1
	if sortByOrder == "asc" {
		sortByOrderInt = 1
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: sortByField, Value: sortByOrderInt}})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	if category != "" {
		filter["category"] = bson.M{
			"$regex":   category,
			"$options": "i",
		}
	}

	cursor, err := r.coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error listing content items: %w", err)
	}

	if err = cursor.All(ctx, &contentItems); err != nil {
		return nil, fmt.Errorf("error decoding content items: %w", err)
	}

	return contentItems, nil
}
