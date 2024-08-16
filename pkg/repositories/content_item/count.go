package content_item

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) Count(userID, category string) (count int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	if category != "" {
		filter["category"] = category
	}

	documentCount, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("error counting content items: %w", err)
	}

	return int(documentCount), nil
}
