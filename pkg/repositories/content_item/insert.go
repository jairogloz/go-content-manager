package content_item

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Insert(contentItem *domain.ContentItem) (insertedID string, err error) {

	contentItem.ID = primitive.NewObjectID().Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = r.coll.InsertOne(ctx, contentItem)
	if err != nil {
		log.Println("Failed to insert contentItem to MongoDB: ", err.Error())
		return "", fmt.Errorf("error inserting contentItem to MongoDB: %w", err)
	}

	return contentItem.ID, nil
}
