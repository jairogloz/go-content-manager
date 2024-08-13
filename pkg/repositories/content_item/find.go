package content_item

import (
	"context"
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Find(id string) (contentItem *domain.ContentItem, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	singleResult := r.coll.FindOne(ctx, bson.M{"_id": id})
	if err := singleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró el documento, pero no hubo error en la consulta
			return nil, nil
		}
		return nil, fmt.Errorf("error finding content item: %w", err) // Error finding the document
	}

	if err := singleResult.Decode(&contentItem); err != nil {
		return nil, err // Error decoding the document
	}

	return contentItem, nil

}
