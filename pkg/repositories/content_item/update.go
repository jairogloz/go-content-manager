package content_item

import (
	"context"
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) Update(id string, contentItemUpdateParam domain.ContentItemUpdateParams) (updatedCount int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	updateResult, err := r.coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": contentItemUpdateParam})
	if err != nil {
		return 0, fmt.Errorf("error updating content item: %w", err)
	}

	return int(updateResult.ModifiedCount), nil
}
